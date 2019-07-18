package forward

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"
	configurationpkg "github.com/mutagen-io/mutagen/pkg/configuration"
	"github.com/mutagen-io/mutagen/pkg/filesystem"
	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/grpcutil"
	"github.com/mutagen-io/mutagen/pkg/prompt"
	"github.com/mutagen-io/mutagen/pkg/selection"
	forwardingsvc "github.com/mutagen-io/mutagen/pkg/service/forwarding"
	"github.com/mutagen-io/mutagen/pkg/url"
	forwardingurl "github.com/mutagen-io/mutagen/pkg/url/forwarding"
)

// loadAndValidateYAMLConfiguration loads a YAML-based configuration, converts
// it to a Protocol Buffers session configuration, and validates it. If the path
// doesn't exist, this function returns a default configuration object.
func loadAndValidateYAMLConfiguration(path string) (*forwarding.Configuration, error) {
	// Load the YAML configuration.
	yamlConfiguration, err := configurationpkg.Load(path)
	if err != nil {
		return nil, err
	}

	// Convert the YAML configuration to a Protocol Buffers representation and
	// validate it.
	configuration := yamlConfiguration.Forwarding.Defaults.Configuration()
	if err := configuration.EnsureValid(false); err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}

	// Success.
	return configuration, nil
}

