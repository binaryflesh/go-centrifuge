// +build testworld

package testworld

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

const typeInvoice string = "invoice"
const typeEntity string = "entity"
const typePO string = "purchaseorder"
const poPrefix string = "po"

var isRunningOnCI = len(os.Getenv("TRAVIS")) != 0

type httpLog struct {
	logger httpexpect.Logger
}

func (h *httpLog) Logf(fm string, args ...interface{}) {
	if !isRunningOnCI {
		h.logger.Logf(fm, args...)
	}
}

func createInsecureClientWithExpect(t *testing.T, baseURL string) *httpexpect.Expect {
	config := httpexpect.Config{
		BaseURL:  baseURL,
		Client:   createInsecureClient(),
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(&httpLog{t}),
		},
	}
	return httpexpect.WithConfig(config)
}

func getDocumentAndCheck(t *testing.T, e *httpexpect.Expect, auth string, documentType string, params map[string]interface{}, checkattrs bool) *httpexpect.Value {
	docIdentifier := params["document_id"].(string)

	objGet := addCommonHeaders(e.GET("/"+documentType+"/"+docIdentifier), auth).
		Expect().Status(http.StatusOK).JSON().NotNull()
	objGet.Path("$.header.document_id").String().Equal(docIdentifier)
	objGet.Path("$.data.currency").String().Equal(params["currency"].(string))
	if checkattrs {
		attrs := objGet.Path("$.data.attributes").Object().Raw()
		eattrs := defaultAttributePayload()
		cattrs := make(map[string]map[string]string)
		for k, v := range attrs {
			atr := v.(map[string]interface{})
			delete(atr, "key")
			atri := make(map[string]string)
			for k1, v1 := range atr {
				atri[k1] = v1.(string)
			}

			cattrs[k] = atri
		}

		assert.Equal(t, eattrs, cattrs)
	}

	return objGet
}
func getEntityAndCheck(e *httpexpect.Expect, auth string, documentType string, params map[string]interface{}) *httpexpect.Value {
	docIdentifier := params["document_id"].(string)

	objGet := addCommonHeaders(e.GET("/"+documentType+"/"+docIdentifier), auth).
		Expect().Status(http.StatusOK).JSON().NotNull()
	objGet.Path("$.header.document_id").String().Equal(docIdentifier)
	objGet.Path("$.data.entity.legal_name").String().Equal(params["legal_name"].(string))

	return objGet
}

func getEntity(e *httpexpect.Expect, auth string, docIdentifier string) *httpexpect.Value {

	objGet := addCommonHeaders(e.GET("/entity/"+docIdentifier), auth).
		Expect().Status(http.StatusOK).JSON().NotNull()

	return objGet
}

func getEntityWithRelation(e *httpexpect.Expect, auth string, documentType string, params map[string]interface{}) *httpexpect.Value {
	relationshipIdentifier := params["r_identifier"].(string)

	objGet := addCommonHeaders(e.GET("/relationship/"+relationshipIdentifier+"/"+documentType), auth).
		Expect().Status(http.StatusOK).JSON().NotNull()

	return objGet
}

func nonexistentEntityWithRelation(e *httpexpect.Expect, auth string, documentType string, params map[string]interface{}) *httpexpect.Value {
	relationshipIdentifier := params["r_identifier"].(string)

	objGet := addCommonHeaders(e.GET("/relationship/"+relationshipIdentifier+"/"+documentType), auth).
		Expect().Status(500).JSON().NotNull()

	return objGet
}

func revokeEntity(e *httpexpect.Expect, auth, entityID string, status int, payload map[string]interface{}) *httpexpect.Object {
	obj := addCommonHeaders(e.POST("/entity/"+entityID+"/revoke"), auth).
		WithJSON(payload).
		Expect().Status(status).JSON().Object()
	return obj
}

func nonExistingDocumentCheck(e *httpexpect.Expect, auth string, documentType string, params map[string]interface{}) *httpexpect.Value {
	docIdentifier := params["document_id"].(string)

	objGet := addCommonHeaders(e.GET("/"+documentType+"/"+docIdentifier), auth).
		Expect().Status(500).JSON().NotNull()
	return objGet
}

func createDocument(e *httpexpect.Expect, auth string, documentType string, status int, payload map[string]interface{}) *httpexpect.Object {
	obj := addCommonHeaders(e.POST("/"+documentType), auth).
		WithJSON(payload).
		Expect().Status(status).JSON().Object()
	return obj
}

func shareEntity(e *httpexpect.Expect, auth, entityID string, status int, payload map[string]interface{}) *httpexpect.Object {
	obj := addCommonHeaders(e.POST("/entity/"+entityID+"/share"), auth).
		WithJSON(payload).
		Expect().Status(status).JSON().Object()
	return obj
}

