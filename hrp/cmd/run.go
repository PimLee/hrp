package cmd

import (
	"github.com/spf13/cobra"

	"github.com/httprunner/hrp"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run path...",
	Short: "run API test",
	Long:  `run yaml/json testcase files for API test`,
	Example: `  $ hrp run demo.json	# run specified json testcase file
  $ hrp run demo.yaml	# run specified yaml testcase file
  $ hrp run examples/	# run testcases in specified folder`,
	Args: cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var paths []hrp.ITestCase
		for _, arg := range args {
			paths = append(paths, &hrp.TestCasePath{Path: arg})
		}
		runner := hrp.NewRunner(nil).
			SetDebug(!silentFlag).
			SetFailfast(!continueOnFailure)
		if proxyUrl != "" {
			runner.SetProxyUrl(proxyUrl)
		}
		return runner.Run(paths...)
	},
}

var (
	continueOnFailure bool
	silentFlag        bool
	proxyUrl          string
)

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVar(&continueOnFailure, "continue-on-failure", false, "continue running next step when failure occurs")
	runCmd.Flags().BoolVarP(&silentFlag, "silent", "s", false, "disable logging request & response details")
	runCmd.Flags().StringVarP(&proxyUrl, "proxy-url", "p", "", "set proxy url")
	// runCmd.Flags().BoolP("gen-html-report", "r", false, "Generate HTML report")
}
