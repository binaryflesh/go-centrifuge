package funding

import (
	"context"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	clientfundingpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/funding"
	"github.com/ethereum/go-ethereum/common/hexutil"
	logging "github.com/ipfs/go-log"
)

var apiLog = logging.Logger("funding-api")

// grpcHandler handles all the funding extension related actions
type grpcHandler struct {
	service Service
	config  config.Service
}

// GRPCHandler returns an implementation of entity.DocumentServiceServer
func GRPCHandler(config config.Service, srv Service) clientfundingpb.FundingServiceServer {
	return &grpcHandler{
		service: srv,
		config:  config,
	}
}

// Create handles a new funding document extension and adds it to an existing document
func (h *grpcHandler) Create(ctx context.Context, req *clientfundingpb.FundingCreatePayload) (*clientfundingpb.FundingResponse, error) {
	apiLog.Debugf("create funding request %v", req)
	ctxHeader, err := contextutil.Context(ctx, h.config)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	identifier, err := hexutil.Decode(req.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, documents.ErrDocumentIdentifier
	}

	// create new funding id
	req.Data.FundingId = newFundingID()

	// returns model with added funding custom fields
	model, err := h.service.DeriveFromPayload(ctxHeader, req, identifier)
	if err != nil {
		return nil, errors.NewTypedError(ErrPayload, err)
	}

	model, jobID, _, err := h.service.Update(ctx, model)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	resp, err := h.service.DeriveFundingResponse(model, req.Data.FundingId)
	if err != nil {
		apiLog.Error(err)
		return nil, errors.NewTypedError(ErrFundingAttr, err)
	}

	resp.Header.JobId = jobID.String()
	return resp, nil
}

// Get returns a funding agreement from an existing document
func (h *grpcHandler) Get(ctx context.Context, req *clientfundingpb.GetRequest) (*clientfundingpb.FundingResponse, error) {
	apiLog.Debugf("Get request %v", req)
	ctxHeader, err := contextutil.Context(ctx, h.config)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	identifier, err := hexutil.Decode(req.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, documents.ErrDocumentIdentifier
	}

	model, err := h.service.GetCurrentVersion(ctxHeader, identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, documents.ErrDocumentNotFound
	}

	resp, err := h.service.DeriveFundingResponse(model, req.FundingId)
	if err != nil {
		apiLog.Error(err)
		return nil, ErrFundingAttr
	}
	return resp, nil
}

// GetList returns all funding agreements of a existing document
func (h *grpcHandler) GetList(ctx context.Context, req *clientfundingpb.GetListRequest) (*clientfundingpb.FundingListResponse, error) {
	apiLog.Debugf("Get request %v", req)
	ctxHeader, err := contextutil.Context(ctx, h.config)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	identifier, err := hexutil.Decode(req.Identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, documents.ErrDocumentIdentifier
	}

	model, err := h.service.GetCurrentVersion(ctxHeader, identifier)
	if err != nil {
		apiLog.Error(err)
		return nil, documents.ErrDocumentNotFound
	}

	resp, err := h.service.DeriveFundingListResponse(model)
	if err != nil {
		apiLog.Error(err)
		return nil, ErrFundingAttr
	}

	return resp, nil
}
