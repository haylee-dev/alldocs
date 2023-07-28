package func_ebay

import (
	"os"
	"strings"
)

func GetSortOrder(sort_order string) string {
	SortOrder := ""

	/*
		'BestMatch' => 'Best Match',
		'BidCountFewest' => 'Bid Count Fewest',
		'BidCountMost' => 'Bid Count Most',
		'CountryAscending' => 'Country Ascending',
		'CountryDescending' => 'Country Descending',
		'CurrentPriceHighest' => 'Current Highest Price',
		'DistanceNearest' => 'Nearest Distance',
		'EndTimeSoonest' => 'End Time Soonest',
		'PricePlusShippingHighest' => 'Price Plus Shipping Highest',
		'PricePlusShippingLowest' => 'Price Plus Shipping Lowest',
		'StartTimeNewest' => 'Start Time Newest'
	*/

	if sort_order == "best-match" {
		SortOrder = "BestMatch"
	}
	if sort_order == "bid-count-fewest" {
		SortOrder = "BidCountFewest"
	}
	if sort_order == "bid-count-most" {
		SortOrder = "BidCountMost"
	}
	if sort_order == "country-ascending" {
		SortOrder = "CountryAscending"
	}
	if sort_order == "country-descending" {
		SortOrder = "CountryDescending"
	}
	if sort_order == "current-price-highest" {
		SortOrder = "CurrentPriceHighest"
	}
	if sort_order == "end-time-soonest" {
		SortOrder = "EndTimeSoonest"
	}
	if sort_order == "price-plus-shipping-highest" {
		SortOrder = "PricePlusShippingHighest"
	}
	if sort_order == "price-plus-shipping-lowest" {
		SortOrder = "PricePlusShippingLowest"
	}
	if sort_order == "start-time-newest" {
		SortOrder = "StartTimeNewest"
	}
	return SortOrder
}
func SetCurrencySymbol(CurrencyId string, GlobalId string) string {
	currency_symbol := "?"
	switch {
	case CurrencyId == "AUD":
		currency_symbol = "$"
	case CurrencyId == "CAD":
		currency_symbol = "£"
	case CurrencyId == "CHF":
		currency_symbol = "₣"
	case CurrencyId == "CNY":
		currency_symbol = "¥"
	case CurrencyId == "EUR":
		currency_symbol = "€"
	case CurrencyId == "GBP":
		currency_symbol = "£"
	case CurrencyId == "HKD":
		currency_symbol = "$"
	case CurrencyId == "INR":
		currency_symbol = "₹"
	case CurrencyId == "MYR":
		currency_symbol = "RM"
	case CurrencyId == "PHP":
		currency_symbol = "₱"
	case CurrencyId == "PLN":
		currency_symbol = "zł"
	case CurrencyId == "SEK":
		currency_symbol = "kr"
	case CurrencyId == "SGD":
		currency_symbol = "$"
	case CurrencyId == "TWD":
		currency_symbol = "$"
	case CurrencyId == "USD":
		currency_symbol = "$"
	case GlobalId == "EBAY-GB":
		currency_symbol = "£"
	}
	return currency_symbol
}
func CheckCatIsSub(category_name string) (bool, string, string, string, string, string, string) {
	category_name_cl := CleanCategory(category_name)

	IsSubCategory := false
	CategoryId := ""
	CategoryName := category_name_cl
	ParentId := ""
	ParentName := ""
	SubId := ""
	SubName := ""

	if category_name_cl == "warning-lights" {
		IsSubCategory = true
		ParentId = "33707"
		ParentName = "lighting-and-bulbs"
		CategoryId = "262212"
		CategoryName = "accessory-lighting"
		SubId = "121958"
		SubName = "warning-lights"
	}
	if category_name_cl == "light-bars" {
		IsSubCategory = true
		ParentId = "33707"
		ParentName = "lighting-and-bulbs"
		CategoryId = "262212"
		CategoryName = "accessory-lighting"
		SubId = "184669"
		SubName = "light-bars"
	}
	if category_name_cl == "accessory-and-off-road-lighting" {
		IsSubCategory = true
		ParentId = "33707"
		ParentName = "lighting-and-bulbs"
		CategoryId = "262212"
		CategoryName = "accessory-lighting"
		SubId = "262213"
		SubName = "accessory-and-off-road-lighting"
	}
	if category_name_cl == "wheel-and-underbody-lights" {
		IsSubCategory = true
		ParentId = "33707"
		ParentName = "lighting-and-bulbs"
		CategoryId = "262212"
		CategoryName = "accessory-lighting"
		SubId = "262214"
		SubName = "wheel-and-underbody-lights"
	}
	if category_name_cl == "batteries" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262094"
		CategoryName = "batteries-bus-and-fuel-cells"
		SubId = "177703"
		SubName = "batteries"
	}
	if category_name_cl == "battery-management-systems" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262094"
		CategoryName = "batteries-bus-and-fuel-cells"
		SubId = "262095"
		SubName = "battery-management-systems"
	}
	if category_name_cl == "battery-and-bus-units" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262094"
		CategoryName = "batteries-bus-and-fuel-cells"
		SubId = "262096"
		SubName = "battery-and-bus-units"
	}
	if category_name_cl == "fuel-cells-and-parts" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262094"
		CategoryName = "batteries-bus-and-fuel-cells"
		SubId = "262097"
		SubName = "fuel-cells-and-parts"
	}
	if category_name_cl == "hinges-latches-and-additional-bonnet-components" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262155"
		CategoryName = "bonnets"
		SubId = "262156"
		SubName = "hinges-latches-and-additional-bonnet-components"
	}
	if category_name_cl == "lift-support-and-prop-rods" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262155"
		CategoryName = "bonnets"
		SubId = "262158"
		SubName = "lift-support-and-prop-rods"
	}
	if category_name_cl == "bumper-inserts-and-covers" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262154"
		CategoryName = "bumpers-and-components"
		SubId = "262146"
		SubName = "bumper-inserts-and-covers"
	}
	if category_name_cl == "bumpers-and-reinforcements" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262154"
		CategoryName = "bumpers-and-components"
		SubId = "33640"
		SubName = "bumpers-and-reinforcements"
	}
	if category_name_cl == "air-dams" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262154"
		CategoryName = "bumpers-and-components"
		SubId = "38658"
		SubName = "air-dams"
	}
	if category_name_cl == "centre-and-overhead-consoles" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262187"
		CategoryName = "centre-overhead-consoles-and-parts"
		SubId = "262188"
		SubName = "centre-and-overhead-consoles"
	}
	if category_name_cl == "centre-and-overhead-console-parts" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262187"
		CategoryName = "centre-overhead-consoles-and-parts"
		SubId = "262189"
		SubName = "centre-and-overhead-console-parts"
	}
	if category_name_cl == "other-charging-components" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "177702"
		SubName = "other-charging-components"
	}
	if category_name_cl == "charging-cables" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "262102"
		SubName = "charging-cables"
	}
	if category_name_cl == "charger-carry-cases" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "262103"
		SubName = "charger-carry-cases"
	}
	if category_name_cl == "charging-converters" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "262104"
		SubName = "charging-converters"
	}
	if category_name_cl == "charger-holders" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "262105"
		SubName = "charger-holders"
	}
	if category_name_cl == "charging-upgrade-kits" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262101"
		CategoryName = "charging-components"
		SubId = "262106"
		SubName = "charging-upgrade-kits"
	}
	if category_name_cl == "charging-stations" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262098"
		CategoryName = "charging-units"
		SubId = "262099"
		SubName = "charging-stations"
	}
	if category_name_cl == "mobile-charging" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262098"
		CategoryName = "charging-units"
		SubId = "262100"
		SubName = "mobile-charging"
	}
	if category_name_cl == "heatsinks" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262107"
		CategoryName = "cooling-components"
		SubId = "262108"
		SubName = "heatsinks"
	}
	if category_name_cl == "fans" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262107"
		CategoryName = "cooling-components"
		SubId = "262109"
		SubName = "fans"
	}
	if category_name_cl == "cooling-plates" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262107"
		CategoryName = "cooling-components"
		SubId = "262110"
		SubName = "cooling-plates"
	}
	if category_name_cl == "dashboard-panels" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262190"
		CategoryName = "dashboard-panels-and-glove-boxes"
		SubId = "262191"
		SubName = "dashboard-panels"
	}
	if category_name_cl == "glove-boxes-doors-and-latches" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262190"
		CategoryName = "dashboard-panels-and-glove-boxes"
		SubId = "33698"
		SubName = "glove-boxes-doors-and-latches"
	}
	if category_name_cl == "doors-and-door-skins" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "179850"
		SubName = "doors-and-door-skins"
	}
	if category_name_cl == "door-handles" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "179851"
		SubName = "door-handles"
	}
	if category_name_cl == "door-seals" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "262148"
		SubName = "door-seals"
	}
	if category_name_cl == "cargo-and-sliding-doors" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "262149"
		SubName = "cargo-and-sliding-doors"
	}
	if category_name_cl == "lift-support-latches-hinges-and-additional-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "262150"
		SubName = "lift-support-latches-hinges-and-additional-parts"
	}
	if category_name_cl == "tail-lifts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "33647"
		SubName = "tail-lifts"
	}
	if category_name_cl == "boot-and-hatch-lids" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262147"
		CategoryName = "doors-boot-lids-and-hatches"
		SubId = "33656"
		SubName = "boot-and-hatch-lids"
	}
	if category_name_cl == "drivetrain-motors" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262111"
		CategoryName = "drivetrain-motors-and-parts"
		SubId = "177708"
		SubName = "drivetrain-motors"
	}
	if category_name_cl == "drivetrain-motor-parts" {
		IsSubCategory = true
		ParentId = "177701"
		ParentName = "electric-hybrid-and-phev-specific-parts"
		CategoryId = "262111"
		CategoryName = "drivetrain-motors-and-parts"
		SubId = "262112"
		SubName = "drivetrain-motor-parts"
	}
	if category_name_cl == "high-pressure-fuel-pumps" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "262070"
		CategoryName = "fuel-injection-parts"
		SubId = "262071"
		SubName = "high-pressure-fuel-pumps"
	}
	if category_name_cl == "additional-fuel-injection-parts" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "262070"
		CategoryName = "fuel-injection-parts"
		SubId = "33553"
		SubName = "additional-fuel-injection-parts"
	}
	if category_name_cl == "fuel-injectors" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "262070"
		CategoryName = "fuel-injection-parts"
		SubId = "33554"
		SubName = "fuel-injectors"
	}
	if category_name_cl == "brackets-and-hardware" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262126"
		CategoryName = "gaskets-seals-and-hardware"
		SubId = "262127"
		SubName = "brackets-and-hardware"
	}
	if category_name_cl == "gasket-kits" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262126"
		CategoryName = "gaskets-seals-and-hardware"
		SubId = "262128"
		SubName = "gasket-kits"
	}
	if category_name_cl == "hardware-kits" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262126"
		CategoryName = "gaskets-seals-and-hardware"
		SubId = "262129"
		SubName = "hardware-kits"
	}
	if category_name_cl == "gasket-and-seals" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262126"
		CategoryName = "gaskets-seals-and-hardware"
		SubId = "33665"
		SubName = "gasket-and-seals"
	}
	if category_name_cl == "individual-gauges" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "33672"
		CategoryName = "gauges"
		SubId = "262192"
		SubName = "individual-gauges"
	}
	if category_name_cl == "individual-clusters" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "33672"
		CategoryName = "gauges"
		SubId = "33675"
		SubName = "individual-clusters"
	}
	if category_name_cl == "instrument-and-gauge-parts" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "33672"
		CategoryName = "gauges"
		SubId = "43951"
		SubName = "instrument-and-gauge-parts"
	}
	if category_name_cl == "car-glass" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "33683"
		CategoryName = "glass-and-window-parts"
		SubId = "33684"
		SubName = "car-glass"
	}
	if category_name_cl == "window-seals-gaskets-and-trims" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "33683"
		CategoryName = "glass-and-window-parts"
		SubId = "33685"
		SubName = "window-seals-gaskets-and-trims"
	}
	if category_name_cl == "window-tinting" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "33683"
		CategoryName = "glass-and-window-parts"
		SubId = "63689"
		SubName = "window-tinting"
	}
	if category_name_cl == "additional-glass-and-window-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "33683"
		CategoryName = "glass-and-window-parts"
		SubId = "6781"
		SubName = "additional-glass-and-window-parts"
	}
	if category_name_cl == "sun-visors-wind-and-bug-deflectors" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262153"
		CategoryName = "guards-and-protection"
		SubId = "38659"
		SubName = "sun-visors-wind-and-bug-deflectors"
	}
	if category_name_cl == "splash-guards-and-mud-flaps" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262153"
		CategoryName = "guards-and-protection"
		SubId = "50455"
		SubName = "splash-guards-and-mud-flaps"
	}
	if category_name_cl == "additional-guards-and-protection" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262153"
		CategoryName = "guards-and-protection"
		SubId = "50457"
		SubName = "additional-guards-and-protection"
	}
	if category_name_cl == "airbags" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262195"
		CategoryName = "interior-safety"
		SubId = "177710"
		SubName = "airbags"
	}
	if category_name_cl == "airbag-parts" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262195"
		CategoryName = "interior-safety"
		SubId = "177711"
		SubName = "airbag-parts"
	}
	if category_name_cl == "other-interior-safety" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262195"
		CategoryName = "interior-safety"
		SubId = "262196"
		SubName = "other-interior-safety"
	}
	if category_name_cl == "seatbelt-and-parts" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262195"
		CategoryName = "interior-safety"
		SubId = "33725"
		SubName = "seatbelt-and-parts"
	}
	if category_name_cl == "body-side-and-quarter-panels" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262165"
		CategoryName = "panels"
		SubId = "262166"
		SubName = "body-side-and-quarter-panels"
	}
	if category_name_cl == "fuel-flaps-and-components" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262165"
		CategoryName = "panels"
		SubId = "262167"
		SubName = "fuel-flaps-and-components"
	}
	if category_name_cl == "additional-panels" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262165"
		CategoryName = "panels"
		SubId = "262168"
		SubName = "additional-panels"
	}
	if category_name_cl == "fenders" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262165"
		CategoryName = "panels"
		SubId = "33644"
		SubName = "fenders"
	}
	if category_name_cl == "parking-cameras" {
		IsSubCategory = true
		ParentId = "262064"
		ParentName = "advanced-driver-assistance-systems"
		CategoryId = "258037"
		CategoryName = "parking-assistance"
		SubId = "174593"
		SubName = "parking-cameras"
	}
	if category_name_cl == "camera-monitor-and-sensor-kits" {
		IsSubCategory = true
		ParentId = "262064"
		ParentName = "advanced-driver-assistance-systems"
		CategoryId = "258037"
		CategoryName = "parking-assistance"
		SubId = "184616"
		SubName = "camera-monitor-and-sensor-kits"
	}
	if category_name_cl == "other-parking-assistance" {
		IsSubCategory = true
		ParentId = "262064"
		ParentName = "advanced-driver-assistance-systems"
		CategoryId = "258037"
		CategoryName = "parking-assistance"
		SubId = "258038"
		SubName = "other-parking-assistance"
	}
	if category_name_cl == "parking-sensors" {
		IsSubCategory = true
		ParentId = "262064"
		ParentName = "advanced-driver-assistance-systems"
		CategoryId = "258037"
		CategoryName = "parking-assistance"
		SubId = "61523"
		SubName = "parking-sensors"
	}
	if category_name_cl == "hinges-latches-and-additional-components" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262175"
		CategoryName = "pick-up-truck-beds-and-parts"
		SubId = "262176"
		SubName = "hinges-latches-and-additional-components"
	}
	if category_name_cl == "pickup-truck-beds-and-repair-sections" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262175"
		CategoryName = "pick-up-truck-beds-and-parts"
		SubId = "262177"
		SubName = "pickup-truck-beds-and-repair-sections"
	}
	if category_name_cl == "tailgates" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262175"
		CategoryName = "pick-up-truck-beds-and-parts"
		SubId = "262178"
		SubName = "tailgates"
	}
	if category_name_cl == "pickup-truck-bed-accessories" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262175"
		CategoryName = "pick-up-truck-beds-and-parts"
		SubId = "33655"
		SubName = "pickup-truck-bed-accessories"
	}
	if category_name_cl == "convertible-tops-and-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262169"
		CategoryName = "roofs-tops-and-sunroofs"
		SubId = "262170"
		SubName = "convertible-tops-and-parts"
	}
	if category_name_cl == "hardtops-roof-panels-and-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262169"
		CategoryName = "roofs-tops-and-sunroofs"
		SubId = "262171"
		SubName = "hardtops-roof-panels-and-parts"
	}
	if category_name_cl == "sunroofs-and-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262169"
		CategoryName = "roofs-tops-and-sunroofs"
		SubId = "262172"
		SubName = "sunroofs-and-parts"
	}
	if category_name_cl == "additional-parts" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262169"
		CategoryName = "roofs-tops-and-sunroofs"
		SubId = "262173"
		SubName = "additional-parts"
	}
	if category_name_cl == "warning-triangles" {
		IsSubCategory = true
		ParentId = "262266"
		ParentName = "safety-and-security-accessories"
		CategoryId = "180136"
		CategoryName = "safety"
		SubId = "121986"
		SubName = "warning-triangles"
	}
	if category_name_cl == "emergency-tools-and-kits" {
		IsSubCategory = true
		ParentId = "262266"
		ParentName = "safety-and-security-accessories"
		CategoryId = "180136"
		CategoryName = "safety"
		SubId = "180138"
		SubName = "emergency-tools-and-kits"
	}
	if category_name_cl == "other-safety-accessories" {
		IsSubCategory = true
		ParentId = "262266"
		ParentName = "safety-and-security-accessories"
		CategoryId = "180136"
		CategoryName = "safety"
		SubId = "180140"
		SubName = "other-safety-accessories"
	}
	if category_name_cl == "safety-jackets-and-vests" {
		IsSubCategory = true
		ParentId = "262266"
		ParentName = "safety-and-security-accessories"
		CategoryId = "180136"
		CategoryName = "safety"
		SubId = "99404"
		SubName = "safety-jackets-and-vests"
	}
	if category_name_cl == "headrests" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262199"
		CategoryName = "seats-parts-and-accessories"
		SubId = "262200"
		SubName = "headrests"
	}
	if category_name_cl == "additional-seat-parts" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262199"
		CategoryName = "seats-parts-and-accessories"
		SubId = "262201"
		SubName = "additional-seat-parts"
	}
	if category_name_cl == "seats" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262199"
		CategoryName = "seats-parts-and-accessories"
		SubId = "33701"
		SubName = "seats"
	}
	if category_name_cl == "seat-covers" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262199"
		CategoryName = "seats-parts-and-accessories"
		SubId = "33702"
		SubName = "seat-covers"
	}
	if category_name_cl == "seat-belt-shoulder-pads" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262199"
		CategoryName = "seats-parts-and-accessories"
		SubId = "50458"
		SubName = "seat-belt-shoulder-pads"
	}
	if category_name_cl == "mirror-assemblies" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262160"
		CategoryName = "side-view-mirrors"
		SubId = "262161"
		SubName = "mirror-assemblies"
	}
	if category_name_cl == "mirror-components" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262160"
		CategoryName = "side-view-mirrors"
		SubId = "262162"
		SubName = "mirror-components"
	}
	if category_name_cl == "mirror-glass" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "262160"
		CategoryName = "side-view-mirrors"
		SubId = "262163"
		SubName = "mirror-glass"
	}
	if category_name_cl == "rubbish-bins" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262202"
		CategoryName = "storage-and-organisers"
		SubId = "180117"
		SubName = "rubbish-bins"
	}
	if category_name_cl == "additional-storage-and-organisers" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262202"
		CategoryName = "storage-and-organisers"
		SubId = "262203"
		SubName = "additional-storage-and-organisers"
	}
	if category_name_cl == "cargo-nets-guards-and-boot-organisers" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262202"
		CategoryName = "storage-and-organisers"
		SubId = "63690"
		SubName = "cargo-nets-guards-and-boot-organisers"
	}
	if category_name_cl == "cup-holders" {
		IsSubCategory = true
		ParentId = "33694"
		ParentName = "interior-parts-and-accessories"
		CategoryId = "262202"
		CategoryName = "storage-and-organisers"
		SubId = "63691"
		SubName = "cup-holders"
	}
	if category_name_cl == "guide-rails" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262131"
		SubName = "guide-rails"
	}
	if category_name_cl == "sprockets" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262132"
		SubName = "sprockets"
	}
	if category_name_cl == "tensioners-and-pulleys" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262133"
		SubName = "tensioners-and-pulleys"
	}
	if category_name_cl == "timing-belts" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262134"
		SubName = "timing-belts"
	}
	if category_name_cl == "timing-chains" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262135"
		SubName = "timing-chains"
	}
	if category_name_cl == "timing-covers" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262136"
		SubName = "timing-covers"
	}
	if category_name_cl == "timing-kits" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262130"
		CategoryName = "timing-components-and-kits"
		SubId = "262137"
		SubName = "timing-kits"
	}
	if category_name_cl == "hitch-pins-and-locks" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "184629"
		SubName = "hitch-pins-and-locks"
	}
	if category_name_cl == "hitch-extenders-and-adapters" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "262234"
		SubName = "hitch-extenders-and-adapters"
	}
	if category_name_cl == "safety-pins-and-cables" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "262235"
		SubName = "safety-pins-and-cables"
	}
	if category_name_cl == "weight-distributions-and-sway-control" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "262236"
		SubName = "weight-distributions-and-sway-control"
	}
	if category_name_cl == "other-towing-accessories" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "262237"
		SubName = "other-towing-accessories"
	}
	if category_name_cl == "tow-ropes-and-chains" {
		IsSubCategory = true
		ParentId = "180143"
		ParentName = "towing-parts-and-accessories"
		CategoryId = "262233"
		CategoryName = "towing-accessories"
		SubId = "61960"
		SubName = "tow-ropes-and-chains"
	}
	if category_name_cl == "intercoolers" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "174107"
		CategoryName = "turbos-superchargers-and-intercoolers"
		SubId = "262073"
		SubName = "intercoolers"
	}
	if category_name_cl == "superchargers-and-parts" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "174107"
		CategoryName = "turbos-superchargers-and-intercoolers"
		SubId = "33741"
		SubName = "superchargers-and-parts"
	}
	if category_name_cl == "turbos-and-parts" {
		IsSubCategory = true
		ParentId = "33549"
		ParentName = "air-and-fuel-delivery"
		CategoryId = "174107"
		CategoryName = "turbos-superchargers-and-intercoolers"
		SubId = "33742"
		SubName = "turbos-and-parts"
	}
	if category_name_cl == "guide-rails" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262131"
		SubName = "guide-rails"
	}
	if category_name_cl == "sprockets" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262132"
		SubName = "sprockets"
	}
	if category_name_cl == "tensioners-and-pulleys" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262133"
		SubName = "tensioners-and-pulleys"
	}
	if category_name_cl == "timing-belts" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262134"
		SubName = "timing-belts"
	}
	if category_name_cl == "timing-chains" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262135"
		SubName = "timing-chains"
	}
	if category_name_cl == "timing-covers" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262136"
		SubName = "timing-covers"
	}
	if category_name_cl == "timing-kits" {
		IsSubCategory = true
		ParentId = "33612"
		ParentName = "engines-and-engine-parts"
		CategoryId = "262138"
		CategoryName = "valvetrain-components"
		SubId = "262137"
		SubName = "timing-kits"
	}
	if category_name_cl == "snow-chains-anti-skid-and-snow-socks" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262264"
		CategoryName = "wheel-and-tyre-accessories"
		SubId = "180090"
		SubName = "snow-chains-anti-skid-and-snow-socks"
	}
	if category_name_cl == "spare-tyre-jack-and-tool-kits" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262264"
		CategoryName = "wheel-and-tyre-accessories"
		SubId = "262265"
		SubName = "spare-tyre-jack-and-tool-kits"
	}
	if category_name_cl == "other-tyre-accessories" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262264"
		CategoryName = "wheel-and-tyre-accessories"
		SubId = "33746"
		SubName = "other-tyre-accessories"
	}
	if category_name_cl == "rim-protectors" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262263"
		CategoryName = "wheel-centre-caps-trims-and-trim-rings"
		SubId = "180123"
		SubName = "rim-protectors"
	}
	if category_name_cl == "wheel-trims-and-trim-rings" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262263"
		CategoryName = "wheel-centre-caps-trims-and-trim-rings"
		SubId = "33744"
		SubName = "wheel-trims-and-trim-rings"
	}
	if category_name_cl == "wheel-centre-caps" {
		IsSubCategory = true
		ParentId = "33743"
		ParentName = "wheels-tyres-and-parts"
		CategoryId = "262263"
		CategoryName = "wheel-centre-caps-trims-and-trim-rings"
		SubId = "43961"
		SubName = "wheel-centre-caps"
	}
	if category_name_cl == "washers-pumps-reservoirs-and-nozzles" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "174111"
		CategoryName = "wipers-and-washers"
		SubId = "174112"
		SubName = "washers-pumps-reservoirs-and-nozzles"
	}
	if category_name_cl == "wiper-arms" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "174111"
		CategoryName = "wipers-and-washers"
		SubId = "174113"
		SubName = "wiper-arms"
	}
	if category_name_cl == "wiper-blades-and-refills" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "174111"
		CategoryName = "wipers-and-washers"
		SubId = "179852"
		SubName = "wiper-blades-and-refills"
	}
	if category_name_cl == "additional-wiper-and-washer-components" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "174111"
		CategoryName = "wipers-and-washers"
		SubId = "262179"
		SubName = "additional-wiper-and-washer-components"
	}
	if category_name_cl == "wiper-linkages-transmissions-and-motors" {
		IsSubCategory = true
		ParentId = "33637"
		ParentName = "exterior-parts-and-accessories"
		CategoryId = "174111"
		CategoryName = "wipers-and-washers"
		SubId = "61941"
		SubName = "wiper-linkages-transmissions-and-motors"
	}
	return IsSubCategory, CategoryId, CategoryName, ParentId, ParentName, SubId, SubName
}
func CleanSubCategories(IsSubCategory bool, CategoryId string, CategoryName string, ParentId string, ParentName string, SubId string, SubName string) (string, string, string, string, string, string) {

	rCategoryId := CategoryId
	rCategoryName := CategoryName
	rParentId := ParentId
	rParentName := ParentName
	rSubId := SubId
	rSubName := SubName

	if IsSubCategory == true {
		rCategoryId = CategoryId
		rCategoryName = CleanCategory(CategoryName)
		rParentId = ParentId
		rParentName = CleanCategory(ParentName)
		rSubId = SubId
		rSubName = CleanCategory(SubName)
	}
	return rCategoryId, rCategoryName, rParentId, rParentName, rSubId, rSubName
}
func SetDisplayImage(PictureURLLarge string, GalleryURL string) string {
	DisplayImage1 := "https://spares2compare.b-cdn.net/no-img/no-img.png"
	switch {
	case len(PictureURLLarge) > 0:
		DisplayImage1 = PictureURLLarge
	case len(GalleryURL) > 0:
		DisplayImage1 = GalleryURL
	}
	DisplayImage := strings.Replace(DisplayImage1, "http:", "https:", -1)

	color.Blue.Println("PictureURLLarge : \t\t", PictureURLLarge)
	color.Blue.Println("GalleryURL : \t\t\t", GalleryURL)
	color.Yellow.Println("DisplayImage -> \t\t", DisplayImage)
	color.Cyan.Print("\n\n")

	return DisplayImage
}
func CreateAffiliateLinksEbay(ItemId string) (string, string) {
	EBAY_CAMPAIGN_ID := os.Getenv("EBAY_CAMPAIGN_ID")
	EBAY_DOMAIN := os.Getenv("EBAY_DOMAIN")
	EBAY_GB_MKR_ID := os.Getenv("EBAY_GB_MKR_ID")
	EBAY_SITE_ID := os.Getenv("EBAY_SITE_ID")
	EBAY_CUSTOM_TRACKING := os.Getenv("EBAY_CUSTOM_TRACKING")

	ebay_affiliate_link := "https://www." + EBAY_DOMAIN + "/itm/" + ItemId + "?mkrid=" + EBAY_GB_MKR_ID + "&siteid=" + EBAY_SITE_ID + "&mkcid=1&campid=" + EBAY_CAMPAIGN_ID + "&toolid=10001&customid=" + EBAY_CUSTOM_TRACKING + "&mkevt=1"
	ebay_affiliate_pixel := ""

	return ebay_affiliate_link, ebay_affiliate_pixel
}
func SetConditionString(conditionid string) string {
	conditionStr := ""
	if conditionid == "1000" {
		conditionStr = "new"
	} else if conditionid == "3000" {
		conditionStr = "used"
	} else if conditionid == "4000" {
		conditionStr = "very-good"
	} else if conditionid == "5000" {
		conditionStr = "good"
	} else if conditionid == "7000" {
		conditionStr = "for-parts-only"
	} else if conditionid == "1500" {
		conditionStr = "new-other"
	} else if conditionid == "1750" {
		conditionStr = "new-with-defects"
	} else if conditionid == "2000" {
		conditionStr = "cert-refurbished"
	} else if conditionid == "2010" {
		conditionStr = "excellent-refurb"
	} else if conditionid == "2020" {
		conditionStr = "very-good-refurb"
	} else if conditionid == "2030" {
		conditionStr = "good-refurb"
	} else if conditionid == "2050" {
		conditionStr = "seller-refurb"
	} else if conditionid == "2750" {
		conditionStr = "like-new"
	} else if conditionid == "6000" {
		conditionStr = "acceptable"
	} else {
		conditionStr = ""
	}
	return conditionStr
}
func CheckHasLogo(seller_username string) string {

	UserNameLogo := ""

	if seller_username == "silverlake-automotive-recycling" {
		UserNameLogo = "silverlake.png"
	}
	if seller_username == "charlestrentlimited" {
		UserNameLogo = "charlestrentlimited.png"
	}
	if seller_username == "hillsmotors" {
		UserNameLogo = "hillsmotors.png"
	}
	if seller_username == "renbreakers" {
		UserNameLogo = "renbreakers.png"
	}
	if seller_username == "365engines" {
		UserNameLogo = "365engines.png"
	}
	if seller_username == "rbsspares" {
		UserNameLogo = "rbsspares.png"
	}
	if seller_username == "easycarpartsltd" {
		UserNameLogo = "easycarpartsltd.png"
	}
	if seller_username == "eurocarparts" {
		UserNameLogo = "eurocarparts.png"
	}
	if seller_username == "evolutioncarspares" {
		UserNameLogo = "evolutioncarspares.png"
	}
	if seller_username == "firstchoicebumpers" {
		UserNameLogo = "firstchoicebumpers.png"
	}
	if seller_username == "gacarspares" {
		UserNameLogo = "gacarspares.png"
	}
	if seller_username == "gmhrecyclers" {
		UserNameLogo = "gmhrecyclers.png"
	}
	if seller_username == "lscltdvauxhall" {
		UserNameLogo = "lscltdvauxhall.png"
	}
	if seller_username == "midlandcarbreakers" {
		UserNameLogo = "midlandcarbreakers.png"
	}
	if seller_username == "morganautosalvage" {
		UserNameLogo = "morganautosalvage.png"
	}
	if seller_username == "scbcarparts" {
		UserNameLogo = "scbcarparts.png"
	}
	if seller_username == "selectvehiclespares2" {
		UserNameLogo = "selectvehiclespares2.png"
	}
	if seller_username == "srsautomotive07817555446" {
		UserNameLogo = "srsautomotive.png"
	}
	if seller_username == "symondsautosalvageltd" {
		UserNameLogo = "symondsautosalvageltd.png"
	}
	if seller_username == "vwbreakersukltd" {
		UserNameLogo = "vwbreakersukltd.png"
	}
	if seller_username == "vansparesuk" {
		UserNameLogo = "vansparesuk.png"
	}
	if seller_username == "vygieauto" {
		UserNameLogo = "vygieauto.png"
	}
	if seller_username == "ecopartcarparts" {
		UserNameLogo = "ecopartcarparts.png"
	}
	if seller_username == "xtremebreakers" {
		UserNameLogo = "xtremebreakers.png"
	}
	if seller_username == "autobitsonlinespares" {
		UserNameLogo = "autobitsonlinespares.png"
	}
	if seller_username == "daautoparts" {
		UserNameLogo = "daautopartsdirect.png"
	}
	if seller_username == "247bumpersltd" {
		UserNameLogo = "247bumpersltd.png"
	}
	if seller_username == "combellackvehiclerecyclersltd" {
		UserNameLogo = "combellackvehiclerecyclersltd.png"
	}
	if seller_username == "carbumpers12" {
		UserNameLogo = "carbumpers12.png"
	}
	if seller_username == "drs-car-parts" {
		UserNameLogo = "drs-car-parts.png"
	}
	if seller_username == "bradgatemotors" {
		UserNameLogo = "bradgatemotors.png"
	}
	if seller_username == "bbautos" {
		UserNameLogo = "bbautos.png"
	}
	if seller_username == "autoverausedbmwparts" {
		UserNameLogo = "autoverausedbmwparts.png"
	}
	if seller_username == "autosaverecycledautoparts" {
		UserNameLogo = "autosaverecycledautoparts.png"
	}
	if seller_username == "aswrindependentbmwdismantlers" {
		UserNameLogo = "aswrindependentbmwdismantlers.png"
	}
	if seller_username == "candccarparts" {
		UserNameLogo = "candccarparts.png"
	}
	if seller_username == "intaparts" {
		UserNameLogo = "intaparts.png"
	}
	if seller_username == "nirroltdvwaudicarparts" {
		UserNameLogo = "nirroltdvwaudicarparts.png"
	}
	if seller_username == "georgeanton" {
		UserNameLogo = "georgeanton.png"
	}
	if seller_username == "german-autospares" {
		UserNameLogo = "ajsparesltd.png"
	}
	if seller_username == "quality.parts.ml6" {
		UserNameLogo = "qualitypartsml6.png"
	}
	if seller_username == "mim-ltd" {
		UserNameLogo = "tradepartsnortheast.png"
	}
	if seller_username == "mollys-car-breakers" {
		UserNameLogo = "ajsparesltd.png"
	}
	if seller_username == "desguace_malvarrosa" {
		UserNameLogo = "desguace_malvarrosa.png"
	}
	if seller_username == "kernowparts" {
		UserNameLogo = "kernowparts.png"
	}
	if seller_username == "desguaces_berrocar" {
		UserNameLogo = "desguacesberrocar.png"
	}
	if seller_username == "desguace_alarcon" {
		UserNameLogo = "desguacealarcon.png"
	}
	if seller_username == "south_hams_car_spares" {
		UserNameLogo = "southhamscarspares.png"
	}
	if seller_username == "quarrymotors" {
		UserNameLogo = "quarrymotors.png"
	}
	if seller_username == "globalparts-uk" {
		UserNameLogo = "globalparts-uk.png"
	}
	if seller_username == "first-autoteile" {
		UserNameLogo = "first-autoteile.png"
	}
	if seller_username == "dronsfieldsindependentmercedes0333" {
		UserNameLogo = "dronsfieldsindependentmercedes0333.png"
	}
	if seller_username == "www_neron-parts_de" {
		UserNameLogo = "www_neron-parts_de.png"
	}
	if seller_username == "daviva16" {
		UserNameLogo = "daviva16.png"
	}
	if seller_username == "premium_parts_online_2" {
		UserNameLogo = "premium-parts-online.png"
	}
	if seller_username == "acarol-autoteile" {
		UserNameLogo = "acarol-autoteile.png"
	}
	if seller_username == "vwpartsltd" {
		UserNameLogo = "vwpartsltd.png"
	}
	if seller_username == "a1germanbreakersltd" {
		UserNameLogo = "a1germanbreakersltd.png"
	}
	if seller_username == "bcgbreakersltd" {
		UserNameLogo = "bcgbreakersltd.png"
	}
	return UserNameLogo
}
