package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gliderlabs/ssh"
)

const (
	sshPort = ":22" //might want to change this as your real SSH daemon is using this port
)

func main() {

	listenErr := ssh.ListenAndServe(sshPort, nil, ssh.PasswordAuth(ConnectionHandler)) //On connection call COnnectionHandler
	if listenErr != nil {
		//handle failures
		log.Println("failed to start ssh server")
		log.Fatalln(listenErr.Error())
		os.Exit(1)
	}

}

type GeoIP struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Message     string  `json:"message"`
}

//ConnectionHandler when an attack connects print their detaails to stdout and close the connection.
func ConnectionHandler(ctx ssh.Context, pass string) bool {
	data, err := RequestLocation(ctx.RemoteAddr().String()[:strings.IndexByte(ctx.RemoteAddr().String(), ':')]) //strip off the port number at end of ip ip:125120 -> ip
	if err != nil {
		//now do real logging now :^)
		log.Println(err)
	}
	log.Printf("%s - %s: '%s' - %s", ctx.RemoteAddr(), ctx.User(), pass, data.Country)
	return false
}

// RequestLocation returns information on an IP address data from IP-API.com
func RequestLocation(ipAddress string) (gipresult GeoIP, err error) {
	api := "http://ip-api.com/json/"

	request := fmt.Sprintf("%s%s", api, ipAddress)
	response, err := http.Get(request)
	if err != nil {
		return gipresult, errors.New("Error")
	}

	if response.StatusCode != 200 {
		return gipresult, errors.New("Status code != 200")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return gipresult, errors.New("failed to read response body")
	}

	var result GeoIP
	err = json.Unmarshal(body, &result)
	if err != nil {
		return gipresult, errors.New("failed to parse JSON")
	}

	if result.Status == "fail" {
		return gipresult, errors.New(result.Message)
	}

	return result, nil
}
