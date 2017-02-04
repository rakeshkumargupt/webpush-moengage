package webpush

import (
	"io/ioutil"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"bytes"
	"net/http"
	"github.com/rakeshkumargupt/models"
)

const (
	JSON_CONTENT_TYPE = "application/json"

	// Endpoint URL of new version
	ENDPOINT_URL = "https://pushapi.moengage.com/v2/transaction/sendpush"

	//Moengage APP_ID
	APP_ID = "AAAAAAAAAAAAAAAAAAA"

	// Moengage API_SECRET_KEY
	API_SECRET_KEY = "AAAAAAAAA"
)


func WebPush(inputWebpush models.UiWebpushModel) error {

	client := &http.Client{}
	payloadVar := GeneratePayload(inputWebpush)
	payload, _ := json.Marshal(&payloadVar)
	req, _ := http.NewRequest("POST", ENDPOINT_URL, bytes.NewReader(payload))
	req.Header.Set("Content-Type", JSON_CONTENT_TYPE)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return nil
}

func GeneratePayload(inputWebpushVar models.UiWebpushModel) models.FinalWebpushDto{

	var WPMVar models.FinalWebpushDto
	WPMVar.AppID = APP_ID
	WPMVar.CampaignName = "aaaaaaa" //Name of Campaign
	sig := WPMVar.AppID + "|" + WPMVar.CampaignName + "|" + API_SECRET_KEY
	byteArr := sha256.Sum256([]byte(sig))
	WPMVar.Signature = hex.EncodeToString(byteArr[:])
	var target []string
	target = append(target, "WEB") //Target platfrom like:WEB,ANDROID,IOS,WINDOW
	WPMVar.TargetPlatform = target
	WPMVar.Payload.WEB.Title = inputWebpushVar.Title
	WPMVar.Payload.WEB.Message = inputWebpushVar.Message
	WPMVar.Payload.WEB.RedirectURL = inputWebpushVar.RedirectURL
	WPMVar.Payload.WEB.IconURL = inputWebpushVar.IconURL
	WPMVar.Payload.WEB.Fallback = struct {
	}{}
	WPMVar.CampaignDelivery.Type = inputWebpushVar.Type
	WPMVar.RequestType = "push" //Request Type push immediate or sheduled.
	WPMVar.TargetAudience = inputWebpushVar.TargetAudience

	return WPMVar
}
