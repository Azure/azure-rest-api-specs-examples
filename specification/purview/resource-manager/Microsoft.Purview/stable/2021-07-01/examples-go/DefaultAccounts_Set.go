package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Set.json
func ExampleDefaultAccountsClient_Set() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefaultAccountsClient().Set(ctx, armpurview.DefaultAccountPayload{
		AccountName:       to.Ptr("myDefaultAccount"),
		ResourceGroupName: to.Ptr("rg-1"),
		Scope:             to.Ptr("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
		ScopeTenantID:     to.Ptr("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
		ScopeType:         to.Ptr(armpurview.ScopeTypeTenant),
		SubscriptionID:    to.Ptr("12345678-1234-1234-12345678aaa"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefaultAccountPayload = armpurview.DefaultAccountPayload{
	// 	AccountName: to.Ptr("myDefaultAccount"),
	// 	ResourceGroupName: to.Ptr("rg-1"),
	// 	Scope: to.Ptr("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
	// 	ScopeTenantID: to.Ptr("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
	// 	ScopeType: to.Ptr(armpurview.ScopeTypeTenant),
	// 	SubscriptionID: to.Ptr("12345678-1234-1234-12345678aaa"),
	// }
}
