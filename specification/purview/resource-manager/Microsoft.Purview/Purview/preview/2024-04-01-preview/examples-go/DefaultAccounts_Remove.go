package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview/v2"
)

// Generated from example definition: 2024-04-01-preview/DefaultAccounts_Remove.json
func ExampleDefaultAccountsClient_Remove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefaultAccountsClient().Remove(ctx, "ee85a74c-405e-4adc-bb47-ffa8ca0c9f31", armpurview.ScopeTypeTenant, &armpurview.DefaultAccountsClientRemoveOptions{
		Scope: to.Ptr("12345678-1234-1234-1234-12345678abcd")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurview.DefaultAccountsClientRemoveResponse{
	// }
}
