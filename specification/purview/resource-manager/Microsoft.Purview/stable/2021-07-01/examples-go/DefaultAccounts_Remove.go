package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Remove.json
func ExampleDefaultAccountsClient_Remove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpurview.NewDefaultAccountsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Remove(ctx,
		"11733A4E-BA84-46FF-91D1-AFF1A3215A90",
		armpurview.ScopeTypeTenant,
		&armpurview.DefaultAccountsClientRemoveOptions{Scope: to.Ptr("11733A4E-BA84-46FF-91D1-AFF1A3215A90")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
