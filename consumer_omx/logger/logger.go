package logger
import  (
		"log"
		"os"
		"fmt"
		)
func Exists(name string) bool {
			_, err := os.Stat(name)
			return !os.IsNotExist(err)
 }		
func LogFile ( c []byte, fname string ) {
/// validate input ///

	if fname != "" {
		// check file ///
		 
		if Exists(fname) {
			fmt.Println("found file")
		}else {
		     fmt.Println("nof found file")
		}

		//write file 
		 // If the file doesn't exist, create it, or append to the file
		file, err2 := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
		if err2 != nil {
			log.Fatalf("failed opening file: %s", err2)
		}
		 defer file.Close()
		stringdata := string(c)
		len, err := file.WriteString(stringdata+"\n")
		/* 	c.Order.SubmissionDate +
		//","+ c.OmxtrackingID +
		//","+ c.Order.Channel +
		//","+ strconv.Itoa( c.Order.OrderType ) +
		//","+ c.Order.OrderStatus +
	//	","+  c.Customer.Ou[0].Subscriber[0].SubscriberNumber +
	//	","+  ou_Offers.OfferName   +
	//	","+ chkExtendinfo.Name +
	//	","+ chkExtendinfo.Value +
	//	"\n" */
		 
		 if err != nil {
		 log.Fatalf("failed writing to file: %s", err)
		 }
		 fmt.Printf("\nLength: %d bytes", len)
		 fmt.Printf("\nFile Name: %s", file.Name())  
	}	
}	