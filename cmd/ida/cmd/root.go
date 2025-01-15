/*
Copyright © 2025 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/apex/log"
	clihander "github.com/apex/log/handlers/cli"
	"github.com/blacktop/go-idalib"
	"github.com/spf13/cobra"
)

var verbose bool

func init() {
	log.SetHandler(clihander.Default)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Enable verbose logging")
	// rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output directory")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "ida",
	Short:         "Run ida commands",
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {

		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		ida := idalib.NewIDALib()
		defer idalib.DeleteIDALib(ida)

		if ret := ida.Init(); ret != 0 {
			log.Fatalf("Failed to initialize IDA library: %d", ret)
		}

		if verbose {
			ida.EnableConsoleMessages(true)
		}

		if ret := ida.OpenDatabase(args[0], false); ret != 0 {
			log.Fatalf("Failed to open database: %d", ret)
		}
		defer ida.CloseDatabase(true)

		// count := ida.FuncCount()
		log.Infof("Function count: %d", ida.FuncCount())

		fn := ida.GetFuncByID(0)
		log.Infof("1st Function: %#v", fn)

		dec := ida.DecompileFunction(fn)
		if dec == "" {
			log.Fatal("Failed to decompile function")
		}
		fmt.Println("Decompiled Function:")
		if err := quick.Highlight(os.Stdout, dec, "c", "terminal256", "nord"); err != nil {
			log.Fatalf("failed to highlight json: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
