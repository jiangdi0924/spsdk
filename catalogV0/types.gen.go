// Package catalogV0 provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package catalogV0

// ASINIdentifier defines model for ASINIdentifier.
type ASINIdentifier struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN"`

	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`
}

// AttributeSetList defines model for AttributeSetList.
type AttributeSetList []AttributeSetListType

// AttributeSetListType defines model for AttributeSetListType.
type AttributeSetListType struct {

	// The actor attributes of the item.
	Actor *[]string `json:"Actor,omitempty"`

	// The artist attributes of the item.
	Artist *[]string `json:"Artist,omitempty"`

	// The aspect ratio attribute of the item.
	AspectRatio *string `json:"AspectRatio,omitempty"`

	// The audience rating attribute of the item.
	AudienceRating *string `json:"AudienceRating,omitempty"`

	// The author attributes of the item.
	Author *[]string `json:"Author,omitempty"`

	// The back finding attribute of the item.
	BackFinding *string `json:"BackFinding,omitempty"`

	// The band material type attribute of the item.
	BandMaterialType *string `json:"BandMaterialType,omitempty"`

	// The binding attribute of the item.
	Binding *string `json:"Binding,omitempty"`

	// The Bluray region attribute of the item.
	BlurayRegion *string `json:"BlurayRegion,omitempty"`

	// The brand attribute of the item.
	Brand *string `json:"Brand,omitempty"`

	// The CERO age rating attribute of the item.
	CeroAgeRating *string `json:"CeroAgeRating,omitempty"`

	// The chain type attribute of the item.
	ChainType *string `json:"ChainType,omitempty"`

	// The clasp type attribute of the item.
	ClaspType *string `json:"ClaspType,omitempty"`

	// The color attribute of the item.
	Color *string `json:"Color,omitempty"`

	// The CPU manufacturer attribute of the item.
	CpuManufacturer *string `json:"CpuManufacturer,omitempty"`

	// The decimal value and unit.
	CpuSpeed *DecimalWithUnits `json:"CpuSpeed,omitempty"`

	// The CPU type attribute of the item.
	CpuType *string `json:"CpuType,omitempty"`

	// The creator attributes of the item.
	Creator *[]CreatorType `json:"Creator,omitempty"`

	// The department attribute of the item.
	Department *string `json:"Department,omitempty"`

	// The director attributes of the item.
	Director *[]string `json:"Director,omitempty"`

	// The decimal value and unit.
	DisplaySize *DecimalWithUnits `json:"DisplaySize,omitempty"`

	// The edition attribute of the item.
	Edition *string `json:"Edition,omitempty"`

	// The episode sequence attribute of the item.
	EpisodeSequence *string `json:"EpisodeSequence,omitempty"`

	// The ESRB age rating attribute of the item.
	EsrbAgeRating *string `json:"EsrbAgeRating,omitempty"`

	// The feature attributes of the item
	Feature *[]string `json:"Feature,omitempty"`

	// The flavor attribute of the item.
	Flavor *string `json:"Flavor,omitempty"`

	// The format attributes of the item.
	Format *[]string `json:"Format,omitempty"`

	// The gem type attributes of the item.
	GemType *[]string `json:"GemType,omitempty"`

	// The genre attribute of the item.
	Genre *string `json:"Genre,omitempty"`

	// The golf club flex attribute of the item.
	GolfClubFlex *string `json:"GolfClubFlex,omitempty"`

	// The decimal value and unit.
	GolfClubLoft *DecimalWithUnits `json:"GolfClubLoft,omitempty"`

	// The hand orientation attribute of the item.
	HandOrientation *string `json:"HandOrientation,omitempty"`

	// The hard disk interface attribute of the item.
	HardDiskInterface *string `json:"HardDiskInterface,omitempty"`

	// The decimal value and unit.
	HardDiskSize *DecimalWithUnits `json:"HardDiskSize,omitempty"`

	// The hardware platform attribute of the item.
	HardwarePlatform *string `json:"HardwarePlatform,omitempty"`

	// The hazardous material type attribute of the item.
	HazardousMaterialType *string `json:"HazardousMaterialType,omitempty"`

	// The adult product attribute of the item.
	IsAdultProduct *bool `json:"IsAdultProduct,omitempty"`

	// The autographed attribute of the item.
	IsAutographed *bool `json:"IsAutographed,omitempty"`

	// The is eligible for trade in attribute of the item.
	IsEligibleForTradeIn *bool `json:"IsEligibleForTradeIn,omitempty"`

	// The is memorabilia attribute of the item.
	IsMemorabilia *bool `json:"IsMemorabilia,omitempty"`

	// The issues per year attribute of the item.
	IssuesPerYear *string `json:"IssuesPerYear,omitempty"`

	// The dimension type attribute of an item.
	ItemDimensions *DimensionType `json:"ItemDimensions,omitempty"`

	// The item part number attribute of the item.
	ItemPartNumber *string `json:"ItemPartNumber,omitempty"`

	// The label attribute of the item.
	Label *string `json:"Label,omitempty"`

	// The languages attribute of the item.
	Languages *[]LanguageType `json:"Languages,omitempty"`

	// The legal disclaimer attribute of the item.
	LegalDisclaimer *string `json:"LegalDisclaimer,omitempty"`

	// The price attribute of the item.
	ListPrice *Price `json:"ListPrice,omitempty"`

	// The manufacturer attribute of the item.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// The decimal value and unit.
	ManufacturerMaximumAge *DecimalWithUnits `json:"ManufacturerMaximumAge,omitempty"`

	// The decimal value and unit.
	ManufacturerMinimumAge *DecimalWithUnits `json:"ManufacturerMinimumAge,omitempty"`

	// The manufacturer parts warranty description attribute of the item.
	ManufacturerPartsWarrantyDescription *string `json:"ManufacturerPartsWarrantyDescription,omitempty"`

	// The material type attributes of the item.
	MaterialType *[]string `json:"MaterialType,omitempty"`

	// The decimal value and unit.
	MaximumResolution *DecimalWithUnits `json:"MaximumResolution,omitempty"`

	// The media type attributes of the item.
	MediaType *[]string `json:"MediaType,omitempty"`

	// The metal stamp attribute of the item.
	MetalStamp *string `json:"MetalStamp,omitempty"`

	// The metal type attribute of the item.
	MetalType *string `json:"MetalType,omitempty"`

	// The model attribute of the item.
	Model *string `json:"Model,omitempty"`

	// The number of discs attribute of the item.
	NumberOfDiscs *int `json:"NumberOfDiscs,omitempty"`

	// The number of issues attribute of the item.
	NumberOfIssues *int `json:"NumberOfIssues,omitempty"`

	// The number of items attribute of the item.
	NumberOfItems *int `json:"NumberOfItems,omitempty"`

	// The number of pages attribute of the item.
	NumberOfPages *int `json:"NumberOfPages,omitempty"`

	// The number of tracks attribute of the item.
	NumberOfTracks *int `json:"NumberOfTracks,omitempty"`

	// The operating system attributes of the item.
	OperatingSystem *[]string `json:"OperatingSystem,omitempty"`

	// The decimal value and unit.
	OpticalZoom *DecimalWithUnits `json:"OpticalZoom,omitempty"`

	// The dimension type attribute of an item.
	PackageDimensions *DimensionType `json:"PackageDimensions,omitempty"`

	// The package quantity attribute of the item.
	PackageQuantity *int `json:"PackageQuantity,omitempty"`

	// The part number attribute of the item.
	PartNumber *string `json:"PartNumber,omitempty"`

	// The PEGI rating attribute of the item.
	PegiRating *string `json:"PegiRating,omitempty"`

	// The platform attributes of the item.
	Platform *[]string `json:"Platform,omitempty"`

	// The processor count attribute of the item.
	ProcessorCount *int `json:"ProcessorCount,omitempty"`

	// The product group attribute of the item.
	ProductGroup *string `json:"ProductGroup,omitempty"`

	// The product type name attribute of the item.
	ProductTypeName *string `json:"ProductTypeName,omitempty"`

	// The product type subcategory attribute of the item.
	ProductTypeSubcategory *string `json:"ProductTypeSubcategory,omitempty"`

	// The publication date attribute of the item.
	PublicationDate *string `json:"PublicationDate,omitempty"`

	// The publisher attribute of the item.
	Publisher *string `json:"Publisher,omitempty"`

	// The region code attribute of the item.
	RegionCode *string `json:"RegionCode,omitempty"`

	// The release date attribute of the item.
	ReleaseDate *string `json:"ReleaseDate,omitempty"`

	// The ring size attribute of the item.
	RingSize *string `json:"RingSize,omitempty"`

	// The decimal value and unit.
	RunningTime *DecimalWithUnits `json:"RunningTime,omitempty"`

	// The scent attribute of the item.
	Scent *string `json:"Scent,omitempty"`

	// The season sequence attribute of the item.
	SeasonSequence *string `json:"SeasonSequence,omitempty"`

	// The Seikodo product code attribute of the item.
	SeikodoProductCode *string `json:"SeikodoProductCode,omitempty"`

	// The shaft material attribute of the item.
	ShaftMaterial *string `json:"ShaftMaterial,omitempty"`

	// The size attribute of the item.
	Size *string `json:"Size,omitempty"`

	// The size per pearl attribute of the item.
	SizePerPearl *string `json:"SizePerPearl,omitempty"`

	// The image attribute of the item.
	SmallImage *Image `json:"SmallImage,omitempty"`

	// The studio attribute of the item.
	Studio *string `json:"Studio,omitempty"`

	// The decimal value and unit.
	SubscriptionLength *DecimalWithUnits `json:"SubscriptionLength,omitempty"`

	// The decimal value and unit.
	SystemMemorySize *DecimalWithUnits `json:"SystemMemorySize,omitempty"`

	// The system memory type attribute of the item.
	SystemMemoryType *string `json:"SystemMemoryType,omitempty"`

	// The theatrical release date attribute of the item.
	TheatricalReleaseDate *string `json:"TheatricalReleaseDate,omitempty"`

	// The title attribute of the item.
	Title *string `json:"Title,omitempty"`

	// The decimal value and unit.
	TotalDiamondWeight *DecimalWithUnits `json:"TotalDiamondWeight,omitempty"`

	// The decimal value and unit.
	TotalGemWeight *DecimalWithUnits `json:"TotalGemWeight,omitempty"`

	// The warranty attribute of the item.
	Warranty *string `json:"Warranty,omitempty"`

	// The price attribute of the item.
	WeeeTaxValue *Price `json:"WeeeTaxValue,omitempty"`
}

// Categories defines model for Categories.
type Categories struct {

	// The identifier for the product category (or browse node).
	ProductCategoryId *string `json:"ProductCategoryId,omitempty"`

	// The name of the product category (or browse node).
	ProductCategoryName *string `json:"ProductCategoryName,omitempty"`

	// The parent product category.
	Parent *map[string]interface{} `json:"parent,omitempty"`
}

// CreatorType defines model for CreatorType.
type CreatorType struct {

	// The role of the value.
	Role *string `json:"Role,omitempty"`

	// The value of the attribute.
	Value *string `json:"value,omitempty"`
}

// DecimalWithUnits defines model for DecimalWithUnits.
type DecimalWithUnits struct {

	// The unit of the decimal value.
	Units *string `json:"Units,omitempty"`

	// The decimal value.
	Value *float32 `json:"value,omitempty"`
}

// DimensionType defines model for DimensionType.
type DimensionType struct {

	// The decimal value and unit.
	Height *DecimalWithUnits `json:"Height,omitempty"`

	// The decimal value and unit.
	Length *DecimalWithUnits `json:"Length,omitempty"`

	// The decimal value and unit.
	Weight *DecimalWithUnits `json:"Weight,omitempty"`

	// The decimal value and unit.
	Width *DecimalWithUnits `json:"Width,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional information that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// GetCatalogItemResponse defines model for GetCatalogItemResponse.
type GetCatalogItemResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// An item in the Amazon catalog.
	Payload *Item `json:"payload,omitempty"`
}

// IdentifierType defines model for IdentifierType.
type IdentifierType struct {
	MarketplaceASIN *ASINIdentifier      `json:"MarketplaceASIN,omitempty"`
	SKUIdentifier   *SellerSKUIdentifier `json:"SKUIdentifier,omitempty"`
}

// Image defines model for Image.
type Image struct {

	// The decimal value and unit.
	Height *DecimalWithUnits `json:"Height,omitempty"`

	// The image URL attribute of the item.
	URL *string `json:"URL,omitempty"`

	// The decimal value and unit.
	Width *DecimalWithUnits `json:"Width,omitempty"`
}

// Item defines model for Item.
type Item struct {

	// A list of attributes for the item.
	AttributeSets *AttributeSetList `json:"AttributeSets,omitempty"`
	Identifiers   IdentifierType    `json:"Identifiers"`

	// A list of variation relationship information, if applicable for the item.
	Relationships *RelationshipList `json:"Relationships,omitempty"`

	// A list of sales rank information for the item by category.
	SalesRankings *SalesRankList `json:"SalesRankings,omitempty"`
}

// ItemList defines model for ItemList.
type ItemList []Item

// LanguageType defines model for LanguageType.
type LanguageType struct {

	// The audio format attribute of the item.
	AudioFormat *string `json:"AudioFormat,omitempty"`

	// The name attribute of the item.
	Name *string `json:"Name,omitempty"`

	// The type attribute of the item.
	Type *string `json:"Type,omitempty"`
}

