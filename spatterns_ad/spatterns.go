package spatterns_ad

type PatternsMongo struct {
	ID               string `bson:"_id" json:"_id"`
	PatternID        string `bson:"pattern_id" json:"pattern_id"`
	ParentPatternID  string `bson:"parent_pattern_id" json:"parent_pattern_id"`
	PatternURL       string `bson:"pattern_url" json:"pattern_url"`
	DateCreated      string `bson:"date_created" json:"date_created"`
	DateEntered      string `bson:"date_entered" json:"date_entered"`
	PrID             string `bson:"pr_id" json:"pr_id"`
	PatternStatus    string `bson:"pattern_status" json:"pattern_status"`
	Title            string `bson:"title" json:"title"`
	DescriptionBrief string `bson:"description_brief" json:"description_brief"`
	Description      string `bson:"description" json:"description"`
	RandomizeList    int    `bson:"randomize_list" json:"randomize_list"`

	PatternInfo PatternInfo `bson:"pattern_info" json:"pattern_info"`
	SetsInfo    SetsInfo    `bson:"sets_info" json:"sets_info"`
	ArtistInfo  ArtistInfo  `bson:"artist_info" json:"artist_info"`
	ImageInfo   ImageInfo   `bson:"image_info" json:"image_info"`
	ColorsInfo  ColorsInfo  `bson:"colors_info" json:"colors_info"`
	PriceInfo   PriceInfo   `bson:"price_info" json:"price_info"`
	MockupInfo  MockupInfo  `bson:"mockup_info" json:"mockup_info"`
	LicenseInfo LicenseInfo `bson:"license_info" json:"license_info"`
	UsageInfo   UsageInfo   `bson:"usage_info" json:"usage_info"`
}

type PatternInfo struct {
	PatternTitle            string   `bson:"pattern_title" json:"pattern_title"`
	PatternDescriptionBrief string   `bson:"pattern_description_brief" json:"pattern_description_brief"`
	PatternDescription      string   `bson:"pattern_description" json:"pattern_description"`
	PatternName             string   `bson:"pattern_name" json:"pattern_name"`
	PatternStyle            string   `bson:"pattern_style,omitempty" json:"pattern_style,omitempty"`             //retro
	PatternTheme            string   `bson:"pattern_theme,omitempty" json:"pattern_theme,omitempty"`             //theme like 'nautical'
	PatternKeywords         []string `bson:"pattern_keywords,omitempty" json:"pattern_keywords,omitempty"`       //mosaic, boho, bold, earthy
	PatternContents         []string `bson:"pattern_content" json:"pattern_content"`                             //lemons, birds, nature, roses (picked up as a sub)
	Gender                  string   `bson:"gender" json:"gender"`                                               //masculine, feminine
	PatternSeason           string   `bson:"pattern_season,omitempty" json:"pattern_season,omitempty"`           //autumn, winter
	PatternCelebration      string   `bson:"pattern_celebration,omitempty" json:"pattern_celebration,omitempty"` //anniversary, birthday
	PatternFestival         string   `bson:"pattern_festival,omitempty" json:"pattern_festival,omitempty"`
}

type UsageInfo struct {
	KidsWear          bool     `bson:"kids_wear" json:"kids_wear"`
	MasculineWear     bool     `bson:"masculine_wear" json:"masculine_wear"`
	FeminineWear      bool     `bson:"feminine_wear" json:"feminine_wear"`
	StationaryItems   bool     `bson:"stationary_items" json:"stationary_items"`
	TechItems         bool     `bson:"tech_items" json:"tech_items"`
	HomewareItems     bool     `bson:"homeware_items" json:"homeware_items"`
	Accessories       bool     `bson:"accessories" json:"accessories"`
	Footwear          bool     `bson:"footwear" json:"footwear"`
	OutdoorAndTravel  bool     `bson:"outdoor_and_travel" json:"outdoor_and_travel"`
	KitchenwareItems  bool     `bson:"kitchenware_items" json:"kitchenware_items"`
	PetAccessories    bool     `bson:"pet_accessories" json:"pet_accessories"`
	BeautyAndPersonal bool     `bson:"beauty_and_personal" json:"beauty_and_personal"`
	Automotive        bool     `bson:"automotive" json:"automotive"`
	PatternUsage      []string `bson:"pattern_usage,omitempty" json:"pattern_usage,omitempty"`
}

type SetsInfo struct {
	ParentName    string   `bson:"parent_name" json:"parent_name"`                           //main aesthetic
	CategoryNames []string `bson:"category_names,omitempty" json:"category_names,omitempty"` //list of other categories
}

type ArtistInfo struct {
	ArtistName string `bson:"artist_name" json:"artist_name"`
}

type ImageInfo struct {
	ImgDir    string `bson:"img_dir" json:"img_dir"`
	ImgAlt    string `bson:"img_alt" json:"img_alt"`
	ImgShape  string `bson:"img_shape" json:"img_shape"`
	ImgHeight int    `bson:"img_height" json:"img_height"`
	ImgWidth  int    `bson:"img_width" json:"img_width"`
}

type ColorsInfo struct {
	ColorsHex   []string `bson:"colors_hex,omitempty" json:"colors_hex,omitempty"`
	ColorsList  []string `bson:"colors_list,omitempty" json:"colors_list,omitempty"`
	ColorsClass string   `bson:"colors_class,omitempty" json:"colors_class,omitempty"`
}
type PriceInfo struct {
	PriceUnlimitedUsd float64 `bson:"price_unlimited_usd" json:"price_unlimited_usd"`
	PriceScreenUsd    float64 `bson:"price_screen_usd" json:"price_screen_usd"`
	PricePrintUsd     float64 `bson:"price_print_usd" json:"price_print_usd"`
	PriceTier         string  `bson:"price_tier" json:"price_tier"`
}
type LicenseInfo struct {
	LicenseType string `bson:"license_type" json:"license_type"`
}

type MockupInfo struct {
	Mockup1 string `bson:"mockup_1" json:"mockup_1"`
	Mockup2 string `bson:"mockup_2" json:"mockup_2"`
	Mockup3 string `bson:"mockup_3" json:"mockup_3"`
	Mockup4 string `bson:"mockup_4" json:"mockup_4"`
	Mockup5 string `bson:"mockup_5" json:"mockup_5"`
	Mockup6 string `bson:"mockup_6" json:"mockup_6"`
}
