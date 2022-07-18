// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package types

type FungibleDeployRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ReserveOwner string `json:"reserveOwner"`
	Supply       uint64 `json:"supply"`
}

type FungibleDeployResponse struct {
	TypeId       string `json:"typeId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Supply       uint64 `json:"supply"`
	ReserveOwner string `json:"reserveOwner"`
	Url          string `json:"url"`
}

type FungibleDescribeRequest struct {
}

type FungibleDescribeResponse FungibleDeployResponse

type FungibleMintRequest struct {
	Supply uint64 `json:"supply"`
}

type FungibleMintResponse struct {
	TypeId        string `json:"typeId"`
	TxEnvelope    string `json:"txEnvelope"`    //base64 (std, padded) encoding of bytes
	TxPayloadHash string `json:"txPayloadHash"` //base64 (std, padded) encoding of bytes
}

type FungibleTransferRequest struct {
	Owner    string `json:"owner"`
	NewOwner string `json:"newOwner"`
	Comment  string `json:"comment"`
}

type FungibleTransferResponse struct {
	TypeId        string `json:"typeId"`
	Owner         string `json:"owner"`
	NewOwner      string `json:"newOwner"`
	TxEnvelope    string `json:"txEnvelope"`    //base64 (std, padded) encoding of bytes
	TxPayloadHash string `json:"txPayloadHash"` //base64 (std, padded) encoding of bytes
}

type FungibleConsolidateRequest struct {
	Owner    string   `json:"owner"`
	Accounts []string `json:"accounts"`
}

type FungibleConsolidateResponse struct {
	TypeId        string `json:"typeId"`
	Owner         string `json:"owner"`
	TxEnvelope    string `json:"txEnvelope"`    //base64 (std, padded) encoding of bytes
	TxPayloadHash string `json:"txPayloadHash"` //base64 (std, padded) encoding of bytes
}

type FungibleSubmitRequest struct {
	TypeId        string `json:"typeId"`
	TxEnvelope    string `json:"txEnvelope"`    //base64 (std, padded) encoding of bytes
	TxPayloadHash string `json:"txPayloadHash"` //base64 (std, padded) encoding of bytes
	Signer        string `json:"signer"`
	Signature     string `json:"signature"` //base64 (std, padded) encoding of bytes
}

type FungibleSubmitResponse struct {
	Account string `json:"account"`
	Url     string `json:"url"`
}

type FungibleAccountRequest struct {
}

type FungibleAccountRecord struct {
	Account string `json:"account"`
	Owner   string `json:"owner"`
	Balance uint64 `json:"balance"`
	Comment string `json:"comment"`
}

type FungibleAccountRecords = []FungibleAccountRecord
