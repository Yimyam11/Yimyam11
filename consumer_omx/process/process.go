package process

import (
	"consumer/formatomx"
	// "fmt"
	//"os"
	// "log"
	//"strconv"
	// "errors"
	//   "encoding/json"
	//	"io/ioutil"
)

//var ErrOpenFile error = errors.New("no such file or directory")
func getOrderName(orderId int) string {
	switch os := orderId; os {
	case 4:
		return "Remove Offer"
	case 3:
		return "ADD Offer"
	case 41:
		return "Create Memo"
	default:
		return "Unknown"
	}
}

func CheckPkg(p string) string {
	listsoc := []string{"PROSPS06"}
	for _, x := range listsoc {
		if x == p {
			return x
			break
		}
	}
	return "DMC"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Do(c *formatomx.Omxkafka) formatomx.Resultlog {
	var resultOmx formatomx.Resultlog
	var imsi string

	for _, c_ou := range c.Customer.Ou {
		for _, ou_Subscriber := range c_ou.Subscriber {
			for _, c_resINnfo := range ou_Subscriber.ResourceInfo {
				if c_resINnfo.ResourceName == "IMSI" {
					imsi = c_resINnfo.ValuesArray
				}
			}
			for _, ou_Offers := range ou_Subscriber.Offers {
				//if  chkpkg := CheckPkg(ou_Offers.OfferName ); chkpkg == ou_Offers.OfferName    {
				for _, chkExtendinfo := range ou_Offers.ExtendedInfo {
					if chkExtendinfo.Name == "FE_OR_CCBS" { //&& chkExtendinfo.Value ==   "FE" {
						/*  file2, err2 := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND, 0644)
												if err2 != nil {
													log.Fatalf("failed opening file: %s", err2)
												}
												 defer file2.Close()
						 							len, err := file2.WriteString(
												     c.Order.SubmissionDate +
													 ","+ c.OmxtrackingID +
													 ","+ c.Order.Channel +
													 ","+ strconv.Itoa( c.Order.OrderType ) +
													 ","+ c.Order.OrderStatus +
												     ","+  c.Customer.Ou[0].Subscriber[0].SubscriberNumber +
										             ","+  ou_Offers.OfferName   +
													 ","+ chkExtendinfo.Name +
													 ","+ chkExtendinfo.Value +
													  "\n"  )
												 	if err != nil {
												 	log.Fatalf("failed writing to file: %s", err)
												 }
												 fmt.Printf("\nLength: %d bytes", len)
												 fmt.Printf("\nFile Name: %s", file2.Name()) */
						orderidOmx := c.Order.OrderType
						orderName := getOrderName(orderidOmx)
						//  fmt.Printf("order name = %T, %v \n" ,orderName , orderName )
						// fmt.Printf("order id = %T, %v \n" ,orderidOmx , orderidOmx )
						/////////// write log to JSON ////
						resultOmx = formatomx.Resultlog{
							ResultlogDetail: formatomx.ResultlogDetail{
								SubmissionDate: c.Order.SubmissionDate,
								OmxtrackingID:  c.OmxtrackingID,
								OrderType:      orderidOmx, // orderidOmx  	,
								OrderDesc:      orderName,
								Channel:        c.Order.Channel,
								ExtendedName:   chkExtendinfo.Name,
								ExtendedValue:  chkExtendinfo.Value,
								Msisdn:         c.Customer.Ou[0].Subscriber[0].SubscriberNumber,
								Imsi:           imsi,
								// Subcategory	 : "C" ,
								Thidid:      c.Customer.CustomerGeneralInfo.Identification,
								OrderDetail: ou_Offers.OfferName,
							},
						}

						//fmt.Println(result);
					}
				}
				//}
			}
		}
	}
	return resultOmx

}
