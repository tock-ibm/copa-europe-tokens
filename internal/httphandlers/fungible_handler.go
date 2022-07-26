package httphandlers

import (
	"encoding/json"
	"github.com/copa-europe-tokens/internal/tokens"
	"github.com/copa-europe-tokens/pkg/constants"
	"github.com/copa-europe-tokens/pkg/types"
	"github.com/gorilla/mux"
	"github.com/hyperledger-labs/orion-server/pkg/logger"
	"net/http"
)

type fungibleHandler assetsHandler

type RequestContext interface {
	Apply(tokens.Operations, map[string]string) (interface{}, int)
}

type RequestContextMaker func() RequestContext

type DeployContext types.FungibleDeployRequest
type DescribeContext types.FungibleDescribeRequest
type MintContext types.FungibleMintRequest
type TransferContext types.FungibleTransferRequest
type ConsolidateContext types.FungibleConsolidateRequest
type SubmitContext types.FungibleSubmitRequest
type AccountContext types.FungibleAccountRequest

func createResponse(res interface{}, err error, successStatus int) (interface{}, int) {
	if err != nil {
		return types.HttpResponseErr{ErrMsg: err.Error()}, tokens.ErrorStatus(err)
	} else {
		return res, successStatus
	}
}

func (rq *DeployContext) Apply(manager tokens.Operations, _ map[string]string) (interface{}, int) {
	res, err := manager.FungibleDeploy((*types.FungibleDeployRequest)(rq))
	return createResponse(res, err, http.StatusCreated)
}

func (rq *DescribeContext) Apply(manager tokens.Operations, params map[string]string) (interface{}, int) {
	res, err := manager.FungibleDescribe(params["typeId"])
	return createResponse(res, err, http.StatusOK)
}

func (rq *MintContext) Apply(manager tokens.Operations, params map[string]string) (interface{}, int) {
	res, err := manager.FungiblePrepareMint(params["typeId"], (*types.FungibleMintRequest)(rq))
	return createResponse(res, err, http.StatusOK)
}

func (rq *TransferContext) Apply(manager tokens.Operations, params map[string]string) (interface{}, int) {
	res, err := manager.FungiblePrepareTransfer(params["typeId"], (*types.FungibleTransferRequest)(rq))
	return createResponse(res, err, http.StatusOK)
}

func (rq *ConsolidateContext) Apply(manager tokens.Operations, params map[string]string) (interface{}, int) {
	res, err := manager.FungiblePrepareConsolidate(params["typeId"], (*types.FungibleConsolidateRequest)(rq))
	return createResponse(res, err, http.StatusOK)
}

func (rq *SubmitContext) Apply(manager tokens.Operations, _ map[string]string) (interface{}, int) {
	res, err := manager.FungibleSubmitTx((*types.FungibleSubmitRequest)(rq))
	return createResponse(res, err, http.StatusOK)
}

func (rq *AccountContext) Apply(manager tokens.Operations, params map[string]string) (interface{}, int) {
	res, err := manager.FungibleAccounts(params["typeId"], params["owner"], params["account"])
	return createResponse(res, err, http.StatusOK)
}

func FetchContext(request *http.Request, maker RequestContextMaker) (RequestContext, map[string]string, error) {
	// Get rest URI vars
	params := mux.Vars(request)

	// Get optional query parameters
	query := request.URL.Query()
	for key, _ := range query {
		params[key] = query.Get(key)
	}

	context := maker()
	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(context)
	return context, params, err
}

func (d *fungibleHandler) GenericHandler(
	maker RequestContextMaker,
	response http.ResponseWriter,
	request *http.Request,
) {
	context, params, err := FetchContext(request, maker)
	if err != nil {
		SendHTTPResponse(response, http.StatusBadRequest, &types.HttpResponseErr{ErrMsg: err.Error()}, d.lg)
		return
	}

	responseJson, status := context.Apply(d.manager, params)
	SendHTTPResponse(response, status, responseJson, d.lg)
}

func (d *fungibleHandler) AddHandler(path string, method string, maker RequestContextMaker) *mux.Route {
	return d.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		d.GenericHandler(maker, writer, request)
	}).Methods(method)
}

func NewFungibleHandler(manager tokens.Operations, lg *logger.SugarLogger) *fungibleHandler {
	handler := &fungibleHandler{
		router:  mux.NewRouter(),
		manager: manager,
		lg:      lg,
	}

	handler.AddHandler(constants.FungibleDeploy, http.MethodPost, func() RequestContext {
		return &DeployContext{}
	})
	handler.AddHandler(string(constants.FungibleDescribe), http.MethodGet, func() RequestContext {
		return &DescribeContext{}
	})
	handler.AddHandler(string(constants.FungibleMint), http.MethodPost, func() RequestContext {
		return &MintContext{}
	})
	handler.AddHandler(string(constants.FungibleTransfer), http.MethodPost, func() RequestContext {
		return &TransferContext{}
	})
	handler.AddHandler(string(constants.FungibleConsolidate), http.MethodPost, func() RequestContext {
		return &ConsolidateContext{}
	})
	handler.AddHandler(constants.FungibleSubmit, http.MethodPost, func() RequestContext {
		return &SubmitContext{}
	})
	handler.AddHandler(string(constants.FungibleAccounts), http.MethodGet, func() RequestContext {
		return &AccountContext{}
	})

	return handler
}

func (d *fungibleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	d.router.ServeHTTP(response, request)
}
