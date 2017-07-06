package routes

import (
	"../config"
)

type RouterRequest struct {
	zarasa string
}

type RouterResponse struct {
	zarasa string
}

func Router(c chan RouterRequest, reglas []config.Regla) {

}
