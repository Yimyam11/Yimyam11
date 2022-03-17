package formatkafka 

type Kafka_message_input struct {
	Order order_inter    `json:"order"` 
	customer  customer   `json:"customer"` 
	OmxtrackingID  string `json:"omxtrackingID"` 
} 

type Order_inter struct {
  Channel   string    `json:"channel"` 
  OrderId   string    `json:"orderId"` 
  OrderType   int    `json:"orderType"` 
  IntegrationMethod   string    `json:"integrationMethod"` 
  OrderStatus   string    `json:"orderStatus"` 
  ProfileStatus   string    `json:"profileStatus"` 
  NetworkStatus   string    `json:"networkStatus"` 
  OrderPriority   string    `json:"orderPriority"`  
  SubmissionDate   string    `json:"submissionDate"` 
  ResponseCode   string    `json:"responseCode"` 
  ResponseMsg   string    `json:"responseMsg"`  
  extendedInfo  []extendedInfo   `json:"extendedInfo"`  
  ProcessConfigExtID   string    `json:"processConfigExtID"`  
  Mnpinfo   mnpinfo     `json:"mnpinfo"`  
     
  
}
  type ExtendedInfo struct {
  	Name       string                 `json:"name"`
	Value       string              `json:"value"`
}

type Mnpinfo struct {} 

type CustomerGeneralInfo struct {
  	Grading      		 string                 `json:"grading"`
	Identification       string              `json:"identification"`
  	IdentificationExpDate  	  string       `json:"identificationExpDate"`
  	IdentificationType       string              `json:"identificationType"`
  	BirthDate       string              `json:"birthDate"`
}

type CustomerTypeInfo struct {
  customerType       int                 `json:"customerType"`
 }
 
 type Account struct {
  AccountId       string              `json:"accountId"`
  AccountingManagementInfo       accountingManagementInfo        `json:"accountingManagementInfo"`
  agreementRefId       string              `json:"agreementRefId"`
  payChannelId      int              `json:"payChannelId"`
  refId             string              `json:"refId"`
  extendedInfo       []extendedInfo              `json:"extendedInfo"` 
 }

 type Subscriber struct {
	Status       int              `json:"status"`
	AccountRefId      string              `json:"accountRefId"`
	ActivityInfo activityInfo     `json:"activityInfo"`
	Offers  []offers 	 `json:"offers"`
	PayChannelIdPrimary      string              `json:"payChannelIdPrimary"`
	RefId      string              `json:"refId"`
    ResourceInfo []resourceInfo    `json:"resourceInfo"`
    SubscriberGeneralInfo  subscriberGeneralInfo   `json:"subscriberGeneralInfo"`
	SubscriberId      string              `json:"subscriberId"`
	SubscriberNumber      string              `json:"subscriberNumber"`
	SubscriberType      string              `json:"subscriberType"`
	ResponseCode      string              `json:"responseCode"`
	ResponseMsg      string              `json:"responseMsg"`
	ExtendedInfo       []extendedInfo              `json:"extendedInfo"` 
 
 }
 type Ou struct {
	Agreement       agreement              `json:"agreement"`
	NumberOfIDD      int              `json:"numberOfIDD"`
	NumberOfIR      int              `json:"numberOfIR"`
	OuId      string              `json:"ouId"`
	RefId      string              `json:"refId"`
	Subscriber []subscriber 	 `json:"subscriber"`
   }


 type AccountingManagementInfo struct {
  
	CompanyCode       string              `json:"companyCode"`
  creditClass       string       `json:"creditClass"`
  creditLimitWaiverInd       string              `json:"creditLimitWaiverInd"`
  personalCreditLimit       int              `json:"personalCreditLimit"`
  accountSubType       string              `json:"accountSubType"`
  tempCreditLimit       int              `json:"tempCreditLimit"`
  tempCreditLimitExpDate       string     `json:"tempCreditLimitExpDate"`
 }

 type Agreement struct {
  AgreementId       string                `json:"agreementId"`
  AgreementType       string              `json:"agreementType"`
  RefId       string              `json:"refId"`
 }

 type ActivityInfo struct {
  ActivityReason       string              `json:"activityReason"`
  UserText       string              `json:"userText"`
 
 }
 type Offers  struct {
  action       string              `json:"action"`
  offerName       string              `json:"offerName"`
  offerInstanceId       int              `json:"offerInstanceId"`
  serviceType       int              `json:"serviceType"`
  soc       string              `json:"soc"`
  ExtendedInfo       []extendedInfo              `json:"extendedInfo"`
}
type ResourceInfo struct {
	ResourceName       string              `json:"resourceName"`
	ValuesArray       string              `json:"valuesArray"`
}
 
type SubscriberGeneralInfo struct {
	Language   string              `json:"language"`
 
  }

 


  type Customer struct {
	CustomerId   string              `json:"customerId"`
	RefId   string              `json:"refId"`
	BillCycleNo   int              `json:"billCycleNo"`
	CustomerGeneralInfo   customerGeneralInfo  `json:"customerGeneralInfo"`
	CustomerTypeInfo   customerTypeInfo  `json:"customerTypeInfo"`
	Account   []account   `json:"account"`
	Ou			[]ou   `json:"ou"`
	ExtendedInfo       []extendedInfo              `json:"extendedInfo"`
 }

 