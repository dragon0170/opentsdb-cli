// Copyright © 2018 Kevin Choi <kevin.choi@paust.io>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
	"encoding/json"
//	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Query struct {
	Aggregator string `json:"aggregator"`
	Metric string `json:"metric"`
	Tags map[string]string `json:"tags"`
}

type Request struct {
	Start string `json:"start"`
	End string `json:"json"`
	Queries []Query `json:"queries"`
}

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("query called")
		tagMap := make(map[string]string)
		for _, a := range args[4:] {
			tags := strings.Split(a, "=")
			tagMap[tags[0]] = tags[1]
		}
//		query := Query{args[2], args[3], tagMap}
		request := Request{args[0], args[1], []Query{Query{args[2], args[3], tagMap}}}
		rbytes, err := json.Marshal(request)
		if err != nil {
			panic(err)
		}
//		fmt.Printf("%s\n", rbytes)
		buff := bytes.NewBuffer(rbytes)
		resp, err := http.Post("http://10.0.81.180:30096/api/query", "application/json", buff)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if len(respBody) == 0 {
			fmt.Printf("No Data\n")
		} else {
			var indentedJSON bytes.Buffer

			err = json.Indent(&indentedJSON, respBody, "", "\t")
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", indentedJSON.Bytes())
		}

	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