func createMain(command *cobra.Command, arguments []string) error {
	// Validate, extract, and parse URLs.
	if len(arguments) != 2 {
		return errors.New("invalid number of endpoint URLs provided")
	}
	source, err := url.Parse(arguments[0], url.Kind_Forwarding, true)
	if err != nil {
		return errors.Wrap(err, "unable to parse source URL")
	}
	destination, err := url.Parse(arguments[1], url.Kind_Forwarding, false)
	if err != nil {
		return errors.Wrap(err, "unable to parse destination URL")
	}

	// If either URL is a local Unix domain socket path, make sure it's
	// normalized.
	if source.Protocol == url.Protocol_Local {
		if protocol, path, err := forwardingurl.Parse(source.Path); err != nil {
			return errors.Wrap(err, "unable to parse source forwarding endpoint URL")
		} else if protocol == "unix" {
			if normalized, err := filesystem.Normalize(path); err != nil {
				return errors.Wrap(err, "unable to normalize source forwarding endpoint socket path")
			} else {
				source.Path = fmt.Sprintf("%s:%s", protocol, normalized)
			}
		}
	}
	if destination.Protocol == url.Protocol_Local {
		if protocol, path, err := forwardingurl.Parse(destination.Path); err != nil {
			return errors.Wrap(err, "unable to parse destination forwarding endpoint URL")
		} else if protocol == "unix" {
			if normalized, err := filesystem.Normalize(path); err != nil {
				return errors.Wrap(err, "unable to normalize destination forwarding endpoint socket path")
			} else {
				destination.Path = fmt.Sprintf("%s:%s", protocol, normalized)
			}
		}
	}

	// Validate the name.
	if err := selection.EnsureNameValid(createConfiguration.name); err != nil {
		return errors.Wrap(err, "invalid session name")
	}

	// Parse, validate, and record labels.
	var labels map[string]string
	if len(createConfiguration.labels) > 0 {
		labels = make(map[string]string, len(createConfiguration.labels))
	}
	for _, label := range createConfiguration.labels {
		components := strings.SplitN(label, "=", 2)
		var key, value string
		key = components[0]
		if len(components) == 2 {
			value = components[1]
		}
		if err := selection.EnsureLabelKeyValid(key); err != nil {
			return errors.Wrap(err, "invalid label key")
		} else if err := selection.EnsureLabelValueValid(value); err != nil {
			return errors.Wrap(err, "invalid label value")
		}
		labels[key] = value
	}

	// Create a default session configuration which will form the basis of our
	// cumulative configuration.
	configuration := &forwarding.Configuration{}

	// Unless disabled, attempt to load configuration from the global
	// configuration file and merge it into our cumulative configuration.
	if !createConfiguration.noGlobalConfiguration {
		// Compute the path to the global configuration file.
		globalConfigurationPath, err := configurationpkg.GlobalConfigurationPath()
		if err != nil {
			return errors.Wrap(err, "unable to compute path to global configuration file")
		}

		// Attempt to load the file. We allow it to not exist.
		globalConfiguration, err := loadAndValidateYAMLConfiguration(globalConfigurationPath)
		if err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrap(err, "unable to load global configuration")
			}
		} else {
			configuration = forwarding.MergeConfigurations(configuration, globalConfiguration)
		}
	}

	// If a configuration file has been specified, then load it and merge it
	// into our cumulative configuration.
	if createConfiguration.configurationFile != "" {
		if c, err := loadAndValidateYAMLConfiguration(createConfiguration.configurationFile); err != nil {
			return errors.Wrap(err, "unable to load configuration file")
		} else {
			configuration = forwarding.MergeConfigurations(configuration, c)
		}
	}

	// Validate and convert socket overwrite mode specifications.
	var socketOverwriteMode, socketOverwriteModeSource, socketOverwriteModeDestination forwarding.SocketOverwriteMode
	if createConfiguration.socketOverwriteMode != "" {
		if err := socketOverwriteMode.UnmarshalText([]byte(createConfiguration.socketOverwriteMode)); err != nil {
			return errors.Wrap(err, "unable to socket overwrite mode")
		}
	}
	if createConfiguration.socketOverwriteModeSource != "" {
		if err := socketOverwriteModeSource.UnmarshalText([]byte(createConfiguration.socketOverwriteModeSource)); err != nil {
			return errors.Wrap(err, "unable to socket overwrite mode for source")
		}
	}
	if createConfiguration.socketOverwriteModeDestination != "" {
		if err := socketOverwriteModeDestination.UnmarshalText([]byte(createConfiguration.socketOverwriteModeDestination)); err != nil {
			return errors.Wrap(err, "unable to socket overwrite mode for destination")
		}
	}

	// Validate socket owner specifications.
	if createConfiguration.socketOwner != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketOwner,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket ownership specification")
		}
	}
	if createConfiguration.socketOwnerSource != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketOwnerSource,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket ownership specification for source")
		}
	}
	if createConfiguration.socketOwnerDestination != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketOwnerDestination,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket ownership specification for destination")
		}
	}

	// Validate socket group specifications.
	if createConfiguration.socketGroup != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketGroup,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket group specification")
		}
	}
	if createConfiguration.socketGroupSource != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketGroupSource,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket group specification for source")
		}
	}
	if createConfiguration.socketGroupDestination != "" {
		if kind, _ := filesystem.ParseOwnershipIdentifier(
			createConfiguration.socketGroupDestination,
		); kind == filesystem.OwnershipIdentifierKindInvalid {
			return errors.New("invalid socket group specification for destination")
		}
	}

	// Validate and convert socket permission mode specifications.
	var socketPermissionMode, socketPermissionModeSource, socketPermissionModeDestination filesystem.Mode
	if createConfiguration.socketPermissionMode != "" {
		if err := socketPermissionMode.UnmarshalText([]byte(createConfiguration.socketPermissionMode)); err != nil {
			return errors.Wrap(err, "unable to parse socket permission mode")
		}
	}
	if createConfiguration.socketPermissionModeSource != "" {
		if err := socketPermissionModeSource.UnmarshalText([]byte(createConfiguration.socketPermissionModeSource)); err != nil {
			return errors.Wrap(err, "unable to parse socket permission mode for source")
		}
	}
	if createConfiguration.socketPermissionModeDestination != "" {
		if err := socketPermissionModeDestination.UnmarshalText([]byte(createConfiguration.socketPermissionModeDestination)); err != nil {
			return errors.Wrap(err, "unable to parse socket permission mode for destination")
		}
	}

	// Create the command line configuration and merge it into our cumulative
	// configuration.
	configuration = forwarding.MergeConfigurations(configuration, &forwarding.Configuration{
		SocketOverwriteMode:  socketOverwriteMode,
		SocketOwner:          createConfiguration.socketOwner,
		SocketGroup:          createConfiguration.socketGroup,
		SocketPermissionMode: uint32(socketPermissionMode),
	})

	// Connect to the daemon and defer closure of the connection.
	daemonConnection, err := daemon.CreateClientConnection(true, true)
	if err != nil {
		return errors.Wrap(err, "unable to connect to daemon")
	}
	defer daemonConnection.Close()

	// Create a session service client.
	sessionService := forwardingsvc.NewForwardingClient(daemonConnection)

	// Invoke the session create method. The stream will close when the
	// associated context is cancelled.
	createContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := sessionService.Create(createContext)
	if err != nil {
		return errors.Wrap(grpcutil.PeelAwayRPCErrorLayer(err), "unable to invoke create")
	}

	// Send the initial request.
	request := &forwardingsvc.CreateRequest{
		Specification: &forwardingsvc.CreationSpecification{
			Source:        source,
			Destination:   destination,
			Configuration: configuration,
			ConfigurationSource: &forwarding.Configuration{
				SocketOverwriteMode:  socketOverwriteModeSource,
				SocketOwner:          createConfiguration.socketOwnerSource,
				SocketGroup:          createConfiguration.socketGroupSource,
				SocketPermissionMode: uint32(socketPermissionModeSource),
			},
			ConfigurationDestination: &forwarding.Configuration{
				SocketOverwriteMode:  socketOverwriteModeDestination,
				SocketOwner:          createConfiguration.socketOwnerDestination,
				SocketGroup:          createConfiguration.socketGroupDestination,
				SocketPermissionMode: uint32(socketPermissionModeDestination),
			},
			Name:   createConfiguration.name,
			Labels: labels,
			Paused: createConfiguration.paused,
		},
	}
	if err := stream.Send(request); err != nil {
		return errors.Wrap(grpcutil.PeelAwayRPCErrorLayer(err), "unable to send create request")
	}

	// Create a status line printer and defer a break.
	statusLinePrinter := &cmd.StatusLinePrinter{}
	defer statusLinePrinter.BreakIfNonEmpty()

	// Receive and process responses until we're done.
	for {
		if response, err := stream.Recv(); err != nil {
			return errors.Wrap(grpcutil.PeelAwayRPCErrorLayer(err), "create failed")
		} else if err = response.EnsureValid(); err != nil {
			return errors.Wrap(err, "invalid create response received")
		} else if response.Session != "" {
			statusLinePrinter.Print(fmt.Sprintf("Created session %s", response.Session))
			return nil
		} else if response.Message != "" {
			statusLinePrinter.Print(response.Message)
			if err := stream.Send(&forwardingsvc.CreateRequest{}); err != nil {
				return errors.Wrap(grpcutil.PeelAwayRPCErrorLayer(err), "unable to send message response")
			}
		} else if response.Prompt != "" {
			statusLinePrinter.BreakIfNonEmpty()
			if response, err := prompt.PromptCommandLine(response.Prompt); err != nil {
				return errors.Wrap(err, "unable to perform prompting")
			} else if err = stream.Send(&forwardingsvc.CreateRequest{Response: response}); err != nil {
				return errors.Wrap(grpcutil.PeelAwayRPCErrorLayer(err), "unable to send prompt response")
			}
		}
	}
}

