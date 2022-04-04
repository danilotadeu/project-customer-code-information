package codeinformation

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/danilotadeu/r-customer-code-information-provider/app"
	codeinformationModel "github.com/danilotadeu/r-customer-code-information-provider/model/codeinformation"
	"github.com/gofiber/fiber/v2"
)

type apiImpl struct {
	apps *app.Container
}

// NewAPI codeinformation function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	g.Post("/", api.CodeInformationProviderHandler)
}

func (p *apiImpl) CodeInformationProviderHandler(c *fiber.Ctx) error {
	requestCodeInformation := new(codeinformationModel.CustomerRetrieveRequest)
	if err := c.BodyParser(requestCodeInformation); err != nil {
		log.Println("api.codeinformation.codeinformation.codeinformation.body_parser", err.Error())
		return err
	}

	if requestCodeInformation.Header.Security.UsernameToken.Username == "" ||
		requestCodeInformation.Header.Security.UsernameToken.Password.Type == "" {
		responseErr := codeinformationModel.CodeInformationResponseError{}
		responseErr.Body.Fault.Faultstring.Text = "No customer with id 54.033.310 exists in the database."
		responseErr.Body.Fault.Faultcode.Ns0 = "ns0:CMS.RC6701"
		body, _ := xml.Marshal(responseErr)

		return c.Status(http.StatusBadRequest).Type("xml").SendString(string(body))
	}
	b, err := ioutil.ReadFile("response.xml")
	if err != nil {
		log.Println("api.codeinformation.codeinformation.codeinformation.ioutil_readfile", err.Error())
		return err
	}

	return c.Status(http.StatusOK).Type("xml").SendString(string(b))
}
