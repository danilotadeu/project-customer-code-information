package codeinformation

import (
	"fmt"

	codeinformationModel "github.com/danilotadeu/r-customer-code-information-provider/model/codeinformation"
	"github.com/valyala/fasthttp"
)

//App is a contract to CodeInformation..
type App interface {
	GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest)
}

type appImpl struct{}

//NewApp init a codeInformation
func NewApp() App {
	return &appImpl{}
}

//GetCodeInformation get a information code in webservice...
func (a appImpl) GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest) {
	fmt.Println(requestCodeInformation)
}
