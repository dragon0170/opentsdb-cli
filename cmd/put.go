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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type DataPoint struct {
	Metric string  `json:"metric"`
	Timestamp int  `json:"timestamp"`
	Value int  `json:"value"`
	Tags map[string]string  `json:"tags"`
}

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("put called")
		timestamp, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		value, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		tags := strings.Split(args[3], "=")

		dataPoint := DataPoint{args[0], timestamp, value, map[string]string{tags[0]: tags[1]}}
		dbytes, err := json.Marshal(dataPoint)
//		fmt.Printf("%s\n", dbytes)
		buff := bytes.NewBuffer(dbytes)
		resp, err := http.Post("http://10.0.81.180:30096/api/put?details", "application/json", buff)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		// Response 체크.
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
//        str := string(respBody)
//        println(str)
		}

		var indentedJSON bytes.Buffer

		err = json.Indent(&indentedJSON, respBody, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", indentedJSON.Bytes())

	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
