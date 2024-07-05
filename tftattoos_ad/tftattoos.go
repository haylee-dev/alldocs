package tftattoos_ad

import (
	"encoding/json"
)

/*==== THE MAIN STRUCT FOR THE MONGO TATTOO IMAGES (FRONTEND) ====*/

type TattooMongo struct {
	ID             string `bson:"_id" json:"_id"`
	TattooStatus   string `bson:"tattoo_status" json:"tattoo_status"`
	TattooID       string `bson:"tattoo_id" json:"tattoo_id"`
	RandomizeList  int    `bson:"randomize_list" json:"randomize_list"`
	TattooUrl      string `bson:"tattoo_url" json:"tattoo_url"`
	ParentTattooID string `bson:"parent_tattoo_id" json:"parent_tattoo_id"`
	CreatorName    string `bson:"creator_name" json:"creator_name"`
	DateCreated    string `bson:"date_created" json:"date_created"`
	DateEntered    string `bson:"date_entered" json:"date_entered"`
	PrID           string `bson:"pr_id" json:"pr_id"`
	Title          string `bson:"title" json:"title"`
	Description    string `bson:"description" json:"description"`

	SubjectInfo SubjectInfo `bson:"subject_info" json:"subject_info"`
	SetsInfo    SetsInfo    `bson:"sets_info" json:"sets_info"`
	MockupInfo  MockupInfo  `bson:"mockup_info" json:"mockup_info"`
	ImgInfo     ImgInfo     `bson:"img_info" json:"img_info"`
	ArtInfo     ArtInfo     `bson:"art_info" json:"art_info"`
	PriceInfo   PriceInfo   `bson:"price_info" json:"price_info"`
}
type ArtInfo struct {
	ArtStylesList []string `bson:"art_styles" json:"art_styles"`
	ArtGenre      string   `bson:"art_genre" json:"art_genre"`
}
type SubjectInfo struct {
	SubjectsList []string `bson:"subjects" json:"subjects"`
}
type SetsInfo struct {
	SetNamesList []string `bson:"set_names" json:"set_names"`
	ParentName   string   `bson:"parent_name" json:"parent_name"`
	CategoryName string   `bson:"category_name" json:"category_name"`
	SubName      string   `bson:"sub_name" json:"sub_name"`
}

type MockupInfo struct {
	MockupIs     string `bson:"mockup_is" json:"mockup_is"`
	MockupGender string `bson:"mockup_gender" json:"mockup_gender"`
	MockupBody   string `bson:"mockup_body" json:"mockup_body"`
}
type ImgInfo struct {
	ImgAlt         string   `bson:"img_alt" json:"img_alt"`
	ImgDir         string   `bson:"img_dir" json:"img_dir"`
	ImgShape       string   `bson:"img_shape" json:"img_shape"`
	ImgHeight      int      `bson:"img_height" json:"img_height"`
	ImgWidth       int      `bson:"img_width" json:"img_width"`
	ImgAspectRatio string   `bson:"img_aspect_ratio" json:"img_aspect_ratio"`
	ImgColorType   string   `bson:"img_color_type" json:"img_color_type"`
	ImgColorsList  []string `bson:"img_colors_list" json:"img_colors_list"`
}
type PriceInfo struct {
	PriceUSD float64 `bson:"price_usd" json:"price_usd"`
	PriceGBP float64 `bson:"price_gbp" json:"price_gbp"`
	PriceEUR float64 `bson:"price_eur" json:"price_eur"`
	PriceCAD float64 `bson:"price_cad" json:"price_cad"`
	PriceAUD float64 `bson:"price_aud" json:"price_aud"`
}

/*==== THE MAIN STRUCT FOR THE MYSQL TATTOO IMAGES (FRONTEND) ====*/

