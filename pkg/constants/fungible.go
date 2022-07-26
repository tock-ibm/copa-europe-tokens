// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package constants

import "strings"

type typeURI string

const (
	FungibleEndpoint = "/tokens/fungible/"

	FungibleDeploy = "/tokens/fungible/deploy"
	FungibleSubmit = "/tokens/fungible/submit"

	FungibleDescribe    typeURI = "/tokens/fungible/{typeId}"
	FungibleMint        typeURI = "/tokens/fungible/{typeId}/mint-prepare"
	FungibleTransfer    typeURI = "/tokens/fungible/{typeId}/transfer-prepare"
	FungibleConsolidate typeURI = "/tokens/fungible/{typeId}/consolidate-prepare"
	FungibleAccounts    typeURI = "/tokens/fungible/{typeId}/accounts"
)

func (s typeURI) Replace(typeId string) string {
	return strings.Replace(string(s), "{typeId}", typeId, 1)
}