var createCommand = &cobra.Command{
	Use:          "create <source> <destination>",
	Short:        "Create and start a new forwarding session",
	RunE:         createMain,
	SilenceUsage: true,
}

var createConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
	// name is the name specification for the session.
	name string
	// labels are the label specifications for the session.
	labels []string
	// paused indicates whether or not to create the session in a pre-paused
	// state.
	paused bool
	// noGlobalConfiguration specifies whether or not the global configuration
	// file should be ignored.
	noGlobalConfiguration bool
	// configurationFile specifies a file from which to load configuration. It
	// should be a path relative to the working directory.
	configurationFile string
	// socketOverwriteMode specifies the socket overwrite mode to use for the
	// session.
	socketOverwriteMode string
	// socketOverwriteModeSource specifies the socket overwrite mode to use for
	// the session, taking priority over socketOverwriteMode on source if
	// specified.
	socketOverwriteModeSource string
	// socketOverwriteModeDestination specifies the socket overwrite mode to use
	// for the session, taking priority over socketOverwriteMode on destination
	// if specified.
	socketOverwriteModeDestination string
	// socketOwner specifies the socket owner identifier to use new Unix domain
	// socket listeners, with endpoint-specific specifications taking priority.
	socketOwner string
	// socketOwnerSource specifies the socket owner identifier to use new Unix
	// domain socket listeners, taking priority over socketOwner on source if
	// specified.
	socketOwnerSource string
	// socketOwnerDestination specifies the socket owner identifier to use new
	// Unix domain socket listeners, taking priority over socketOwner on
	// destination if specified.
	socketOwnerDestination string
	// socketGroup specifies the socket owner identifier to use new Unix domain
	// socket listeners, with endpoint-specific specifications taking priority.
	socketGroup string
	// socketGroupSource specifies the socket owner identifier to use new Unix
	// domain socket listeners, taking priority over socketGroup on source if
	// specified.
	socketGroupSource string
	// socketGroupDestination specifies the socket owner identifier to use new
	// Unix domain socket listeners, taking priority over socketGroup on
	// destination if specified.
	socketGroupDestination string
	// socketPermissionMode specifies the socket permission mode to use for new
	// Unix domain socket listeners, with endpoint-specific specifications
	// taking priority.
	socketPermissionMode string
	// socketPermissionModeSource specifies the socket permission mode to use
	// for new Unix domain socket listeners on source, taking priority over
	// socketPermissionMode on source if specified.
	socketPermissionModeSource string
	// socketPermissionModeDestination specifies the socket permission mode to
	// use for new Unix domain socket listeners on destination, taking priority
	// over socketPermissionMode on destination if specified.
	socketPermissionModeDestination string
}

