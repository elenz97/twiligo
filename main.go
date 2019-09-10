package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
	Play    string   `xml:",omitempty"`
}

func main() {
	fmt.Println("Starting Twiligo Server...")
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/makecall", makeCall)
	http.ListenAndServe(":2255", nil)
}

func twiml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request for /twiml")
	twiml := TwiML{Play: os.Getenv("TWILIO_MP3_URL")}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func makeCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request for /makecall")
	var (
		accountSid     = os.Getenv("TWILIO_ACCOUNT_SID")
		authToken      = os.Getenv("TWILIO_ACCOUNT_TOKEN")
		receiver       = os.Getenv("TWILIO_RECEIVER")
		sender         = os.Getenv("TWILIO_SENDER")
		webEndpointURL = os.Getenv("TWILIO_ENDPOINT_URL")
	)

	if accountSid == "" {
		log.Fatal("'TWILIO_ACCOUNT_SID' environment variable needs to be set!")
	} else if authToken == "" {
		log.Fatal("'TWILIO_ACCOUNT_TOKEN' environment variable needs to be set!")
	} else if receiver == "" {
		log.Fatal("'TWILIO_RECEIVER' environment variable needs to be set!")
	} else if sender == "" {
		log.Fatal("'TWILIO_SENDER' environment variable needs to be set!")
	} else if webEndpointURL == "" {
		log.Fatal("'TWILIO_ENDPOINT_URL' environment variable needs to be set!")
	}
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Calls.json"
	fmt.Print(urlStr)
	v := url.Values{}
	v.Set("To", "+"+receiver)
	v.Set("From", "+"+sender)
	v.Set("Url", webEndpointURL+":2255/twiml")
	rb := *strings.NewReader(v.Encode())
	fmt.Println(&rb)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(req)
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println("Response status from Twilio API:", resp.Status, "\n")
		w.Write([]byte("Something went wrong trying to POST to the Twilio API\nPlease check your credentials specified in .env,\nor refer to the logs\n"))
	}
}
