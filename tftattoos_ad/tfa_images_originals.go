package tftattoos_ad

import "time"

type TfaImagesOriginals struct {
	OriginalID        string    `gorm:"column:original_id" json:"original_id"`
	DateEntered       time.Time `gorm:"column:date_entered" json:"date_entered"`
	CreatorName       string    `gorm:"column:creator_name" json:"creator_name"`
	APIService        string    `gorm:"column:api_service" json:"api_service"`
	JobId             string    `gorm:"column:job_id" json:"job_id"`
	MessageId         string    `gorm:"column:message_id" json:"message_id"`
	Status            string    `gorm:"column:status" json:"status"`
	ImgFileStatus     string    `gorm:"column:img_file_status" json:"img_file_status"`
	ImgArtStyle       string    `gorm:"column:img_art_style" json:"img_art_style"`
	ImgSetsList       string    `gorm:"column:img_sets_list" json:"img_sets_list"`
	ImgSubject        string    `gorm:"column:img_subject" json:"img_subject"`
	ImgSubjectContext string    `gorm:"column:img_subject_context" json:"img_subject_context"`
	DateGenerated     string    `gorm:"column:date_generated" json:"date_generated"`
	TimeGenerated     string    `gorm:"column:time_generated" json:"time_generated"`
	ImgDir            string    `gorm:"column:img_dir" json:"img_dir"`
	ImgKeywords       string    `gorm:"column:img_keywords" json:"img_keywords"`
	ImgDiscordCDN     string    `gorm:"column:img_discord_cdn" json:"img_discord_cdn"`
	ImgColorType      string    `gorm:"column:img_color_type" json:"img_color_type"`
	ImgColorsList     string    `gorm:"column:img_colors_list" json:"img_colors_list"`
	ImgFilename       string    `gorm:"column:img_filename" json:"img_filename"`
	ImgFilenameCDN    string    `gorm:"column:img_filename_cdn" json:"img_filename_cdn"`
	ImgAspectRatio    string    `gorm:"column:img_aspect_ratio" json:"img_aspect_ratio"`
	ImgWidth          string    `gorm:"column:img_width" json:"img_width"`
	ImgHeight         string    `gorm:"column:img_height" json:"img_height"`
	ImgShape          string    `gorm:"column:img_shape" json:"img_shape"`
	Prompt            string    `gorm:"column:prompt" json:"prompt"`
	PromptId          string    `gorm:"column:prompt_id" json:"prompt_id"`
	AiModel           string    `gorm:"column:ai_model" json:"ai_model"`
	MockupBody        string    `gorm:"column:mockup_body" json:"mockup_body"`
	MockupGender      string    `gorm:"column:mockup_gender" json:"mockup_gender"`
}