func init() {
	// Grab a handle for the command line flags.
	flags := createCommand.Flags()

	// Disable sourcebetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&createConfiguration.help, "help", "h", false, "Show help information")

	// Wire up name and label flags.
	flags.StringVarP(&createConfiguration.name, "name", "n", "", "Specify a name for the session")
	flags.StringSliceVarP(&createConfiguration.labels, "label", "l", nil, "Specify labels")

	// Wire up paused flags.
	flags.BoolVarP(&createConfiguration.paused, "paused", "p", false, "Create the session pre-paused")

	// Wire up general configuration flags.
	flags.BoolVar(&createConfiguration.noGlobalConfiguration, "no-global-configuration", false, "Ignore the global configuration file")
	flags.StringVarP(&createConfiguration.configurationFile, "configuration-file", "c", "", "Specify a file from which to load session configuration")

	// Wire up socket flags.
	flags.StringVar(&createConfiguration.socketOverwriteMode, "socket-overwrite-mode", "", "Specify socket overwrite mode (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOverwriteModeSource, "socket-overwrite-mode-source", "", "Specify socket overwrite mode for source (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOverwriteModeDestination, "socket-overwrite-mode-destination", "", "Specify socket overwrite mode for destination (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOwner, "socket-owner", "", "Specify socket owner")
	flags.StringVar(&createConfiguration.socketOwnerSource, "socket-owner-source", "", "Specify socket owner for source")
	flags.StringVar(&createConfiguration.socketOwnerDestination, "socket-owner-destination", "", "Specify socket owner for destination")
	flags.StringVar(&createConfiguration.socketGroup, "socket-group", "", "Specify socket group")
	flags.StringVar(&createConfiguration.socketGroupSource, "socket-group-source", "", "Specify socket group for source")
	flags.StringVar(&createConfiguration.socketGroupDestination, "socket-group-destination", "", "Specify socket group for destination")
	flags.StringVar(&createConfiguration.socketPermissionMode, "socket-permission-mode", "", "Specify socket permission mode")
	flags.StringVar(&createConfiguration.socketPermissionModeSource, "socket-permission-mode-source", "", "Specify socket permission mode for source")
	flags.StringVar(&createConfiguration.socketPermissionModeDestination, "socket-permission-mode-destination", "", "Specify socket permission mode for destination")
}
