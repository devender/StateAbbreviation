package StateAbbreviation

import (
	"strings"
	"log"
)

var usStates = make(map[string]string)

func init() {
	//https://en.wikipedia.org/wiki/Alabama
	usStates["alabama"] = "AL"
	usStates["al"] = "AL"
	usStates["ala"] = "AL"

	//https://en.wikipedia.org/wiki/Alaska
	usStates["alaska"] = "AK"
	usStates["ak"] = "AK"
	usStates["alas"] = "AK"

	//https://en.wikipedia.org/wiki/Arizona
	usStates["arizona"] = "AZ"
	usStates["az"] = "AZ"
	usStates["ariz"] = "AZ"

	//https://en.wikipedia.org/wiki/Arkansas
	usStates["arkansas"] = "AR"
	usStates["ar"] = "AR"
	usStates["ark"] = "AR"

	//https://en.wikipedia.org/wiki/California
	usStates["california"] = "CA"
	usStates["ca"] = "CA"
	usStates["calif"] = "CA"
	usStates["cal"] = "CA"

	//https://en.wikipedia.org/wiki/Colorado
	usStates["colorado"] = "CO"
	usStates["co"] = "CO"
	usStates["colo"] = "CO"
	usStates["col"] = "CO"

	//https://en.wikipedia.org/wiki/Connecticut
	usStates["connecticut"] = "CT"
	usStates["ct"] = "CT"
	usStates["conn"] = "CT"

	//https://en.wikipedia.org/wiki/Delaware
	usStates["delaware"] = "DE"
	usStates["de"] = "DE"
	usStates["del"] = "DE"

	//https://en.wikipedia.org/wiki/Florida
	usStates["florida"] = "FL"
	usStates["fl"] = "FL"
	usStates["fla"] = "FL"
	usStates["flor"] = "FL"

	//https://en.wikipedia.org/wiki/Georgia_(U.S._state)
	usStates["georgia"] = "GA"
	usStates["ga"] = "GA"

	//https://en.wikipedia.org/wiki/Hawaii
	usStates["hawaii"] = "HI"
	usStates["hi"] = "HI"

	//https://en.wikipedia.org/wiki/Idaho
	usStates["idaho"] = "ID"
	usStates["id"] = "ID"
	usStates["ida"] = "ID"

	//https://en.wikipedia.org/wiki/Illinois
	usStates["illinois"] = "IL"
	usStates["il"] = "IL"
	usStates["ill"] = "IL"
	usStates["ills"] = "IL"

	//https://en.wikipedia.org/wiki/Indiana
	usStates["indiana"] = "IN"
	usStates["in"] = "IN"
	usStates["ind"] = "IN"

	//https://en.wikipedia.org/wiki/Iowa
	usStates["iowa"] = "IA"
	usStates["ia"] = "IA"
	usStates["ioa"] = "IA"

	//https://en.wikipedia.org/wiki/Kansas
	usStates["kansas"] = "KS"
	usStates["ks"] = "KS"
	usStates["kans"] = "KS"
	usStates["kan"] = "KS"

	//https://en.wikipedia.org/wiki/Kentucky
	usStates["kentucky"] = "KY"
	usStates["ky"] = "KY"
	usStates["ken"] = "KY"
	usStates["kent"] = "KY"

	//https://en.wikipedia.org/wiki/Louisiana
	usStates["louisiana"] = "LA"
	usStates["la"] = "LA"

	//https://en.wikipedia.org/wiki/Maine
	usStates["maine"] = "ME"
	usStates["me"] = "ME"

	//https://en.wikipedia.org/wiki/Maryland
	usStates["maryland"] = "MD"
	usStates["md"] = "MD"

	//https://en.wikipedia.org/wiki/Massachusetts
	usStates["massachusetts"] = "MA"
	usStates["ma"] = "MA"
	usStates["mass"] = "MA"

	//https://en.wikipedia.org/wiki/Michigan
	usStates["michigan"] = "MI"
	usStates["mi"] = "MI"
	usStates["mich"] = "MI"

	//https://en.wikipedia.org/wiki/Minnesota
	usStates["minnesota"] = "MN"
	usStates["mn"] = "MN"
	usStates["minn"] = "MN"

	//https://en.wikipedia.org/wiki/Mississippi
	usStates["mississippi"] = "MS"
	usStates["ms"] = "MS"
	usStates["miss"] = "MS"

	//https://en.wikipedia.org/wiki/Missouri
	usStates["missouri"] = "MO"
	usStates["mo"] = "MO"

	//https://en.wikipedia.org/wiki/Montana
	usStates["montana"] = "MT"
	usStates["mt"] = "MT"
	usStates["mont"] = "MT"

	//https://en.wikipedia.org/wiki/Nebraska
	usStates["nebraska"] = "NE"
	usStates["ne"] = "NE"
	usStates["neb"] = "NE"
	usStates["nebr"] = "NE"

	//https://en.wikipedia.org/wiki/Nevada
	usStates["nevada"] = "NV"
	usStates["nv"] = "NV"
	usStates["nev"] = "NV"

	//https://en.wikipedia.org/wiki/New_Hampshire
	usStates["new hampshire"] = "NH"
	usStates["nh"] = "NH"

	//https://en.wikipedia.org/wiki/New_Jersey
	usStates["new jersey"] = "NJ"
	usStates["nj"] = "NJ"

	//https://en.wikipedia.org/wiki/New_Mexico
	usStates["new mexico"] = "NM"
	usStates["nm"] = "NM"
	usStates["n mex"] = "NM"
	usStates["new me"] = "NM"

	usStates["new york"] = "NY"
	usStates["ny"] = "NY"

	//https://en.wikipedia.org/wiki/North_Carolina
	usStates["north carolina"] = "NC"
	usStates["nc"] = "NC"
	usStates["n car"] = "NC"

	//https://en.wikipedia.org/wiki/North_Dakota
	usStates["north dakota"] = "ND"
	usStates["nd"] = "ND"
	usStates["n dak"] = "ND"
	usStates["nodak"] = "ND"

	//https://en.wikipedia.org/wiki/Ohio
	usStates["ohio"] = "OH"
	usStates["oh"] = "OH"
	usStates["o"] = "OH"

	//https://en.wikipedia.org/wiki/Oklahoma
	usStates["oklahoma"] = "OK"
	usStates["ok"] = "OK"
	usStates["okla"] = "OK"

	//https://en.wikipedia.org/wiki/Oregon
	usStates["oregon"] = "OR"
	usStates["or"] = "OR"
	usStates["oreg"] = "OR"
	usStates["ore"] = "OR"

	//https://en.wikipedia.org/wiki/Pennsylvania
	usStates["pennsylvania"] = "PA"
	usStates["pa"] = "PA"
	usStates["penn"] = "PA"
	usStates["penna"] = "PA"

	//https://en.wikipedia.org/wiki/Rhode_Island
	usStates["rhode island"] = "RI"
	usStates["ri"] = "RI"
	usStates["r isl"] = "RI"

	//https://en.wikipedia.org/wiki/South_Carolina
	usStates["south carolina"] = "SC"
	usStates["sc"] = "SC"
	usStates["s car"] = "SC"

	//https://en.wikipedia.org/wiki/South_Dakota
	usStates["south dakota"] = "SD"
	usStates["sd"] = "SD"
	usStates["s dak"] = "SD"
	usStates["sodak"] = "SD"

	//https://en.wikipedia.org/wiki/Tennessee
	usStates["tennessee"] = "TN"
	usStates["tenn"] = "TN"
	usStates["tn"] = "TN"

	//https://en.wikipedia.org/wiki/Texas
	usStates["texas"] = "TX"
	usStates["tx"] = "TX"
	usStates["tex"] = "TX"

	//https://en.wikipedia.org/wiki/Utah
	usStates["utah"] = "UT"
	usStates["ut"] = "UT"

	//https://en.wikipedia.org/wiki/Vermont
	usStates["vermont"] = "VT"
	usStates["vt"] = "VT"

	//https://en.wikipedia.org/wiki/Virginia
	usStates["virginia"] = "VA"
	usStates["virg"] = "VA"
	usStates["va"] = "VA"

	//https://en.wikipedia.org/wiki/Washington_(state)
	usStates["washington"] = "WA"
	usStates["wash"] = "WA"
	usStates["wn"] = "WA"
	usStates["wa"] = "WA"

	//https://en.wikipedia.org/wiki/Wisconsin
	usStates["wisconsin"] = "WI"
	usStates["wi"] = "WI"
	usStates["wisc"] = "WI"
	usStates["wis"] = "WI"

	//https://en.wikipedia.org/wiki/West_Virginia
	usStates["west virginia"] = "WV"
	usStates["w va"] = "WV"
	usStates["w virg"] = "WV"
	usStates["wv"] = "WV"

	//https://en.wikipedia.org/wiki/Wyoming
	usStates["wyoming"] = "WY"
	usStates["wi"] = "WY"
	usStates["wy"] = "WY"
	usStates["wyo"] = "WY"

	//https://en.wikipedia.org/wiki/Washington,_D.C.
	usStates["district of columbia"] = "DC"
	usStates["dc"] = "DC"
	usStates["wash dc"] = "DC"
	usStates["washington dc"] = "DC"
}

func GetAbbreviation(state string) string {
	if i, ok := usStates[strings.ToLower(strings.Replace(state, ".", "", -1))]; ok {
		return i
	} else {
		log.Printf("no mapping for %s found \n", state)
		return state
	}
}
