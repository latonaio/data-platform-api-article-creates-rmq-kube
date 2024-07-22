package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-article-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Article:                         *data.Article,
		ArticleType:                     data.ArticleType,
		ArticleOwner:                    data.ArticleOwner,
		ArticleOwnerBusinessPartnerRole: data.ArticleOwnerBusinessPartnerRole,
		PersonResponsible:             	 data.PersonResponsible,
		ValidityStartDate:             	 data.ValidityStartDate,
		ValidityStartTime:             	 data.ValidityStartTime,
		ValidityEndDate:               	 data.ValidityEndDate,
		ValidityEndTime:               	 data.ValidityEndTime,
		Description:                   	 data.Description,
		LongText:                      	 data.LongText,
		Introduction:                  	 data.Introduction,
		Site:                          	 data.Site,
		Shop:                          	 data.Shop,
		Project:                       	 data.Project,
		WBSElement:                    	 data.WBSElement,
		Tag1:                          	 data.Tag1,
		Tag2:                          	 data.Tag2,
		Tag3:                          	 data.Tag3,
		Tag4:                          	 data.Tag4,
		DistributionProfile:           	 data.DistributionProfile,
		QuestionnaireType:			   	 data.QuestionnaireType,
		QuestionnaireTemplate:		   	 data.QuestionnaireTemplate,
		LastChangeDate:                	 data.LastChangeDate,
		LastChangeTime:                	 data.LastChangeTime,
		LastChangeUser:				   	 data.LastChangeUser,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Article:                 *dataHeader.Article,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		Article:        *dataHeader.Article,
		AddressID:      data.AddressID,
		PostalCode:     data.PostalCode,
		LocalSubRegion: data.LocalSubRegion,
		LocalRegion:    data.LocalRegion,
		Country:        data.Country,
		GlobalRegion:   data.GlobalRegion,
		TimeZone:       data.TimeZone,
		District:       data.District,
		StreetName:     data.StreetName,
		CityName:       data.CityName,
		Building:       data.Building,
		Floor:          data.Floor,
		Room:           data.Room,
		XCoordinate:    data.XCoordinate,
		YCoordinate:    data.YCoordinate,
		ZCoordinate:    data.ZCoordinate,
		Site:           data.Site,
	}
}

func ConvertToCounterUpdates(header dpfm_api_input_reader.Header, counter dpfm_api_input_reader.Counter) *CounterUpdates {
	dataHeader := header
	data := counter

	return &CounterUpdates{
		Article:				*dataHeader.Article,
		NumberOfLikes:			data.NumberOfLikes,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
	}
}

func ConvertToLikeUpdates(header dpfm_api_input_reader.Header, like dpfm_api_input_reader.Like) *LikeUpdates {
	dataHeader := header
	data := like

	return &LikeUpdates{
		Article:				*dataHeader.Article,
		BusinessPartner:		data.BusinessPartner,
		Like:					data.Like,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
	}
}
