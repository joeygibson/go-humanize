package main

//func main() {
//	if len(os.Args) < 2 {
//		fmt.Println("Usage: thou <number> [... <number>]")
//		os.Exit(1)
//	}
//
//	for _, number := range os.Args[1:] {
//		actualNumber, err := strconv.Atoi(number)
//		if err != nil {
//			fmt.Printf("Error converting %s to number: %v\n", number, err)
//			os.Exit(1)
//		}
//
//		numberWithCommas := humanize.Comma(int64(actualNumber))
//		fmt.Printf("%s -> %s\n", number, numberWithCommas)
//	}
//}

import (
	"fmt"
	"github.com/joeygibson/go-humanize/pkg/humanize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Use:   "humanize",
		Short: "Convert things to human-friendly versions",
		Long:  "Convert things to human-friendly versions",
		Run:   CmdRoot,
	}

	version string
	build   string
)

func CmdRoot(cmd *cobra.Command, args []string) {
	if viper.GetBool("help") {
		cmd.Help()
		os.Exit(0)
	}

	if viper.GetBool("commas") {
		for _, arg := range args {
			if num, err := strconv.ParseInt(arg, 10, 64); err == nil {
				val := humanize.Comma(num)
				fmt.Printf("%s -> %s\n", arg, val)

				continue
			}

			if num, err := strconv.ParseFloat(arg, 64); err == nil {
				val := humanize.Commaf(num)
				fmt.Printf("%s -> %s\n", arg, val)

				continue
			}

			fmt.Printf("%s can't be converted to a number\n", arg)
		}
	}
}

func init() {
	// Look for env vars starting with `HUMANIZE`, replacing `.` in keys
	// with `_` for env vars. Automatically bind what it finds
	viper.SetEnvPrefix("HUMANIZE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	rootCmd.Flags().BoolP("commas", "c", false, "format big numbers with thousands separators")

	viper.BindPFlag("commas", rootCmd.Flags().Lookup("commas"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
