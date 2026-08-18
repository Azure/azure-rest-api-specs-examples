package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/FirstPartyServiceTagGet.json
func ExampleFirstPartyServiceTagsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirstPartyServiceTagsClient().Get(ctx, "rg1", "myServiceTag", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.FirstPartyServiceTagsClientGetResponse{
	// 	FirstPartyServiceTag: armnetwork.FirstPartyServiceTag{
	// 		Name: to.Ptr("myServiceTag"),
	// 		Type: to.Ptr("Microsoft.Network/firstPartyServiceTags"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firstPartyServiceTags/myServiceTag"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 		Properties: &armnetwork.FirstPartyServiceTagPropertiesFormat{
	// 			Value: to.Ptr("myServiceTagValue"),
	// 			FailedReason: to.Ptr(""),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	},
	// }
}
