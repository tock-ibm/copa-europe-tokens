package tokens

import (
	"github.com/copa-europe-tokens/pkg/constants"
	"github.com/copa-europe-tokens/pkg/types"
)

func (m *Manager) FungibleDeploy(request *types.FungibleDeployRequest) (*types.FungibleDeployResponse, error) {
	typeId := "placeholder"
	return &types.FungibleDeployResponse{
		TypeId:       typeId,
		Name:         request.Name,
		Description:  request.Description,
		ReserveOwner: request.ReserveOwner,
		Supply:       request.Supply,
		Url:          constants.FungibleDescribe.Replace(typeId),
	}, nil
}

func (m *Manager) FungibleDescribe(typeId string) (*types.FungibleDescribeResponse, error) {
	return &types.FungibleDescribeResponse{
		TypeId: typeId,
		Url:    constants.FungibleDescribe.Replace(typeId),
	}, nil
}

func (m *Manager) FungiblePrepareMint(typeId string, request *types.FungibleMintRequest) (*types.FungibleMintResponse, error) {
	return &types.FungibleMintResponse{
		TypeId: typeId,
	}, nil
}

func (m *Manager) FungiblePrepareTransfer(typeId string, request *types.FungibleTransferRequest) (*types.FungibleTransferResponse, error) {
	return &types.FungibleTransferResponse{
		TypeId:   typeId,
		Owner:    request.Owner,
		NewOwner: request.NewOwner,
	}, nil
}

func (m *Manager) FungiblePrepareConsolidate(typeId string, request *types.FungibleConsolidateRequest) (*types.FungibleConsolidateResponse, error) {
	return &types.FungibleConsolidateResponse{
		TypeId: typeId,
		Owner:  request.Owner,
	}, nil
}

func (m *Manager) FungibleSubmitTx(request *types.FungibleSubmitRequest) (*types.FungibleSubmitResponse, error) {
	return &types.FungibleSubmitResponse{}, nil
}

func (m *Manager) FungibleAccounts(typeId string, owner string, account string) (*types.FungibleAccountRecords, error) {
	return &types.FungibleAccountRecords{}, nil
}
