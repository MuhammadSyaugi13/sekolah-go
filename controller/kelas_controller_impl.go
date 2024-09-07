package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MuhammadSyaugi13/go-project-1/helper"
	"github.com/MuhammadSyaugi13/go-project-1/model/web"
	"github.com/MuhammadSyaugi13/go-project-1/service"
	"github.com/julienschmidt/httprouter"
)

type KelasControllerImpl struct {
	KelasService service.KelasService
}

func (controller *KelasControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)

	kelasCreateRequest := web.KelasCreateRequest{}
	err := decoder.Decode(&kelasCreateRequest)
	helper.PanicIfError(err)

	kelasResponse := controller.KelasService.Create(request.Context(), kelasCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kelasResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
