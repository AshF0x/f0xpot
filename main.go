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
	"time"

	"github.com/gliderlabs/ssh"
)

const (
	sshPort = ":22" //might want to change this as your real SSH daemon is using this port
)

type geoIP struct {
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

func main() {

	listenErr := ssh.ListenAndServe(sshPort, nil, ssh.PasswordAuth(ConnectionHandler)) //On connection call COnnectionHandler
	if listenErr != nil {
		//handle failures
		log.Println("failed to start ssh server")
		log.Fatalln(listenErr.Error())
		os.Exit(1)
	}

}

//ConnectionHandler when an attack connects print their detaails to stdout and close the connection.
func ConnectionHandler(ctx ssh.Context, pass string) bool {
	//strip off the port number at end of ip ip:125120 -> ip
	ip := ctx.RemoteAddr().String()[:strings.IndexByte(ctx.RemoteAddr().String(), ':')]
	data, err := requestLocation(ip)
	if err != nil {
		//now do real logging now :^)
		log.Println(err)
	}
	log.Printf("%s - %s:%s - %s", ip, ctx.User(), pass, data.Country)
	writeInflux(ip, ctx.User(), pass, data.Country, data.City) // time already handled
	return false
}

// RequestLocation returns information on an IP address data from IP-API.com
func requestLocation(ipAddress string) (gipresult geoIP, err error) {
	api := "http://ip-api.com/json"

	request := fmt.Sprintf("%s/%s", api, ipAddress)
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

	var result geoIP
	err = json.Unmarshal(body, &result)
	if err != nil {
		return gipresult, errors.New("failed to parse JSON")
	}

	if result.Status == "fail" {
		return gipresult, errors.New(result.Message)
	}

	return result, nil
}

func writeInflux(ip string, username string, password string, country string, city string) {
	file, err := os.OpenFile("log.influx", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)

	_, err = file.WriteString(fmt.Sprintf("IP=%q,Username=%q,Password=%q,Country=%q,City=%q %d\n",
		ip,
		username,
		password,
		country,
		city,
		time.Now().UnixNano()))

	if err != nil {
		log.Println(err.Error())
	}

	file.Close()
}

func checkError(err error) {
	if err != nil {
		panic(1)
	}
}

func createKeyValue(key string, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}