func updateDocument(e *httpexpect.Expect, auth string, documentType string, status int, docIdentifier string, payload map[string]interface{}) *httpexpect.Object {
	obj := addCommonHeaders(e.PUT("/"+documentType+"/"+docIdentifier), auth).
		WithJSON(payload).
		Expect().Status(status).JSON().Object()
	return obj
}

func getDocumentIdentifier(t *testing.T, response *httpexpect.Object) string {
	docIdentifier := response.Value("header").Path("$.document_id").String().NotEmpty().Raw()
	if docIdentifier == "" {
		t.Error("docIdentifier empty")
	}
	return docIdentifier
}

func getTransactionID(t *testing.T, resp *httpexpect.Object) string {
	txID := resp.Value("header").Path("$.job_id").String().Raw()
	if txID == "" {
		t.Error("transaction ID empty")
	}

	return txID
}

func getDocumentCurrentVersion(t *testing.T, resp *httpexpect.Object) string {
	versionID := resp.Value("header").Path("$.version").String().Raw()
	if versionID == "" {
		t.Error("version ID empty")
	}

	return versionID
}

func mintUnpaidInvoiceNFT(e *httpexpect.Expect, auth string, httpStatus int, documentID string, payload map[string]interface{}) *httpexpect.Object {
	resp := addCommonHeaders(e.POST("/token/mint/invoice/unpaid/"+documentID), auth).
		WithJSON(payload).
		Expect().Status(httpStatus)

	httpObj := resp.JSON().Object()
	return httpObj
}

func mintNFT(e *httpexpect.Expect, auth string, httpStatus int, payload map[string]interface{}) *httpexpect.Object {
	resp := addCommonHeaders(e.POST("/token/mint"), auth).
		WithJSON(payload).
		Expect().Status(httpStatus)

	httpObj := resp.JSON().Object()
	return httpObj
}

func getProof(e *httpexpect.Expect, auth string, httpStatus int, documentID string, payload map[string]interface{}) *httpexpect.Object {
	resp := addCommonHeaders(e.POST("/document/"+documentID+"/proof"), auth).
		WithJSON(payload).
		Expect().Status(httpStatus)
	return resp.JSON().Object()
}

func getNodeConfig(e *httpexpect.Expect, auth string, httpStatus int) *httpexpect.Object {
	resp := addCommonHeaders(e.GET("/config"), auth).
		Expect().Status(httpStatus)
	return resp.JSON().Object()
}

func getAccount(e *httpexpect.Expect, auth string, httpStatus int, identifier string) *httpexpect.Object {
	resp := addCommonHeaders(e.GET("/accounts/"+identifier), auth).
		Expect().Status(httpStatus)
	return resp.JSON().Object()
}

func getAllAccounts(e *httpexpect.Expect, auth string, httpStatus int) *httpexpect.Object {
	resp := addCommonHeaders(e.GET("/accounts"), auth).
		Expect().Status(httpStatus)
	return resp.JSON().Object()
}

func generateAccount(e *httpexpect.Expect, auth string, httpStatus int) *httpexpect.Object {
	resp := addCommonHeaders(e.POST("/accounts/generate"), auth).
		Expect().Status(httpStatus)
	return resp.JSON().Object()
}

// TODO add rest of the endpoints for config

func createInsecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func getTransactionStatusAndMessage(e *httpexpect.Expect, auth string, txID string) (string, string) {
	emptyResponseTolerance := 5
	emptyResponsesEncountered := 0
	for {
		resp := addCommonHeaders(e.GET("/jobs/"+txID), auth).Expect().Status(200).JSON().Object().Raw()
		status, ok := resp["status"].(string)
		if !ok {
			emptyResponsesEncountered++
			if emptyResponsesEncountered > emptyResponseTolerance {
				panic("transaction api non-responsive")
			}
			time.Sleep(1 * time.Second)
			continue
		}

		if status == "pending" {
			time.Sleep(1 * time.Second)
			continue
		}

		message, ok := resp["message"].(string)

		if !ok {
			message = "Unknown error while processing transaction"
		}

		return status, message
	}
}

func addCommonHeaders(req *httpexpect.Request, auth string) *httpexpect.Request {
	return req.
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithHeader("authorization", auth)
}

func getAccounts(accounts *httpexpect.Array) map[string]string {
	accIDs := make(map[string]string)
	for i := 0; i < int(accounts.Length().Raw()); i++ {
		val := accounts.Element(i).Path("$.identity_id").String().NotEmpty().Raw()
		accIDs[val] = val
	}
	return accIDs
}