// ListCatalogCategoriesResponse defines model for ListCatalogCategoriesResponse.
type ListCatalogCategoriesResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList        `json:"errors,omitempty"`
	Payload *ListOfCategories `json:"payload,omitempty"`
}

// ListCatalogItemsResponse defines model for ListCatalogItemsResponse.
type ListCatalogItemsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList                 `json:"errors,omitempty"`
	Payload *ListMatchingItemsResponse `json:"payload,omitempty"`
}

// ListMatchingItemsResponse defines model for ListMatchingItemsResponse.
type ListMatchingItemsResponse struct {

	// A list of items.
	Items *ItemList `json:"Items,omitempty"`
}

// ListOfCategories defines model for ListOfCategories.
type ListOfCategories []Categories

// NumberOfOfferListingsList defines model for NumberOfOfferListingsList.
type NumberOfOfferListingsList []OfferListingCountType

// OfferListingCountType defines model for OfferListingCountType.
type OfferListingCountType struct {

	// The number of offer listings.
	Count int32 `json:"Count"`

	// The condition of the item.
	Condition string `json:"condition"`
}

// Price defines model for Price.
type Price struct {

	// The amount.
	Amount *float32 `json:"Amount,omitempty"`

	// The currency code of the amount.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
}

// QualifiersType defines model for QualifiersType.
type QualifiersType struct {

	// The fulfillment channel for the item. Possible values:
	//
	// * Amazon - Fulfilled by Amazon.
	// * Merchant - Fulfilled by the seller.
	FulfillmentChannel string `json:"FulfillmentChannel"`

	// The condition of the item. Possible values: New, Used, Collectible, Refurbished, or Club.
	ItemCondition string `json:"ItemCondition"`

	// The item subcondition for the offer listing. Possible values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	ItemSubcondition string `json:"ItemSubcondition"`

	// (98-100%, 95-97%, 90-94%, 80-89%, 70-79%, Less than 70%, or Just launched ) – Indicates the percentage of feedback ratings that were positive over the past 12 months.
	SellerPositiveFeedbackRating string           `json:"SellerPositiveFeedbackRating"`
	ShippingTime                 ShippingTimeType `json:"ShippingTime"`

	// Indicates whether the marketplace specified in the request and the location that the item ships from are in the same country. Possible values: True, False, or Unknown.
	ShipsDomestically string `json:"ShipsDomestically"`
}

