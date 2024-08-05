package tftattoos_ad

import "time"

type TfaImagesOriginals struct {
	OriginalID        string    `json:"original_id"`
	DateEntered       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"date_entered"`
	CreatorName       string    `json:"creator_name"`
	APIService        string    `json:"api_service"`
	JobId             string    `json:"job_id"`
	MessageId         string    `json:"message_id"`
	Status            string    `json:"status"`
	ImgFileStatus     string    `json:"img_file_status"`
	ImgArtStyle       string    `json:"img_art_style"`
	ImgSetsList       string    `json:"img_sets_list"`
	ImgSubject        string    `json:"img_subject"`
	ImgSubjectContext string    `json:"img_subject_context"`
	DateGenerated     string    `json:"date_generated"`
	TimeGenerated     string    `json:"time_generated"`
	ImgDir            string    `json:"img_dir"`
	ImgKeywords       string    `json:"img_keywords"`
	ImgDiscordCDN     string    `json:"img_discord_cdn"`
	ImgColorType      string    `json:"img_color_type"`
	ImgColorsList     string    `json:"img_colors_list"`
	ImgFilename       string    `json:"img_filename"`
	ImgFilenameCDN    string    `json:"img_filename_cdn"`
	ImgAspectRatio    string    `json:"img_aspect_ratio"`
	ImgWidth          string    `json:"img_width"`
	ImgHeight         string    `json:"img_height"`
	ImgShape          string    `json:"img_shape"`
	Prompt            string    `json:"prompt"`
	PromptId          string    `json:"prompt_id"`
	AiModel           string    `json:"ai_model"`
	MockupBody        string    `json:"mockup_body"`
	MockupGender      string    `json:"mockup_gender"`
}
