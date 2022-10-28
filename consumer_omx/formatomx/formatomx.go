package formatomx

type Omxkafka struct {
	Order struct {
		Channel           string `json:"channel"`
		OrderID           string `json:"orderId"`
		OrderType         int    `json:"orderType"`
		IntegrationMethod string `json:"integrationMethod"`
		OrderStatus       string `json:"orderStatus"`
		ProfileStatus     string `json:"profileStatus"`
		NetworkStatus     string `json:"networkStatus"`
		OrderPriority     string `json:"orderPriority"`
		SubmissionDate    string `json:"submissionDate"`
		ResponseCode      string `json:"responseCode"`
		ResponseMsg       string `json:"responseMsg"`
		ExtendedInfo      []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"extendedInfo"`
		ProcessConfigExtID string `json:"processConfigExtID"`
		Mnpinfo            struct {
		} `json:"mnpinfo"`
	} `json:"order"`
	Customer struct {
		CustomerID          string `json:"customerId"`
		RefID               string `json:"refId"`
		BillCycleNo         int    `json:"billCycleNo"`
		CustomerGeneralInfo struct {
			Grading               string `json:"grading"`
			Identification        string `json:"identification"`
			IdentificationExpDate string `json:"identificationExpDate"`
			IdentificationType    string `json:"identificationType"`
			BirthDate             string `json:"birthDate"`
		} `json:"customerGeneralInfo"`
		CustomerTypeInfo struct {
			CustomerType int `json:"customerType"`
		} `json:"customerTypeInfo"`
		Account []struct {
			AccountID                string `json:"accountId"`
			AccountingManagementInfo struct {
				CompanyCode          string  `json:"companyCode"`
				CreditClass          string  `json:"creditClass"`
				CreditLimitWaiverInd string  `json:"creditLimitWaiverInd"`
				PersonalCreditLimit  float64 `json:"personalCreditLimit"`
				AccountSubType       string  `json:"accountSubType"`
				TempCreditLimit      float64 `json:"tempCreditLimit"`
			} `json:"accountingManagementInfo"`
			AgreementRefID string `json:"agreementRefId"`
			PayChannelID   string `json:"payChannelId"`
			RefID          string `json:"refId"`
			ExtendedInfo   []struct {
				Name  string `json:"name"`
				Value string `json:"value,omitempty"`
			} `json:"extendedInfo"`
		} `json:"account"`
		Ou []struct {
			Agreement struct {
				AgreementID   string `json:"agreementId"`
				AgreementType string `json:"agreementType"`
				RefID         string `json:"refId"`
			} `json:"agreement"`
			NumberOfIDD int    `json:"numberOfIDD"`
			NumberOfIR  int    `json:"numberOfIR"`
			OuID        int    `json:"ouId"`
			RefID       string `json:"refId"`
			Subscriber  []struct {
				Status       int    `json:"status"`
				AccountRefID string `json:"accountRefId"`
				ActivityInfo struct {
					ActivityReason string `json:"activityReason"`
					UserText       string `json:"userText"`
				} `json:"activityInfo"`
				Offers []struct {
					Action          string `json:"action"`
					OfferName       string `json:"offerName"`
					OfferInstanceID int    `json:"offerInstanceId"`
					ServiceType     int    `json:"serviceType"`
					Soc             string `json:"soc"`
					ExtendedInfo    []struct {
						Name  string `json:"name"`
						Value string `json:"value,omitempty"`
					} `json:"extendedInfo"`
					ParentOfferInstanceID int `json:"parentOfferInstanceId,omitempty"`
					RelatedOffersArray    []struct {
						OfferName   string `json:"offerName"`
						ServiceType int    `json:"serviceType"`
						Soc         string `json:"soc"`
					} `json:"relatedOffersArray,omitempty"`
					OfferParameterInfo []struct {
						ParamName     string `json:"paramName"`
						ValuesArray   string `json:"valuesArray"`
						EffectiveDate string `json:"effectiveDate"`
					} `json:"offerParameterInfo,omitempty"`
				} `json:"offers"`
				PayChannelIDPrimary string `json:"payChannelIdPrimary"`
				RefID               string `json:"refId"`
				ResourceInfo        []struct {
					ResourceName     string `json:"resourceName,omitempty"`
					ValuesArray      string `json:"valuesArray,omitempty"`
					ResourceCategory string `json:"resourceCategory,omitempty"`
				} `json:"resourceInfo"`
				SubscriberGeneralInfo struct {
					Language string `json:"language"`
				} `json:"subscriberGeneralInfo"`
				SubscriberID     string `json:"subscriberId"`
				SubscriberNumber string `json:"subscriberNumber"`
				SubscriberType   string `json:"subscriberType"`
				ResponseCode     string `json:"responseCode"`
				ResponseMsg      string `json:"responseMsg"`
				ExtendedInfo     []struct {
					Name  string `json:"name"`
					Value string `json:"value,omitempty"`
				} `json:"extendedInfo"`
			} `json:"subscriber"`
		} `json:"ou"`
		ExtendedInfo []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"extendedInfo"`
	} `json:"customer"`
	OmxtrackingID string `json:"omxtrackingID"`
}

type ResourceIn []struct {
	ResourceName     string `json:"resourceName,omitempty"`
	ValuesArray      string `json:"valuesArray,omitempty"`
	ResourceCategory string `json:"resourceCategory,omitempty"`
}

type ResultlogDetail struct {
	SubmissionDate string `json:"submissionDate,omitempty"`
	OmxtrackingID  string `json:"omxtrackingID,omitempty"`
	OrderType      int    `json:"orderType,omitempty"`
	OrderDesc      string `json:"orderDesc,omitempty"`
	Channel        string `json:"channel,omitempty"`
	ExtendedName   string `json:"extendedName,omitempty"`
	ExtendedValue  string `json:"extendedValue,omitempty"`
	Msisdn         string `json:"msisdn,omitempty"`
	Imsi           string `json:"imsi,omitempty"`
	//	Subcategory	 string `json:"subcategory,omitempty"`
	Thidid      string `json:"thaiid,omitempty"`
	OrderDetail string `json:"orderDetail,omitempty"`
}

type Resultlog struct {
	ResultlogDetail `json:"ResuletOMX"`
}