// RelationshipList defines model for RelationshipList.
type RelationshipList []RelationshipType

// RelationshipType defines model for RelationshipType.
type RelationshipType struct {

	// The color variation of the item.
	Color *string `json:"Color,omitempty"`

	// The edition variation of the item.
	Edition *string `json:"Edition,omitempty"`

	// The flavor variation of the item.
	Flavor *string `json:"Flavor,omitempty"`

	// The gem type variations of the item.
	GemType *[]string `json:"GemType,omitempty"`

	// The golf club flex variation of an item.
	GolfClubFlex *string `json:"GolfClubFlex,omitempty"`

	// The decimal value and unit.
	GolfClubLoft *DecimalWithUnits `json:"GolfClubLoft,omitempty"`

	// The hand orientation variation of an item.
	HandOrientation *string `json:"HandOrientation,omitempty"`

	// The hardware platform variation of an item.
	HardwarePlatform *string         `json:"HardwarePlatform,omitempty"`
	Identifiers      *IdentifierType `json:"Identifiers,omitempty"`

	// The dimension type attribute of an item.
	ItemDimensions *DimensionType `json:"ItemDimensions,omitempty"`

	// The material type variations of an item.
	MaterialType *[]string `json:"MaterialType,omitempty"`

	// The metal type variation of an item.
	MetalType *string `json:"MetalType,omitempty"`

	// The model variation of an item.
	Model *string `json:"Model,omitempty"`

	// The operating system variations of an item.
	OperatingSystem *[]string `json:"OperatingSystem,omitempty"`

	// The package quantity variation of an item.
	PackageQuantity *int `json:"PackageQuantity,omitempty"`

	// The product type subcategory variation of an item.
	ProductTypeSubcategory *string `json:"ProductTypeSubcategory,omitempty"`

	// The ring size variation of an item.
	RingSize *string `json:"RingSize,omitempty"`

	// The scent variation of an item.
	Scent *string `json:"Scent,omitempty"`

	// The shaft material variation of an item.
	ShaftMaterial *string `json:"ShaftMaterial,omitempty"`

	// The size variation of an item.
	Size *string `json:"Size,omitempty"`

	// The size per pearl variation of an item.
	SizePerPearl *string `json:"SizePerPearl,omitempty"`

	// The decimal value and unit.
	TotalDiamondWeight *DecimalWithUnits `json:"TotalDiamondWeight,omitempty"`

	// The decimal value and unit.
	TotalGemWeight *DecimalWithUnits `json:"TotalGemWeight,omitempty"`
}