type TattooMysql struct {
	TableID        int             `stbl:"table_id" json:"table_id"`
	TattooID       string          `stbl:"tattoo_id" json:"tattoo_id"`
	TattooUrl      string          `stbl:"tattoo_url" json:"tattoo_url"`
	TattooStatus   string          `stbl:"tattoo_status" json:"tattoo_status"`
	ParentTattooID string          `stbl:"parent_tattoo_id" json:"parent_tattoo_id"`
	CreatorName    string          `stbl:"creator_name" json:"creator_name"`
	DateCreated    string          `stbl:"date_created" json:"date_created"`
	DateEntered    string          `stbl:"date_entered" json:"date_entered"`
	MockupIs       string          `stbl:"mockup_is" json:"mockup_is"`
	MockupGender   string          `stbl:"mockup_gender" json:"mockup_gender"`
	MockupBody     string          `stbl:"mockup_body" json:"mockup_body"`
	SetNamesList   json.RawMessage `stbl:"set_names_list" json:"set_names_list"`
	ParentName     string          `stbl:"parent_name" json:"parent_name"`
	CategoryName   string          `stbl:"category_name" json:"category_name"`
	SubName        string          `stbl:"sub_name" json:"sub_name"`
	SubjectsList   json.RawMessage `stbl:"subjects_list" json:"subjects_list"`
	ArtStylesList  json.RawMessage `stbl:"art_styles_list" json:"art_styles_list"`
	ArtGenre       string          `stbl:"art_genre" json:"art_genre"`
	PrID           string          `stbl:"pr_id" json:"pr_id"`
	Title          string          `stbl:"title" json:"title"`
	Description    string          `stbl:"description" json:"description"`
	ImgAlt         string          `stbl:"img_alt" json:"img_alt"`
	ImgDir         string          `stbl:"img_dir" json:"img_dir"`
	ImgShape       string          `stbl:"img_shape" json:"img_shape"`
	ImgHeight      int             `stbl:"img_height" json:"img_height"`
	ImgWidth       int             `stbl:"img_width" json:"img_width"`
	ImgAspectRatio string          `stbl:"img_aspect_ratio" json:"img_aspect_ratio"`
	ImgColorType   string          `stbl:"img_color_type" json:"img_color_type"`
	ImgColorsList  json.RawMessage `stbl:"img_colors_list" json:"img_colors_list"`
	PriceUSD       float64         `stbl:"price_usd" json:"price_usd"`
	PriceGBP       float64         `stbl:"price_gbp" json:"price_gbp"`
	PriceEUR       float64         `stbl:"price_eur" json:"price_eur"`
	PriceCAD       float64         `stbl:"price_cad" json:"price_cad"`
	PriceAUD       float64         `stbl:"price_aud" json:"price_aud"`
}

/*==== THE MAIN STRUCT FOR THE MYSQL MYMJ ORIGINAL(X4) IMAGES (BACKEND) ====*/

type TattooImageMyMj struct {
	TableID           string `stbl:"tableID" json:"table_id"`
	Status            string `stbl:"status" json:"status"`
	CreatorName       string `stbl:"creator_name" json:"creator_name"`
	Subject           string `stbl:"subject" json:"subject"`
	IsOriginal        string `stbl:"is_original" json:"is_original"`
	DateGenerated     string `stbl:"date_generated" json:"date_generated"`
	TimeGenerated     string `stbl:"time_generated" json:"time_generated"`
	IsDownloadedLocal string `stbl:"is_downloaded_local" json:"is_downloaded_local"`
	IsElastic         string `stbl:"is_elastic" json:"is_elastic"`
	IsUpscale         string `stbl:"is_upscale" json:"is_upscale"`

	IsUploadedCdn  string `stbl:"is_uploaded_cdn" json:"is_uploaded_cdn"`
	HasWatermark   string `stbl:"has_watermark" json:"has_watermark"`
	HasSign        string `stbl:"has_sign" json:"has_sign"`
	JobId          string `stbl:"job_id" json:"job_id"`
	MessageId      string `stbl:"message_id" json:"message_id"`
	ImgTitle       string `stbl:"img_title" json:"img_title"`
	ImgDescription string `stbl:"img_description" json:"img_description"`
	ImgAlt         string `stbl:"img_alt" json:"img_alt"`
	ImgNotes       string `stbl:"img_notes" json:"img_notes"`
	ImageDir       string `stbl:"image_dir" json:"image_dir"`

	DateEntered    string `stbl:"date_entered" json:"date_entered"`
	ImageName      string `stbl:"image_name" json:"image_name"`
	ImgAspectRatio string `stbl:"img_aspect_ratio" json:"img_aspect_ratio"`
	ImgWidth       string `stbl:"img_width" json:"img_width"`
	ImgHeight      string `stbl:"img_height" json:"img_height"`
	ImgShape       string `stbl:"img_shape" json:"img_shape"`
	Content        string `stbl:"content" json:"content"`
	Prompt         string `stbl:"prompt" json:"prompt"`
	PromptId       string `stbl:"prompt_id" json:"prompt_id"`
	AiModel        string `stbl:"ai_model" json:"ai_model"`

	ParentJobId     string `stbl:"parent_jobid" json:"parent_jobid"`
	ParentName      string `stbl:"parent_name" json:"parent_name"`
	CategoryName    string `stbl:"category_name" json:"category_name"`
	SubName         string `stbl:"sub_name" json:"sub_name"`
	Style           string `stbl:"style" json:"style"`
	StyleType       string `stbl:"style_type" json:"style_type"`
	ImgColor        string `stbl:"img_color" json:"img_color"`
	Keywords        string `stbl:"keywords" json:"keywords"`
	ImageDiscordCdn string `stbl:"image_discord_cdn" json:"image_discord_cdn"`
	MockupBody      string `stbl:"mockup_body" json:"mockup_body"`
	MockupGender    string `stbl:"mockup_gender" json:"mockup_gender"`
}
