package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_AddRootCollectionAdmin.json
func ExampleAccountsClient_AddRootCollectionAdmin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAccountsClient().AddRootCollectionAdmin(ctx, "SampleResourceGroup", "account1", armpurview.CollectionAdminUpdate{
		ObjectID: to.Ptr("7e8de0e7-2bfc-4e1f-9659-2a5785e4356f"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
