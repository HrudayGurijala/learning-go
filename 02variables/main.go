package main
import "fmt"
// import "mypackage"

const logginstring string = "Logged In"

func main(){
	// fmt.Println("variables")
	var username string 
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	var isLoggedIn bool 
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)
	var smolVal = 123.32
	fmt.Println(smolVal)
	fmt.Printf("variable is of type: %T \n", smolVal)
	
	fmt.Println(logginstring)
	fmt.Printf("variable is of type: %T \n", logginstring)

	// println(mypackage.PublicConstant) // OK: Accessible because it's public
	// println(mypackage.privateConstant) // Error: Not accessible because it's private
	
}