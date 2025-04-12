/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the joke notifier",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			joke := getRandomJoke()
			if joke != "" {
				beeep.Notify("Dad Joke", joke, "")
			}
			time.Sleep(6 * time.Hour)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
