package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getTMs() (r string) {
	url := "https://api.testcleveraim.com/v1.0/tm/overview"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(`{""}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer b9cdd71724759787840dcb5214a44101d66448982684ed87d334b6ed2615d706f57e1254fa635bcdbe35273c8c532d09716934951d2bd1b9ce57b4eecb775094")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r = string(body)
	return
}

func getItems() (r string) {
	url := "https://clevermonitor.sharepoint.com/_api/web/lists/GetByTitle('EmptyFolder')/items"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(`{""}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Cookie", "FeatureOverrides_enableFeatures=; FeatureOverrides_disableFeatures=; splnu=0; WSS_FullScreenMode=false; WOPISessionContext=https%3A%2F%2Fclevermonitor%2Esharepoint%2Ecom%2FA%5Fteam%2520Docs%2FForms%2FAllItems%2Easpx; rtFa=52PM85N6ipssRIRz4wUrnEFPHArkAouNSy4bjYDzQ4gmOUI3NzM0RjQtMTk5Mi00NUZDLUEzMzctOTgxODIzMDc5MTRBYER0Jd2vIIstZ83kByNGSylBkZCcAm1Z/H2zGTh8VJKXhozK/tv8j8u2J/iidqbL+iXgU+hG/qTwllb667ZOMfhxjnUsdplGcrSY1NdtfDDLV1Hsr3lxVltRc6noTWUjFTJWBwcdpWKlZQj/06iSnrks1mIB4+wRw8aRfwm1ly6bQI1PD+uUNenT9mOk6eWu6R8AnlTE8zx3OuK5k8/M5ngkhVYLtBdItobWpf6a3zDUfGVp1zMVvVVJkVkh6HgeuIsbkJJup8QmawVW84YpejfyONeUaFir6InJWDqDH7W8JRz2dqNkzdHpBXcTjeS8Fp2Trhjd99EiQUSrWQ/VW0UAAAA=; FedAuth=77u/PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48U1A+VjUsMGguZnxtZW1iZXJzaGlwfDEwMDNiZmZkYWRmMmI2YmVAbGl2ZS5jb20sMCMuZnxtZW1iZXJzaGlwfHZhY2xhdi5oYW51c0BjbGV2ZXJtb25pdG9yLmNvbSwxMzE4MzAzMDI4MjAwMDAwMDAsMTMxODI5NTcwOTMwMDAwMDAwLDEzMTk1Mzg5NzQyMjE2NDI3MSwwLjAuMC4wLDMsOWI3NzM0ZjQtMTk5Mi00NWZjLWEzMzctOTgxODIzMDc5MTRhLCxWMiExMDAzQkZGREFERjJCNkJFITEzMTgzMDMwMjgyLGI1MjNjMTllLWIwMDEtODAwMC00OGM4LTAwNjRjZmM1YjdmMCxiNTIzYzE5ZS1iMDAxLTgwMDAtNDhjOC0wMDY0Y2ZjNWI3ZjAsLDAsMTMxOTQ5NjEzNDIxNTM5MjU5LDEzMTk1MjE2OTQyMTUzOTI1OSwsbjcvMnZCdGRPeXJ2ZVNici9wWEFrc0loTFVQak1EOWlXcFIxWGlmSkxIbDVuWUE4RzBaY0ljLzUxUGVpSTczeWlXa2Y5UE16blRjQ241UFZIWnBpekpIYUNIWENSSzFpcnBvOVFLYk16ZUs3MkZoVHdPT0VwRGtWaktGNjAwL0J1bzFxQlhSbjQxY1RraHl2RmpKVlFmNUdCb1BjUXhqTUJUNDVFdkYyWHBvbmE5bjFZdVMxKzVBRGRoa2Q1RzNzeXVZcFQ3V0RWWUdwQ3FsY2Z5OGgvKzF4Q2FML2tEQ2doVDd6QnZ4RjBvbU05dnUrZExPa25Td3AvbUxXemVUQng2aU90Y0pFQ1J4UWw0Y0NHVzFoTVR4Uk1ER3VMZUlneTJ5Z3pIQ1AzcWJCaE5COGI1c05lRmNvRnlieDlzQThieVUxZlRQRXN1TUdIMm5BRmw0YTFnPT08L1NQPg==; CCSInfo=MjEuMDIuMjAxOSAxMjowNzoyMTLglOWir4AaJ6pokMgFKJAF6grGLv92m0KbaU66PC/9oMrOhzeDhNGW+wzoJwcf0cL6zkTkmLItbqzf3yMkhowcFROU7TDBFFtIFev6jYSXitSKUIcXlX5CZkhbXGjRRGBaerraQnySwLk3pcwoxNiWqSaMJWhmEnU9aipSTg5o8CjUp3aAA9qSpZpdTaEHE74pGd/avsMME+D8IxF2wyp/cL3vQ0bbc0tFb2tfwWIYXr8lyloN6JjTO6CDyhRl7wu4mfsYmyQufpjkTywgF7CHbL/6aN/siu/zKQWf/etebejAby0CTqYJV99e12LBRJWBQ0eHJ1wsVovAp7J+a7ETAAAA; SPWorkLoadAttribution=Url=https://clevermonitor.sharepoint.com/SitePages/Home.aspx&AppTitle=RenderClientSideBasePage")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r = string(body)
	return
}

func addItem(b string) {
	url := "https://clevermonitor.sharepoint.com/_api/web/lists/GetByTitle('EmptyFolder')/items"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(b)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Cookie", "FeatureOverrides_enableFeatures=; FeatureOverrides_disableFeatures=; splnu=0; WSS_FullScreenMode=false; WOPISessionContext=https%3A%2F%2Fclevermonitor%2Esharepoint%2Ecom%2FA%5Fteam%2520Docs%2FForms%2FAllItems%2Easpx; rtFa=52PM85N6ipssRIRz4wUrnEFPHArkAouNSy4bjYDzQ4gmOUI3NzM0RjQtMTk5Mi00NUZDLUEzMzctOTgxODIzMDc5MTRBYER0Jd2vIIstZ83kByNGSylBkZCcAm1Z/H2zGTh8VJKXhozK/tv8j8u2J/iidqbL+iXgU+hG/qTwllb667ZOMfhxjnUsdplGcrSY1NdtfDDLV1Hsr3lxVltRc6noTWUjFTJWBwcdpWKlZQj/06iSnrks1mIB4+wRw8aRfwm1ly6bQI1PD+uUNenT9mOk6eWu6R8AnlTE8zx3OuK5k8/M5ngkhVYLtBdItobWpf6a3zDUfGVp1zMVvVVJkVkh6HgeuIsbkJJup8QmawVW84YpejfyONeUaFir6InJWDqDH7W8JRz2dqNkzdHpBXcTjeS8Fp2Trhjd99EiQUSrWQ/VW0UAAAA=; CCSInfo=MjEuMDIuMjAxOSAxMjowNzoyMTLglOWir4AaJ6pokMgFKJAF6grGLv92m0KbaU66PC/9oMrOhzeDhNGW+wzoJwcf0cL6zkTkmLItbqzf3yMkhowcFROU7TDBFFtIFev6jYSXitSKUIcXlX5CZkhbXGjRRGBaerraQnySwLk3pcwoxNiWqSaMJWhmEnU9aipSTg5o8CjUp3aAA9qSpZpdTaEHE74pGd/avsMME+D8IxF2wyp/cL3vQ0bbc0tFb2tfwWIYXr8lyloN6JjTO6CDyhRl7wu4mfsYmyQufpjkTywgF7CHbL/6aN/siu/zKQWf/etebejAby0CTqYJV99e12LBRJWBQ0eHJ1wsVovAp7J+a7ETAAAA; FedAuth=77u/PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48U1A+VjUsMGguZnxtZW1iZXJzaGlwfDEwMDNiZmZkYWRmMmI2YmVAbGl2ZS5jb20sMCMuZnxtZW1iZXJzaGlwfHZhY2xhdi5oYW51c0BjbGV2ZXJtb25pdG9yLmNvbSwxMzE4MzAzMDI4MjAwMDAwMDAsMTMxODI5NTcwOTMwMDAwMDAwLDEzMTk1NDc5MTg0NTM2MjU1OSwwLjAuMC4wLDMsOWI3NzM0ZjQtMTk5Mi00NWZjLWEzMzctOTgxODIzMDc5MTRhLCxWMiExMDAzQkZGREFERjJCNkJFITEzMTgzMDMwMjgyLGI1MjNjMTllLWIwMDEtODAwMC00OGM4LTAwNjRjZmM1YjdmMCwwMTc5YzE5ZS04MDg4LTgwMDAtZDg2ZS04ODU4NDNhNjUyNDIsLDAsMTMxOTQ5NjEzNDIxNTM5MjU5LDEzMTk1MjE2OTQyMTUzOTI1OSwsZFRHN1dDRDJabFBKQTdDMW5GcFA0SVU1bjhuTjA4VllhdGhjN3did2tEa1NWam9VMFRuTEdzcGJML0s2d2piakFKRHF6RFg3NjdFTldVRWR1LzFkejc0U3FjbHVMOTlSWlB4U0N6eDNtUUUrcStmRExlRHFDcVVOREt3VGVabll5MDdZVGROV3MySU1oYk9wMTZpVThDQXB0UmlvUDU4M09obksrT0QzQTlZay9mUEdodkMrMDZ0WXQ0WGNSMG5iVGVZaXhXblpGSGVQODdUbEs0bnREZ3FIbXNMT0I4eStXRlZ6aXNKSStrMHMzaGQwUVFHZE9xVlNwS2VpYitScS9kSHhuK2FqK0dndGVaM3hvWUpXalRtMDFXVGlYaGthcU1XOUo2R01XWVYxSUVKcFF2TUVyZmlGTHR3WDhqTE5mUzNCamJSMFFRbHZhbkl3OUNqZHB3PT08L1NQPg==; odbn=1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-RequestDigest", "0xC0558782E70B16AF51F4755FC7CEE5A902EECAB4345CE0D61587E809FB4F1D634F5924F165D87BC72E0EDAB81808C27488D174E7C9F933A9760B86FE57530001,19 Feb 2019 14:08:28 -0000")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func updateItem(b string, itemID string) {
	url := "https://clevermonitor.sharepoint.com/_api/web/lists/GetByTitle('EmptyFolder')/items" + itemID
	fmt.Println("URL:>", url)
	var jsonStr = []byte(b)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Cookie", "FeatureOverrides_enableFeatures=; FeatureOverrides_disableFeatures=; splnu=0; WSS_FullScreenMode=false; WOPISessionContext=https%3A%2F%2Fclevermonitor%2Esharepoint%2Ecom%2FA%5Fteam%2520Docs%2FForms%2FAllItems%2Easpx; rtFa=52PM85N6ipssRIRz4wUrnEFPHArkAouNSy4bjYDzQ4gmOUI3NzM0RjQtMTk5Mi00NUZDLUEzMzctOTgxODIzMDc5MTRBYER0Jd2vIIstZ83kByNGSylBkZCcAm1Z/H2zGTh8VJKXhozK/tv8j8u2J/iidqbL+iXgU+hG/qTwllb667ZOMfhxjnUsdplGcrSY1NdtfDDLV1Hsr3lxVltRc6noTWUjFTJWBwcdpWKlZQj/06iSnrks1mIB4+wRw8aRfwm1ly6bQI1PD+uUNenT9mOk6eWu6R8AnlTE8zx3OuK5k8/M5ngkhVYLtBdItobWpf6a3zDUfGVp1zMVvVVJkVkh6HgeuIsbkJJup8QmawVW84YpejfyONeUaFir6InJWDqDH7W8JRz2dqNkzdHpBXcTjeS8Fp2Trhjd99EiQUSrWQ/VW0UAAAA=; CCSInfo=MjEuMDIuMjAxOSAxMjowNzoyMTLglOWir4AaJ6pokMgFKJAF6grGLv92m0KbaU66PC/9oMrOhzeDhNGW+wzoJwcf0cL6zkTkmLItbqzf3yMkhowcFROU7TDBFFtIFev6jYSXitSKUIcXlX5CZkhbXGjRRGBaerraQnySwLk3pcwoxNiWqSaMJWhmEnU9aipSTg5o8CjUp3aAA9qSpZpdTaEHE74pGd/avsMME+D8IxF2wyp/cL3vQ0bbc0tFb2tfwWIYXr8lyloN6JjTO6CDyhRl7wu4mfsYmyQufpjkTywgF7CHbL/6aN/siu/zKQWf/etebejAby0CTqYJV99e12LBRJWBQ0eHJ1wsVovAp7J+a7ETAAAA; FedAuth=77u/PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48U1A+VjUsMGguZnxtZW1iZXJzaGlwfDEwMDNiZmZkYWRmMmI2YmVAbGl2ZS5jb20sMCMuZnxtZW1iZXJzaGlwfHZhY2xhdi5oYW51c0BjbGV2ZXJtb25pdG9yLmNvbSwxMzE4MzAzMDI4MjAwMDAwMDAsMTMxODI5NTcwOTMwMDAwMDAwLDEzMTk1NDc5MTg0NTM2MjU1OSwwLjAuMC4wLDMsOWI3NzM0ZjQtMTk5Mi00NWZjLWEzMzctOTgxODIzMDc5MTRhLCxWMiExMDAzQkZGREFERjJCNkJFITEzMTgzMDMwMjgyLGI1MjNjMTllLWIwMDEtODAwMC00OGM4LTAwNjRjZmM1YjdmMCwwMTc5YzE5ZS04MDg4LTgwMDAtZDg2ZS04ODU4NDNhNjUyNDIsLDAsMTMxOTQ5NjEzNDIxNTM5MjU5LDEzMTk1MjE2OTQyMTUzOTI1OSwsZFRHN1dDRDJabFBKQTdDMW5GcFA0SVU1bjhuTjA4VllhdGhjN3did2tEa1NWam9VMFRuTEdzcGJML0s2d2piakFKRHF6RFg3NjdFTldVRWR1LzFkejc0U3FjbHVMOTlSWlB4U0N6eDNtUUUrcStmRExlRHFDcVVOREt3VGVabll5MDdZVGROV3MySU1oYk9wMTZpVThDQXB0UmlvUDU4M09obksrT0QzQTlZay9mUEdodkMrMDZ0WXQ0WGNSMG5iVGVZaXhXblpGSGVQODdUbEs0bnREZ3FIbXNMT0I4eStXRlZ6aXNKSStrMHMzaGQwUVFHZE9xVlNwS2VpYitScS9kSHhuK2FqK0dndGVaM3hvWUpXalRtMDFXVGlYaGthcU1XOUo2R01XWVYxSUVKcFF2TUVyZmlGTHR3WDhqTE5mUzNCamJSMFFRbHZhbkl3OUNqZHB3PT08L1NQPg==; odbn=1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-RequestDigest", "0xC0558782E70B16AF51F4755FC7CEE5A902EECAB4345CE0D61587E809FB4F1D634F5924F165D87BC72E0EDAB81808C27488D174E7C9F933A9760B86FE57530001,19 Feb 2019 14:08:28 -0000")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func decodeInput(b string, shpItems map[string][]string) {
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	jsonData, err := json.Marshal(result["data"])
	if err != nil {
		log.Fatal(err)
	}
	type Message struct {
		AttachmentCount  int    `json:"attachmentCount"`
		ID               string `json:"id"`
		LastModifiedDate string `json:"lastModifiedDate"`
		MessageID        string `json:"messageId"`
		ReceiveDate      string `json:"receiveDate"`
		SenderEmail      string `json:"senderEmail"`
		SenderName       string `json:"senderName"`
		SentDate         string `json:"sentDate"`
		Status           string `json:"status"`
		Subject          string `json:"subject"`
		SubscriberEmail  string `json:"subscriberEmail"`
	}
	var Messages []Message
	er := json.Unmarshal(jsonData, &Messages)
	if er != nil {
		log.Fatal(er)
	}
	for index := range Messages {
		if val, ok := shpItems[Messages[index].ID]; ok {
			fmt.Println("Current state: " + val[0])
			if Messages[index].Status == shpItems[Messages[index].ID][0] {
				fmt.Println("Already in list")
				requestBody := `{"Title": "` + Messages[index].Subject + `","GUID0": "` + Messages[index].ID + `","Status": "` + Messages[index].Status + `"}`
				fmt.Println(requestBody)
			} else {
				requestBody := `{"Status": "` + Messages[index].Status + `","LastChange": "` + Messages[index].LastModifiedDate + `"}`
				fmt.Println(requestBody)
				//updateItem(requestBody,shpItems[Messages[index].ID][1])
			}
		} else {
			requestBody := `{"Title": "` + Messages[index].Subject + `","GUID0": "` + Messages[index].ID + `","Status": "` + Messages[index].Status + `","LastChange": "` + Messages[index].LastModifiedDate + `"}`
			fmt.Println(requestBody)
			//addItem(requestBody)
		}
	}
}

func decodeShpInput(b string) map[string][]string {
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	jsonData, err := json.Marshal(result["value"])
	if err != nil {
		log.Fatal(err)
	}
	type Message struct {
		Otype                  string `json:"odata.type"`
		Oid                    string `json:"odata.id"`
		Oetag                  string `json:"odata.etag"`
		Oedit                  string `json:"odata.editLink"`
		FileSystemObjectType   int    `json:"FileSystemObjectType"`
		Sid                    int    `json:"Id"`
		ServerRedirectedEmbed  string `json:"ServerRedirectedEmbedUri"`
		ServerRedirectedEmbedU string `json:"ServerRedirectedEmbedUrl"`
		ContentTypeID          string `json:"ContentTypeId"`
		Subject                string `json:"Title"`
		ComplianceAssetID      string `json:"ComplianceAssetId"`
		GUID                   string `json:"GUID0"`
		Status                 string `json:"Status"`
		LastCange              string `json:"LastCange"`
		ID                     int    `json:"ID"`
		Modified               string `json:"Modified"`
		Created                string `json:"Created"`
		AuthorID               int    `json:"AuthorId"`
		EditorID               int    `json:"EditorId"`
		OUI                    string `json:"OData__UIVersionString"`
		Attachments            bool   `json:"Attachments"`
		OGUID                  string `json:"GUID"`
	}
	var Messages []Message
	er := json.Unmarshal(jsonData, &Messages)
	if er != nil {
		log.Fatal(er)
	}
	shpitems := make(map[string][]string)
	for index := range Messages {
		fmt.Println("---------------------------------")
		fmt.Println(Messages[index].GUID + " > " + Messages[index].Status)
		shpitems[Messages[index].GUID][0] = Messages[index].Status
		shpitems[Messages[index].GUID][1] = string(Messages[index].ID)
	}
	return shpitems
}

func main() {
	shpData := getItems()
	shpItems := decodeShpInput(shpData)
	fmt.Println("---------------------------------")
	tmsData := getTMs()
	decodeInput(tmsData, shpItems)
	fmt.Println("---------------------------------")
	//addItem()

}
