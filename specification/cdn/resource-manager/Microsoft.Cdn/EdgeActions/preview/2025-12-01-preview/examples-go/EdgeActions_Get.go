package armedgeactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeactions/armedgeactions"
)

// Generated from example definition: 2025-12-01-preview/EdgeActions_Get.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "testrg", "edgeAction1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armedgeactions.ClientGetResponse{
	// 	EdgeAction: &armedgeactions.EdgeAction{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1"),
	// 		Name: to.Ptr("edgeAction1"),
	// 		Type: to.Ptr("Microsoft.Cdn/edgeActions"),
	// 		Location: to.Ptr("global"),
	// 		SKU: &armedgeactions.SKUType{
	// 			Name: to.Ptr("Standard"),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		Properties: &armedgeactions.EdgeActionProperties{
	// 			ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
	// 			Attachments: []*armedgeactions.EdgeActionAttachment{
	// 			},
	// 		},
	// 	},
	// }
}
