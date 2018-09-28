package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var seperator = "-----------------------------------------"
var seed = genSeed()
var listenPort string
var excludedKeys = []string{"authorization", "password"}

func main() {
	listenPort = getEnv("PORT", "8080")

	logStart()

	http.HandleFunc("/", globHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil))
}

func globHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req)

	msg := fmt.Sprintf("[%d %s] It's all good baby, baby\n", seed, time.Now().Format("15:04:05"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func logStart() {
	fmt.Println(seperator)
	fmt.Printf("Running with seed %d\n", seed)
	fmt.Printf("Listening on %s\n", listenPort)
	fmt.Println(seperator)
}

func logRequest(req *http.Request) {
	fmt.Println()
	fmt.Println()
	fmt.Println(seperator)
	fmt.Printf("REQUEST @ %s\n", time.Now().Format(time.RFC3339))
	fmt.Println(seperator)
	fmt.Printf("%v %v %v\n", req.Method, req.Proto, req.URL.Path)
	fmt.Printf("%v%v%v\n", req.URL.Scheme, req.Host, req.URL.RequestURI())
	fmt.Println()

	logMap("Query Params", req.URL.Query())
	logMap("Headers", req.Header)

	logRequestBody(req)
}

func logMap(title string, data map[string][]string) {
	fmt.Println(title)
	fmt.Println("---")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	for k, v := range data {
		if stringInSlice(strings.ToLower(k), excludedKeys) {
			continue
		}
		fmt.Fprintf(w, "%v\t \"%s\"\n", k, strings.Join(v, ","))
	}
	w.Flush()
	fmt.Println()
}

func logRequestBody(req *http.Request) {
	fmt.Println("Request Body")
	fmt.Println("---")

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("unable to read request body")
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	fmt.Printf("%v\n", rdr1)
	req.Body = rdr2
}

func genSeed() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return 1000 + rand.Intn(8999)
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
