package cmd

import (
    "log"

    "github.com/spf13/cobra"
    "github.com/z3oxs/goshot/modes"
)

var rootCmd = cobra.Command {
    Short: "Simple tool to take screenshots",
    Use: "goshot",
    Run: func(c *cobra.Command, args []string) {
        fullscreen, _ := c.Flags().GetBool("fullscreen")
        selection, _ := c.Flags().GetBool("selection")
        display, _ := c.Flags().GetInt("display")
        clipboard, _ := c.Flags().GetBool("clipboard")
        save, _ := c.Flags().GetString("save")
        output, _ := c.Flags().GetBool("output")
        upload, _ := c.Flags().GetBool("upload")
        selectiongui, _ := c.Flags().GetBool("selectiongui")

        if (!clipboard && save == "" && !output && !upload) {
            log.Fatalln("You need to specify at least a save method, check 'goshot --help'.")

        }

        if fullscreen {
           modes.Fullscreen(save, clipboard, output, upload)

        } else if selection {
            modes.Selection(save, clipboard, output, upload)

        } else if selectiongui {
            modes.SelectionGUI(save, clipboard, output, upload)

        } else if display != -1 {
            modes.Display(save, clipboard, output, upload, display)

        } else {
            c.Help()

        }
    },
}

func init() {
    rootCmd.PersistentFlags().BoolP("fullscreen", "f", false, "Fullscreen mode")
    rootCmd.PersistentFlags().BoolP("selection", "s", false, "Selection mode")
    rootCmd.PersistentFlags().BoolP("clipboard", "c", false, "Save output to clipboard")
    rootCmd.PersistentFlags().IntP("display", "d", -1, "Display mode")
    rootCmd.PersistentFlags().StringP("save", "S", "", "Save output to file")
    rootCmd.PersistentFlags().BoolP("output", "o", false, "Outputs screenshot to stdout")
    rootCmd.PersistentFlags().BoolP("upload", "u", false, "Upload the screenshot to AnonFiles")
    rootCmd.PersistentFlags().BoolP("selectiongui", "g", false, "Selection with GUI mode.")

}

func Execute() {
    rootCmd.Execute()

}
