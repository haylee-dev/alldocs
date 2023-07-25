package alldocs

/* MAIN PRODUCT LISTING DOC FOR BOTH AMAZON AND EBAY */

type ListingDoc struct {
	ID                 string             `bson:"_id" json:"_id"`
	URLID              string             `bson:"urlid,omitempty" json:"urlid"`
	OurRef             string             `bson:"our_ref,omitempty" json:"our_ref"`
	DateEntered        string             `bson:"date_entered,omitempty" json:"date_entered"`
	ListingStartDate   string             `bson:"listing_start_date,omitempty" json:"listing_start_date"`
	ListingEndDate     string             `bson:"listing_end_date,omitempty" json:"listing_end_date"`
	SourceURL          string             `bson:"source_url,omitempty" json:"source_url"`
	Timestamp          string             `bson:"timestamp,omitempty" json:"timestamp"`
	TimeLeft           string             `bson:"time_left,omitempty" json:"time_left"`
	GeoLocation        string             `bson:"geo_location,omitempty" json:"geo_location"`
	ItemCard           ItemCard           `bson:"item_card,omitempty" json:"item_card"`
	AffInfo            AffInfo            `bson:"aff_info,omitempty" json:"aff_info"`
	AmazonInfo         AmazonInfo         `bson:"amazon_info,omitempty" json:"amazon_info"`
	AttributesInfo     AttributesInfo     `bson:"attributes_info,omitempty" json:"attributes_info"`
	CallInfo           CallInfo           `bson:"call_info,omitempty" json:"call_info"`
	CategoryInfo       CategoryInfo       `bson:"category_info,omitempty" json:"category_info"`
	ConditionInfo      ConditionInfo      `bson:"condition_info,omitempty" json:"condition_info"`
	DomainInfo         DomainInfo         `bson:"domain_info,omitempty" json:"domain_info"`
	EbayInfo           EbayInfo           `bson:"ebay_info,omitempty" json:"ebay_info"`
	FeaturesInfo       FeaturesInfo       `bson:"features_info,omitempty" json:"features_info"`
	ImageInfo          ImageInfo          `bson:"image_info,omitempty" json:"image_info"`
	ItemSizeInfo       ItemSizeInfo       `bson:"item_size_info,omitempty" json:"item_size_info"`
	KeywordsInfo       KeywordsInfo       `bson:"keywords_info,omitempty" json:"keywords_info"`
	LocationInfo       LocationInfo       `bson:"location_info,omitempty" json:"location_info"`
	PriceInfo          PriceInfo          `bson:"price_info,omitempty" json:"price_info"`
	RatingsInfo        RatingsInfo        `bson:"ratings_info,omitempty" json:"ratings_info"`
	ReviewsInfo        ReviewsInfo        `bson:"reviews_info,omitempty" json:"reviews_info"`
	SellerInfo         SellerInfo         `bson:"seller_info,omitempty" json:"seller_info"`
	ShippingInfo       ShippingInfo       `bson:"shipping_info,omitempty" json:"shipping_info"`
	SpecificationsInfo SpecificationsInfo `bson:"specifications_info,omitempty" json:"specifications_info"`
	UpdatedDocInfo     UpdatedDocInfo     `bson:"updated_doc_info,omitempty" json:"updated_doc_info"`
	//VideoInfo          VideoInfo          `bson:"video_info,omitempty" json:"video_info"`
	CompatibilityInfo CompatibilityInfo `bson:"compatibility_info,omitempty" json:"compatibility_info"`
	PositionInfo      PositionInfo      `bson:"position_info,omitempty" json:"position_info"`
}
type ItemCard struct {
	ItemName            string   `bson:"item_name,omitempty" json:"item_name,omitempty"`
	ItemAccurateName    string   `bson:"item_long_name,omitempty" json:"item_long_name"`
	ItemCondition       string   `bson:"item_condition,omitempty" json:"item_condition,omitempty"`
	ItemDescription     string   `bson:"item_description,omitempty" json:"item_description"`
	ItemImage           string   `bson:"item_image,omitempty" json:"item_image"`
	ItemTitle           string   `bson:"item_title,omitempty" json:"item_title"`
	ItemTitleUrl        string   `bson:"item_title_url,omitempty" json:"item_title_url"`
	ItemSubTitle        string   `bson:"item_subtitle,omitempty" json:"item_subtitle,omitempty"`
	ItemPrice           string   `bson:"item_price,omitempty" json:"item_price"`
	ItemShipping        string   `bson:"item_shipping,omitempty" json:"item_shipping"`
	ItemShipsTo         []string `bson:"item_ships_to,omitempty" json:"item_ships_to"`
	ItemCountry         string   `bson:"item_country,omitempty" json:"item_country"`
	ItemCurrency        string   `bson:"item_currency,omitempty" json:"item_currency"`
	ItemCurrencySymbol  string   `bson:"item_currency_symbol,omitempty" json:"item_currency_symbol"`
	ItemStockStatus     string   `bson:"item_stock_status,omitempty" json:"item_stock_status"`
	ItemSellingStatus   string   `bson:"item_selling_status,omitempty" json:"item_selling_status"`
	ItemWebsite         string   `bson:"item_website,omitempty" json:"item_website"`
	ItemWebsiteCode     string   `bson:"item_website_code,omitempty" json:"item_website_code"` // 3849 : ebay 4942 : amazon
	ItemCategoryName    string   `bson:"item_category_name,omitempty" json:"item_category_name"`
	ItemCategoryId      string   `bson:"item_category_id,omitempty" json:"item_category_id"`
	ItemEndTime         string   `bson:"item_end_time,omitempty" json:"item_end_time"`
	ItemMaterials       []string `bson:"item_materials,omitempty" json:"item_materials,omitempty"`
	ItemMaterial        string   `bson:"item_material,omitempty" json:"item_material,omitempty"`
	ItemColors          []string `bson:"item_colors,omitempty" json:"item_colors,omitempty"`
	ItemColor           string   `bson:"item_color,omitempty" json:"item_color,omitempty"`
	ItemSize            string   `bson:"item_size,omitempty" json:"item_size,omitempty"`
	ItemSiteLink        string   `bson:"item_site_link,omitempty" json:"item_site_link,omitempty"`
	ItemManufCode       string   `bson:"item_manuf_code,omitempty" json:"item_manuf_code"`
	ItemFamily          string   `bson:"item_family,omitempty" json:"item_family"`
	ItemBrand           string   `bson:"item_brand,omitempty" json:"item_brand"`
	ItemManufacturer    string   `bson:"item_manufacturer,omitempty" json:"item_manufacturer"`
	ItemReturnsAccepted string   `bson:"item_returns_accepted,omitempty" json:"item_returns_accepted"`
	ItemSourceLink      string   `bson:"item_source_link,omitempty" json:"item_source_link"`
	ItemWeight          string   `bson:"item_weight,omitempty" json:"item_weight"`
	ItemModelNumber     string   `bson:"item_model_number,omitempty" json:"item_model_number"`
	ItemDimensions      string   `bson:"item_dimensions,omitempty" json:"item_dimensions"`
	ItemPaymentMethod   string   `bson:"item_payment_method,omitempty" json:"item_payment_method"`
}
type AffInfo struct {
	AffLink  string `bson:"aff_link,omitempty" json:"aff_link"`
	AffPixel string `bson:"aff_pixel,omitempty" json:"aff_pixel"`
}
type AmazonInfo struct {
	AmazonRatingsTotal int         `bson:"amazon_ratings_total,omitempty" json:"amazon_ratings_total,omitempty"`
	AmazonIsPrime      bool        `bson:"amazon_is_prime,omitempty" json:"amazon_is_prime,omitempty"`
	ASIN               string      `bson:"asin,omitempty" json:"asin"`
	AmazonRating       float64     `bson:"amazon_rating,omitempty" json:"amazon_rating,omitempty"`
	AmazonIsBundle     bool        `bson:"amazon_is_bundle,omitempty" json:"amazon_is_bundle,omitempty"`
	SellOnAmazon       bool        `bson:"sell_on_amazon,omitempty" json:"sell_on_amazon,omitempty"`
	FirstAvailableDate string      `bson:"first_available_date,omitempty" json:"first_available_date,omitempty"`
	Sponsored          bool        `bson:"sponsored,omitempty" json:"sponsored,omitempty"`
	SearchAlias        SearchAlias `bson:"search_alias,omitempty" json:"search_alias,omitempty"`
}
type AttributesInfo struct {
	IsGenuine      bool `bson:"is_genuine,omitempty" json:"is_genuine,omitempty"`
	IsOEM          bool `bson:"is_oem,omitempty" json:"is_oem,omitempty"`
	HasWarranty    bool `bson:"has_warranty,omitempty" json:"has_warranty,omitempty"`
	IsUniversal    bool `bson:"is_universal,omitempty" json:"is_universal,omitempty"`
	IsGeneric      bool `bson:"is_generic,omitempty" json:"is_generic,omitempty"`
	IsPremium      bool `bson:"is_premium,omitempty" json:"is_premium,omitempty"`
	IsSports       bool `bson:"is_sports,omitempty" json:"is_sports,omitempty"`
	IsOriginal     bool `bson:"is_original,omitempty" json:"is_original,omitempty"`
	IsComplete     bool `bson:"is_complete,omitempty" json:"is_complete,omitempty"`
	IsPerformance  bool `bson:"is_performance,omitempty" json:"is_performance,omitempty"`
	IsHeavyDuty    bool `bson:"is_heavy_duty,omitempty" json:"is_heavy_duty,omitempty"`
	IsQuickFit     bool `bson:"is_quick_fit,omitempty" json:"is_quick_fit,omitempty"`
	IsQuickRelease bool `bson:"is_quick_release,omitempty" json:"is_quick_release,omitempty"`
	IsVintage      bool `bson:"is_vintage,omitempty" json:"is_vintage,omitempty"`
	IsRetrofit     bool `bson:"is_retrofit,omitempty" json:"is_retrofit,omitempty"`
}
type CallInfo struct {
	ItemName               string `bson:"item_name,omitempty" json:"item_name"`
	Website                string `bson:"website,omitempty" json:"website"`
	Geo                    string `bson:"geo,omitempty" json:"geo"`
	AmazonCallCategoryId   string `bson:"amazon_call_category_id,omitempty" json:"amazon_call_category_id"`
	EbayCallCategoryId     string `bson:"ebay_call_category_id,omitempty" json:"ebay_call_category_id"`
	AmazonCallCategoryName string `bson:"amazon_call_category_name,omitempty" json:"amazon_call_category_name"`
	EbayCallCategoryName   string `bson:"ebay_call_category_name,omitempty" json:"ebay_call_category_name"`
	ASIN                   string `bson:"asin,omitempty" json:"asin"`
}
type CategoryInfo struct {
	AmazonCategories []struct {
		Name       string `bson:"name,omitempty" json:"name"`
		Link       string `bson:"link,omitempty" json:"link"`
		CategoryID string `bson:"category_id,omitempty" json:"category_id"`
	} `bson:"amazon_categories,omitempty" json:"amazon_categories,omitempty"`

	SiteParentName string `bson:"site_parent_name" json:"site_parent_name"`
	SiteParentID   string `bson:"site_parent_id" json:"site_parent_id"`

	SiteCategoryName string `bson:"site_category_name" json:"site_category_name"`
	SiteCategoryID   string `bson:"site_category_id" json:"site_category_id"`

	SiteGlobalParentName string `bson:"site_global_parent_name" json:"site_global_parent_name"`
	SiteGlobalParentID   string `bson:"site_global_parent_id" json:"site_global_parent_id"`

	SiteGlobalCategoryName string `bson:"site_global_category_name" json:"site_global_category_name"`
	SiteGlobalCategoryID   string `bson:"site_global_category_id" json:"site_global_category_id"`

	SiteSubName string `bson:"site_sub_name,omitempty" json:"site_sub_name"`
	SiteSubID   string `bson:"site_sub_id,omitempty" json:"site_sub_id"`

	SecondaryCategoryName string `bson:"secondary_category_name,omitempty" json:"secondary_category_name"`
	SecondaryCategoryID   string `bson:"secondary_category_id,omitempty" json:"secondary_category_id"`
}
type CompatibilityInfo struct {
	Make          string   `bson:"make,omitempty" json:"make,omitempty"`
	Model         string   `bson:"model,omitempty" json:"model,omitempty"`
	YearStart     string   `bson:"year_start,omitempty" json:"year_start,omitempty"`
	YearEnd       string   `bson:"year_end,omitempty" json:"year_end,omitempty"`
	ValveSize     string   `bson:"valve_size,omitempty" json:"valve_size,omitempty"`
	EngineSize    string   `bson:"engine_size,omitempty" json:"engine_size,omitempty"`
	GearboxSpeeds string   `bson:"gearbox_speeds,omitempty" json:"gearbox_speeds,omitempty"`
	Transmission  string   `bson:"transmission,omitempty" json:"transmission,omitempty"`
	FuelType      string   `bson:"fuel_type,omitempty" json:"fuel_type,omitempty"`
	BodyType      string   `bson:"body_type,omitempty" json:"body_type,omitempty"`
	MkNumber      []string `bson:"mk_number,omitempty" json:"mk_number,omitempty"`
	Years         []string `bson:"years,omitempty" json:"years,omitempty"`
	Compatibility string   `bson:"compatibility,omitempty" json:"compatibility,omitempty"`
}
type PositionInfo struct {
	ItemPosition string `bson:"item_position,omitempty" json:"item_position"`
	ItemSide     string `bson:"item_side,omitempty" json:"item_side"`
	IsUpper      bool   `bson:"is_upper,omitempty" json:"is_upper"`
	IsLower      bool   `bson:"is_lower,omitempty" json:"is_lower"`
	IsFront      bool   `bson:"is_front,omitempty" json:"is_front"`
	IsRear       bool   `bson:"is_rear,omitempty" json:"is_rear"`
	IsRightSide  bool   `bson:"is_right,omitempty" json:"is_right"`
	IsLeftSide   bool   `bson:"is_left,omitempty" json:"is_left"`
	IsInterior   bool   `bson:"is_interior,omitempty" json:"is_interior"`
	IsExterior   bool   `bson:"is_exterior,omitempty" json:"is_exterior"`
}
type ConditionInfo struct {
	ConditionId string `bson:"condition_id,omitempty" json:"condition_id"`
	Condition   string `bson:"condition,omitempty" json:"condition"`
	IsExcellent bool   `bson:"is_excellent,omitempty" json:"is_excellent"`
	IsMint      bool   `bson:"is_mint,omitempty" json:"is_mint"`
	IsGood      bool   `bson:"is_good,omitempty" json:"is_good"`
}
type DomainInfo struct {
	Name   string `bson:"name,omitempty" json:"name"`
	Domain string `bson:"domain,omitempty" json:"domain"`
	URL    string `bson:"url,omitempty" json:"url"`
	Geo    string `bson:"geo,omitempty" json:"geo"`
}
type EbayInfo struct {
	IsTopRatedListing       string `bson:"is_top_rated_listing,omitempty" json:"is_top_rated_listing"`
	EbayID                  string `bson:"ebay_id,omitempty" json:"ebay_id"`
	ProductId               string `bson:"product_id,omitempty" json:"product_id"`
	BuyItNowAvailable       string `bson:"buy_it_now_available,omitempty" json:"buy_it_now_available"`
	BuyItNowPrice           string `bson:"buy_it_now_price,omitempty" json:"buy_it_now_price"`
	EbayGlobalID            string `bson:"global_item_id,omitempty" json:"global_item_id"`
	ListingType             string `bson:"listing_type,omitempty" json:"listing_type"`
	IsMultiVariationListing string `bson:"is_multi_variation_listing,omitempty" json:"is_multi_variation_listing"`
	WatchCount              string `bson:"watch_count,omitempty" json:"watch_count"`
	BestOfferEnabled        string `bson:"best_offer_enabled,omitempty" json:"best_offer_enabled"`
	Gift                    string `bson:"gift,omitempty" json:"gift"`
	DiscountPriceInfo       string `bson:"discount_price_info,omitempty" json:"discount_price_info"`
	ReturnsAccepted         string `bson:"returns_accepted,omitempty" json:"returns_accepted"`
	TimeLeft                string `bson:"time_left,omitempty" json:"time_left"`
	CurrencyId              string `bson:"currency_id,omitempty" json:"currency_id"`
	StoreName               string `bson:"store_name,omitempty" json:"store_name"`
	StoreURL                string `bson:"store_url,omitempty" json:"store_url"`
	BidCount                string `bson:"bid_count,omitempty" json:"bid_count"`
	EbayPlusEnabled         bool   `bson:"ebay_plus_enabled,omitempty" json:"ebay_plus_enabled"`
}
type FeaturesInfo struct {
	FeatureBulletsCount int32  `bson:"feature_bullets_count,omitempty" json:"feature_bullets_count"`
	FeatureBulletsFlat  string `bson:"feature_bullets_flat,omitempty" json:"feature_bullets_flat"`
}
type ImageInfo struct {
	DisplayImage string `bson:"display_image,omitempty" json:"display_image"`
	LargeImgUrl  string `bson:"large_img_url,omitempty" json:"large_img_url"`
	MediumImgUrl string `bson:"medium_img_url,omitempty" json:"medium_img_url"`
	SmallImgUrl  string `bson:"small_img_url,omitempty" json:"small_img_url"`
	ImagesCount  int    `bson:"images_count,omitempty" json:"images_count"`
	Images       []struct {
		Link    string `bson:"link,omitempty" json:"link"`
		Variant string `bson:"variant,omitempty" json:"variant"`
	}
}
type ItemSizeInfo struct {
	ItemSize string `bson:"item_size,omitempty" json:"item_size"`
	IsLarge  bool   `bson:"is_large,omitempty" json:"is_large"`
	IsMedium bool   `bson:"is_medium,omitempty" json:"is_medium"`
	IsSmall  bool   `bson:"is_small,omitempty" json:"is_small"`
	IsLong   bool   `bson:"is_long,omitempty" json:"is_long"`
	IsShort  bool   `bson:"is_short,omitempty" json:"is_short"`
}
type KeywordsInfo struct {
	Keywords string `bson:"keywords,omitempty" json:"keywords"`
}
type LocationInfo struct {
	Location   string `bson:"location,omitempty" json:"location"`
	PostalCode string `bson:"postal_code,omitempty" json:"postal_code"`
	Country    string `bson:"country,omitempty" json:"country"`
}
type PriceInfo struct {
	DisplayPrice      string   `bson:"display_price,omitempty" json:"display_price"`
	CurrentPrice      string   `bson:"current_price,omitempty" json:"current_price"`
	Currency          string   `bson:"currency,omitempty" json:"currency"`
	Symbol            string   `bson:"symbol,omitempty" json:"symbol"`
	CurrencyID        string   `bson:"currency_id,omitempty" json:"currency_id"`
	IsOffersAvailable string   `bson:"is_offers_available,omitempty" json:"is_offers_available"`
	IsBuyItNow        string   `bson:"is_buy_it_now,omitempty" json:"is_buy_it_now"`
	BuyItNowPrice     float64  `bson:"buy_it_now_price,omitempty" json:"buy_it_now_price"`
	PaymentMethods    []string `bson:"payment_methods,omitempty" json:"payment_methods"`
	PriceRaw          string   `bson:"price_raw,omitempty" json:"price_raw"`
	DiscountPriceInfo string   `bson:"discount_price_info,omitempty" json:"discount_price_info"`
}
type RatingsInfo struct {
	Rating          float64         `bson:"rating,omitempty" json:"rating"`
	RatingsTotal    int             `bson:"ratings_total,omitempty" json:"ratings_total"`
	RatingBreakdown RatingBreakdown `bson:"rating_breakdown,omitempty" json:"rating_breakdown"`
}
type ReviewsInfo struct {
	ReviewsTotal int `bson:"reviews_total,omitempty" json:"reviews_total"`
}
type SearchAlias struct {
	Title string `bson:"title,omitempty" json:"title"`
	Value string `bson:"value,omitempty" json:"value"`
}
type SellerInfo struct {
	UserName                string `bson:"user_name,omitempty" json:"user_name"`
	SellerUserName          string `bson:"seller_username,omitempty" json:"seller_username"`
	FeedbackRatingStar      string `bson:"feedback_rating_star,omitempty" json:"feedback_rating_star"`
	FeedbackScore           string `bson:"feedback_score,omitempty" json:"feedback_score"`
	PositiveFeedbackPercent string `bson:"positive_feedback_percent,omitempty" json:"positive_feedback_percent"`
	TopRatedSeller          string `bson:"top_rated_seller,omitempty" json:"top_rated_seller"`
	SellerLogo              string `bson:"seller_logo" json:"seller_logo"`
}
type ShippingInfo struct {
	ShipsTo                 []string `bson:"ships_to,omitempty" json:"ships_to"`
	ShipsToWorldwide        bool     `bson:"ships_to_worldwide,omitempty" json:"ships_to_worldwide"`
	DisplayShippingPrice    string   `bson:"display_shipping_price,omitempty" json:"display_shipping_price"`
	ShippingPrice           float32  `bson:"shipping_price,omitempty" json:"shipping_price"`
	IsFree                  bool     `bson:"is_free,omitempty" json:"is_free"`
	Type                    string   `bson:"type,omitempty" json:"type"`
	IsNextDay               string   `bson:"is_next_day,omitempty" json:"is_next_day"`
	Tagline                 string   `bson:"tagline,omitempty" json:"tagline"`
	ShippingType            string   `bson:"shipping_type,omitempty" json:"shipping_type"`
	ExpeditedShipping       string   `bson:"expedited_shipping,omitempty" json:"expedited_shipping"`
	OneDayShippingAvailable bool     `bson:"one_day_shipping_available,omitempty" json:"one_day_shipping_available"`
	HandlingTime            string   `bson:"handling_time,omitempty" json:"handling_time"`
}
type SpecificationsInfo struct {
	SpecificationsFlat string `bson:"specifications_flat,omitempty" json:"specifications_flat"`
}
type RatingBreakdown struct {
	FiveStar  FiveStar  `bson:"five_star,omitempty" json:"five_star"`
	FourStar  FourStar  `bson:"four_star,omitempty" json:"four_star"`
	ThreeStar ThreeStar `bson:"three_star,omitempty" json:"three_star"`
	TwoStar   TwoStar   `bson:"two_star,omitempty" json:"two_star"`
	OneStar   OneStar   `bson:"one_star,omitempty" json:"one_star"`
}
type FiveStar struct {
	Percentage int `bson:"percentage,omitempty" json:"percentage"`
	Count      int `bson:"count,omitempty" json:"count"`
}
type FourStar struct {
	Percentage int `bson:"percentage,omitempty" json:"percentage"`
	Count      int `bson:"count,omitempty" json:"count"`
}
type ThreeStar struct {
	Percentage int `bson:"percentage,omitempty" json:"percentage"`
	Count      int `bson:"count,omitempty" json:"count"`
}
type TwoStar struct {
	Percentage int `bson:"percentage,omitempty" json:"percentage"`
	Count      int `bson:"count,omitempty" json:"count"`
}
type OneStar struct {
	Percentage int `bson:"percentage,omitempty" json:"percentage"`
	Count      int `bson:"count,omitempty" json:"count"`
}
type UpdatedDocInfo struct {
	UpdatedYears        bool `bson:"update_years,omitempty" json:"update_years"`
	UpdatedCats         bool `bson:"update_cats,omitempty" json:"update_cats"`
	UpdatedItemNameLong bool `bson:"updated_itemname_long,omitempty" json:"updated_itemname_long"`
	UpdatedMake         bool `bson:"updated_make,omitempty" json:"updated_make"`
	UpdatedModel        bool `bson:"updated_model,omitempty" json:"updated_model"`
}

//type VideoInfo struct {
//	Videos []struct {
//		DurationSeconds int    `bson:"duration_seconds,omitempty" json:"duration_seconds"`
//		Width           int32  `bson:"width,omitempty" json:"width"`
//		Height          int32  `bson:"height,omitempty" json:"height"`
//		Link            string `bson:"link,omitempty" json:"link"`
//		Thumbnail       string `bson:"thumbnail,omitempty" json:"thumbnail"`
//		IsHeroVideo     bool   `bson:"is_hero_video,omitempty" json:"is_hero_video"`
//		Variant         string `bson:"variant,omitempty" json:"variant"`
//		GroupID         string `bson:"group_id,omitempty" json:"group_id"`
//		GroupType       string `bson:"group_type,omitempty" json:"group_type"`
//		Title           string `bson:"title,omitempty" json:"title"`
//	} `bson:"videos,omitempty" json:"omitempty"`
//	VideoCount int `bson:"video_count,omitempty" json:"video_count"`
//}

type ListingDocs struct {
	ListingDocs []ListingDoc `bson:"listing_docs,omitempty" json:"listing_docs"`
}
