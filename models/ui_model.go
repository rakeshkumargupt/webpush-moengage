package models


type UiWebpushModel struct {
	TargetAudience string `json:"targetAudience" form:"targetAudience" binding:"required"`
	Message string `json:"message" form:"message" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	RedirectURL string `json:"redirectURL" form:"redirectURL"`
	IconURL string `json:"iconURL" form:"iconURL"`
	Type string `json:"type" form:"type" binding:"required"`
}