// SalesRankList defines model for SalesRankList.
type SalesRankList []SalesRankType

// SalesRankType defines model for SalesRankType.
type SalesRankType struct {

	// Identifies the item category from which the sales rank is taken.
	ProductCategoryId string `json:"ProductCategoryId"`

	// The sales rank of the item within the item category.
	Rank int32 `json:"Rank"`
}

// SellerSKUIdentifier defines model for SellerSKUIdentifier.
type SellerSKUIdentifier struct {

	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`

	// The seller identifier submitted for the operation.
	SellerId string `json:"SellerId"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU"`
}

// ShippingTimeType defines model for ShippingTimeType.
type ShippingTimeType struct {

	// (0-2 days, 3-7 days, 8-13 days, or 14 or more days) – Indicates the maximum time within which the item will likely be shipped once an order has been placed.
	Max *string `json:"Max,omitempty"`
}

// ListCatalogCategoriesParams defines parameters for ListCatalogCategories.
type ListCatalogCategoriesParams struct {

	// A marketplace identifier. Specifies the marketplace for the item.
	MarketplaceId string `json:"MarketplaceId"`

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`

	// Used to identify items in the given marketplace. SellerSKU is qualified by the seller's SellerId, which is included with every operation that you submit.
	SellerSKU *string `json:"SellerSKU,omitempty"`
}

// ListCatalogItemsParams defines parameters for ListCatalogItems.
type ListCatalogItemsParams struct {

	// A marketplace identifier. Specifies the marketplace for which items are returned.
	MarketplaceId string `json:"MarketplaceId"`

	// Keyword(s) to use to search for items in the catalog. Example: 'harry potter books'.
	Query *string `json:"Query,omitempty"`

	// An identifier for the context within which the given search will be performed. A marketplace might provide mechanisms for constraining a search to a subset of potential items. For example, the retail marketplace allows queries to be constrained to a specific category. The QueryContextId parameter specifies such a subset. If it is omitted, the search will be performed using the default context for the marketplace, which will typically contain the largest set of items.
	QueryContextId *string `json:"QueryContextId,omitempty"`

	// Used to identify an item in the given marketplace. SellerSKU is qualified by the seller's SellerId, which is included with every operation that you submit.
	SellerSKU *string `json:"SellerSKU,omitempty"`

	// A 12-digit bar code used for retail packaging.
	UPC *string `json:"UPC,omitempty"`

	// A European article number that uniquely identifies the catalog item, manufacturer, and its attributes.
	EAN *string `json:"EAN,omitempty"`

	// The unique commercial book identifier used to identify books internationally.
	ISBN *string `json:"ISBN,omitempty"`

	// A Japanese article number that uniquely identifies the product, manufacturer, and its attributes.
	JAN *string `json:"JAN,omitempty"`
}

// GetCatalogItemParams defines parameters for GetCatalogItem.
type GetCatalogItemParams struct {

	// A marketplace identifier. Specifies the marketplace for the item.
	MarketplaceId string `json:"MarketplaceId"`
}
