package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	topMoviesAPIPath  = "/top/movies"
	topTVShowsAPIPath = "/top/tv-shows"
	keyLimit          = "limit"
	keyNumVotes       = "numVotes"
)

var rootCmd = &cobra.Command{Use: "./cli"}

var topMoviesCmd = &cobra.Command{
	Use:   "top-movies",
	Short: "Get top movies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("top-movies called")
		callTopTitlesAPI(topMoviesAPIPath, viper.GetInt("limit"), viper.GetInt("num-votes"))
	},
}

var topTVShowsCmd = &cobra.Command{
	Use:   "top-tv-shows",
	Short: "Get top tv shows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("top-tv-shows called", args)
		callTopTitlesAPI(topTVShowsAPIPath, viper.GetInt("limit"), viper.GetInt("num-votes"))
	},
}

var baseURL = "http://localhost:8000"

func init() {
	rootCmd.AddCommand(topMoviesCmd)
	rootCmd.AddCommand(topTVShowsCmd)

	topMoviesCmd.Flags().IntP("limit", "l", 10, "Limit")
	topMoviesCmd.Flags().IntP("num-votes", "n", 1000, "Num votes")
	topMoviesCmd.Flags().String("base-url", baseURL, "Base URL")
	viper.BindPFlag("limit", topMoviesCmd.Flags().Lookup("limit"))
	viper.BindPFlag("num-votes", topMoviesCmd.Flags().Lookup("num-votes"))
	viper.BindPFlag("base-url", topMoviesCmd.Flags().Lookup("base-url"))

	topTVShowsCmd.Flags().IntP("limit", "l", 10, "Limit")
	topTVShowsCmd.Flags().IntP("num-votes", "n", 1000, "Num votes")
	topTVShowsCmd.Flags().String("base-url", baseURL, "Base URL")
	viper.BindPFlag("limit", topTVShowsCmd.Flags().Lookup("limit"))
	viper.BindPFlag("num-votes", topTVShowsCmd.Flags().Lookup("num-votes"))
	viper.BindPFlag("base-url", topTVShowsCmd.Flags().Lookup("base-url"))
}

func callTopTitlesAPI(urlStr string, limit, numVotes int) {
	params := url.Values{keyLimit: {strconv.Itoa(limit)}, keyNumVotes: {strconv.Itoa(numVotes)}}
	response, err := makeHTTPGetRequest(urlStr, params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	formattedResponse, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(formattedResponse))
}

func makeHTTPGetRequest(urlStr string, params url.Values) (map[string]interface{}, error) {
	base := viper.GetString("base-url")
	if base == "" {
		base = baseURL
	}
	urlStr = base + urlStr
	urlStr = urlStr + "?" + params.Encode()
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var responseMessage map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseMessage)
	if err != nil {
		return nil, err
	}
	return responseMessage, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
