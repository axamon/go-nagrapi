package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Critical maps data for critical problems
type Critical2 struct {
	hostname                   string
	servicedescription         string
	modifiedattributes         int
	checkcommand               string
	checkperiod                string
	notificationperiod         string
	checkinterval              float64
	retryinterval              float64
	eventhandler               interface{}
	hasbeenchecked             int
	shouldbescheduled          int
	checkexecutiontime         float64
	checklatency               float64
	checktype                  int
	currentstate               int
	lasthardstate              int
	lasteventid                int
	currenteventid             int
	currentproblemid           int
	lastproblemid              int
	currentattempt             int
	maxattempts                int
	statetype                  int
	laststatechange            int
	lasthardstatechange        int
	lasttimeok                 int
	lasttimewarning            int
	lasttimeunknown            int
	lasttimecritical           int
	pluginoutput               string
	longpluginoutput           interface{}
	performancedata            interface{}
	lastcheck                  int
	nextcheck                  int
	checkoptions               int
	currentnotificationnumber  int
	currentnotificationid      int
	lastnotification           int
	nextnotification           int
	nomorenotifications        int
	notificationsenabled       int
	activechecksenabled        int
	passivechecksenabled       int
	eventhandlerenabled        int
	problemhasbeenacknowledged int
	acknowledgementtype        int
	flapdetectionenabled       int
	failurepredictionenabled   int
	processperformancedata     int
	obsessoverservice          int
	lastupdate                 int
	isflapping                 int
	percentstatechange         float64
	scheduleddowntimedepth     int
}

type Critical struct {
	Hostname           interface{}
	Currentattempt     int
	Servicestatus      int
	Servicedescription string
	Pluginoutput       string
}

var f = flag.String("f", "status.dat", "Path to the status.dat file to parse")
var e = flag.Int("e", 2, "Nagios code to retrieve")

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}

	file := string(b)

	filesplittato := strings.Split(file, "}")

	var servizi []Critical

	//cicla file.
	for _, linea := range filesplittato {
		//fmt.Println(string(b))

		//recupera servicestatus {
		if strings.Contains(linea, "servicestatus {") {
			var servizio Critical
			//fmt.Println(linea)
			linasplittata := strings.Split(linea, "\n")
			for _, sublinea := range linasplittata {
				if strings.ContainsAny(sublinea, "{}") {
					continue
				}
				//fmt.Println(estraivalore(sublinea, "host_name"))
				if strings.Contains(sublinea, "host_name") {
					servizio.Hostname, _ = estraivalore(sublinea, "host_name")
				}
				if strings.Contains(sublinea, "current_state") {
					output, err := estraivalore(sublinea, "current_state")
					if err != nil {
						fmt.Println("sti cazzi")
					}
					servizio.Servicestatus, err = strconv.Atoi(output.(string))
					if err != nil {
						fmt.Println("aristicazzi")
					}

				}
				if strings.Contains(sublinea, "current_attempt") {
					output, _ := estraivalore(sublinea, "current_attempt")
					servizio.Currentattempt, err = strconv.Atoi(output.(string))
					if err != nil {
						fmt.Println("aristicazzi2")
					}
				}
				if strings.Contains(sublinea, "service_description") {
					output, _ := estraivalore(sublinea, "service_description")
					servizio.Servicedescription = output.(string)
				}
				if strings.Contains(sublinea, "plugin_output") {
					output, _ := estraivalore(sublinea, "plugin_output")
					servizio.Pluginoutput = output.(string)
				}

				//fmt.Println(servizio)
				servizi = append(servizi, servizio)
			}

		}

	}
	fmt.Println(len(servizi))
	for _, servizio := range servizi {
		if servizio.Servicestatus == *e && servizio.Currentattempt < 1 {
			fmt.Println(servizio)
		}
	}

	//inserisci campi in mappa
}

func estraivalore(sublinea, chiave string) (value interface{}, err error) {
	elemento := strings.Split(sublinea, "=")
	if strings.Contains(elemento[0], chiave) {

		value = elemento[1]
		//fmt.Println(value)
		return value, nil
	}
	return nil, fmt.Errorf("problema")
}
