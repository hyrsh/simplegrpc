package checkconfig

import (
	"log"
	"math/rand"
	"net"
	"os"
	"simpleGRPC/configchanger"
	"simpleGRPC/configstruct"
	"strconv"
	"strings"
	"time"
)

//Init the settings check
func Init() {
	log.Println("Config check ...")

	/*
		Environment check (kubernetes)
	*/
	checkEnvVar()

	/*
		Settings checks
	*/
	checkPortUint(configstruct.CurrentConfig.SimpleGRPC.Settings.Listenport)
	checkIntervall(configstruct.CurrentConfig.SimpleGRPC.Settings.Interval)
	checkAddress(configstruct.CurrentConfig.SimpleGRPC.Settings.Target)
	checkMessage(configstruct.CurrentConfig.SimpleGRPC.Settings.Message)
	checkMessage(configstruct.CurrentConfig.SimpleGRPC.Settings.Answer)
	checkRunType(configstruct.CurrentConfig.SimpleGRPC.Settings.RunType)

	//Success message
	log.Println("Config valid! Ready for launch!")
}

func checkEnvVar() {
	pod := os.Getenv("SGRPC_K8S_PODNAME")
	if pod == "" {
		log.Println("No kubernetes pod name specified!")
	} else {
		log.Println("My name is: " + pod)
		configstruct.CurrentConfig.SimpleGRPC.Settings.Message = configstruct.CurrentConfig.SimpleGRPC.Settings.Message + " (" + pod + ")"
		configstruct.CurrentConfig.SimpleGRPC.Settings.Answer = configstruct.CurrentConfig.SimpleGRPC.Settings.Answer + "(" + pod + ")"
	}
}

func checkRunType(t string) {
	if t != "client" && t != "server" {
		log.Fatal("Run type not recognized! (client or server)")
	}
}

func checkIntervall(d string) {
	_, intervallErr := time.ParseDuration(d)
	if intervallErr != nil {
		log.Fatal(intervallErr)
	}
}

func checkMessage(msg string) {
	const keylen int = 20            //character length of random generated msg
	var dummyMSG string              //init > empty
	rand.Seed(time.Now().UnixNano()) //get a pseudo-secure seed
	//default uppercase + numeric + "-"
	pool := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-"}
	//pseudo random composition of the ID (no UUID)
	if msg == "" || msg == "none" {
		for i := 0; i < keylen; i++ {
			p := rand.Intn(len(pool))
			dummyMSG += pool[p]
		}

		configstruct.CurrentConfig.SimpleGRPC.Settings.Message = "Hi from " + dummyMSG

		//re-write config
		configchanger.UpdateConfig("Message")
	}
}

//check given string after conversion to uint16 for validity of range and exceptions
func checkPortStr(portStr string) {
	//we want to restrict some port allocations
	portArray := [2]uint16{0, 22}                       //no 0, no SSH
	port, portErr := strconv.ParseUint(portStr, 10, 16) //parse for uint16 decimal (base 10) representation, since the port range is numerically based on it
	if portErr != nil {
		log.Println(1, "Port is invalid "+portStr, false)
		log.Println(portErr)
	}
	for _, v := range portArray {
		if uint16(port) == v { //type casting is allowed because of strconv.ParseUint
			log.Println(2, "Port "+portStr+" is not allowed to set!", true)
		}
	}
}

func checkPortUint(port uint16) {
	if port == 0 {
		configstruct.CurrentConfig.SimpleGRPC.Settings.Listenport = 8080
		configchanger.UpdateConfig("Listenport")
	} else {
		//we want to restrict some port allocations
		portArray := [2]uint16{0, 22} //no 0, no SSH
		for _, v := range portArray {
			if port == v {
				log.Println("Use of forbidden port detected (0, 22)!")
				os.Exit(1)
			}
		}
	}
}

//check IP address validity
func checkIP(ipadr string) {
	ip := net.ParseIP(ipadr)
	if ip == nil {
		log.Println(1, "Cannot parse IP as numeric ("+ipadr+"). If you use an URL it is up to you that it is correct.", false)
	}
}

func checkAddress(addr string) {
	div := strings.Split(addr, ":") //I know its stupid
	if len(div) != 2 {              //I want people to enter a valid port if they provide an URL
		log.Println(2, "Target not parsable ("+div[0]+"). If you use an URL append the port (e.g.: https://my-destination-url.io:6443)", true)
	} else { //default checks
		checkIP(div[0])
		checkPortStr(div[1])
	}

}
