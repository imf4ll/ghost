package cmd

import (
    "github.com/z3oxs/goshot/modes"
    "github.com/spf13/cobra"
)

var rootCmd = cobra.Command {
    Short: "Simple tool to take screenshots",
    Use: "goshot",
    Run: func(c *cobra.Command, args []string) {
        fullscreen, _ := c.Flags().GetBool("fullscreen")
        selection, _ := c.Flags().GetBool("selection")
        monitor, _ := c.Flags().GetInt("monitor")
        clipboard, _ := c.Flags().GetBool("clipboard")
        save, _ := c.Flags().GetString("save")
        output, _ := c.Flags().GetBool("output")

        if fullscreen {
           modes.Fullscreen(save, clipboard, output)

        } else if selection {
            modes.Selection(save, clipboard, output)

        } else if monitor != -1 {
            modes.Monitor(save, clipboard, monitor, output)

        } else {
            c.Help()

        }
    },
}

func init() {
    rootCmd.PersistentFlags().BoolP("fullscreen", "f", false, "Fullscreen Mode")
    rootCmd.PersistentFlags().BoolP("selection", "s", false, "Selection mode")
    rootCmd.PersistentFlags().BoolP("clipboard", "c", false, "Save output to clipboard")
    rootCmd.PersistentFlags().IntP("monitor", "m", -1, "Monitor mode")
    rootCmd.PersistentFlags().StringP("save", "S", "", "Save output to file")
    rootCmd.PersistentFlags().BoolP("output", "o", false, "Outputs screenshot to stdout")
}

func Execute() {
    rootCmd.Execute()

}
