package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
)

// Runtime exit codes
var exitOK = 0
var exitUnknown = 1

// All variables here are provided as build-time arguments, with clear default values
var (
	appVersion = "[Not Provided]"
	cliVersion = "[Not Provided]"
	buildDate  = "[Not Provided]"
	gitCommit  = "[Not Provided]"
)

func main() {
	run()
}

func run() int {
	// Colorized logging output as the default slog logger
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  false,
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
	})))
	return exitOK
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "my-go-app",
	}

	versionCmd := &cobra.Command{
		Use: "version",
		Run: runVersion,
	}
	cmd.AddCommand(versionCmd)

	return cmd
}

func runVersion(cmd *cobra.Command, _ []string) {

	versionString := fmt.Sprintf(`
CLIVersion:     %s
GitCommit:      %s
Build Date:     %s
Platform:       %s
Go Version:     %s
Compiler:       %s
`,
		appVersion,
		gitCommit,
		buildDate,
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		runtime.Version(),
		runtime.Compiler,
	)

	strings.NewReader(versionString).WriteTo(cmd.OutOrStdout())
}
