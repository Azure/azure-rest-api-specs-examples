package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: 2022-04-01/GetProviderOperationsRP.json
func ExampleProviderOperationsMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderOperationsMetadataClient().Get(ctx, "resourceProviderNamespace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armauthorization.ProviderOperationsMetadataClientGetResponse{
	// 	ProviderOperationsMetadata: &armauthorization.ProviderOperationsMetadata{
	// 		Name: to.Ptr("name"),
	// 		Type: to.Ptr("type"),
	// 		DisplayName: to.Ptr("displayName"),
	// 		ID: to.Ptr("id"),
	// 		ResourceTypes: []*armauthorization.ResourceType{
	// 			{
	// 				Name: to.Ptr("name"),
	// 				DisplayName: to.Ptr("name"),
	// 				Operations: []*armauthorization.ProviderOperation{
	// 				},
	// 			},
	// 		},
	// 		Operations: []*armauthorization.ProviderOperation{
	// 		},
	// 	},
	// }
}
