package models

type FinalWebpushDto struct {
	AppID string `json:"appId"`
	CampaignName string `json:"campaignName"`
	RequestType string `json:"requestType"`
	TargetPlatform []string `json:"targetPlatform"`
	TargetAudience string `json:"targetAudience"`
	Signature string `json:"signature"`
	Payload struct {
		      WEB struct {
				  Message string `json:"message"`
				  Title string `json:"title"`
				  RedirectURL string `json:"redirectURL"`
				  IconURL string `json:"iconURL"`
				  Fallback struct {
					  } `json:"fallback"`
			  } `json:"WEB"`
	      } `json:"payload"`
	CampaignDelivery struct {
		      Type string `json:"type"`
	      } `json:"campaignDelivery"`
}
