/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new entries",
	Long: `This command inserts new data to the 
	Phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty!")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		if number == "" {
			fmt.Println("Name is empty!")
			return
		}

		surname, _ := cmd.Flags().GetString("surname")
		if number == "" {
			fmt.Println("Surname is empty!")
			return
		}

		// Create request in two steps for readability
		URL := "http://" + SERVER + ":" + PORT + "/insert/"
		URL = URL + "/" + name + "/" + surname + "/" + number

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println("**", err)
			return
		}

		// Check HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		// Read data
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println("*", err)
			return
		}

		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP("name", "n", "", "Name value")
	insertCmd.Flags().StringP("surname", "s", "", "Surname value")
	insertCmd.Flags().StringP("tel", "t", "", "Telephone value")
}
