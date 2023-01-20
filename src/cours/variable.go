package cours

import "fmt"

var LeNomDeMaVariablePublic = "LeNomDeMaVariablePublic"
var leNomDeMaVariablePrivate = "leNomDeMaVariablePrivate"

var LeNomDeMaVariablePublicAvecType int = 3
var leNomDeMaVariablePrivateAvecType int = 3

var LeNomDeMaVariablePublicAvecTypeSansValeur int  // nil
var leNomDeMaVariablePrivateAvecTypeSansValeur int // nil

const LeNomDeMaConstantePublic = "LeNomDeMaConstantePublic"
const leNomDeMaConstantePrivate = "leNomDeMaConstantePrivate"

func DisplayVariable() {
	fmt.Println(leNomDeMaVariablePrivate)
	fmt.Println(leNomDeMaVariablePrivateAvecType)
	fmt.Println(leNomDeMaVariablePrivateAvecTypeSansValeur)
	fmt.Println(leNomDeMaConstantePrivate)
}
