package models

type UniversalInfo struct {
	UinPK             *string `json:"uin_pk"`
	UinCatalogLC      *string `json:"uin_catalog_lc"`
	UinAbbreviationLC *string `json:"uin_abbreviation_lc"`
	UinDescLC         *string `json:"uin_desc_lc"`
	UinCatalogEN      *string `json:"uin_catalog_en"`
	UinAbbreviationEN *string `json:"uin_abbreviation_en"`
	UinDescEN         *string `json:"uin_desc_en"`
	UinInfo           *string `json:"uin_info"`
	UinRef            *string `json:"uin_ref"`
}
