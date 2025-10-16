package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/ResourceGuardCRUD/GetDefaultDeleteProtectedItemRequests.json
func ExampleResourceGuardsClient_GetDefaultDeleteProtectedItemRequestsObject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceGuardsClient().GetDefaultDeleteProtectedItemRequestsObject(ctx, "SampleResourceGroup", "swaggerExample", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.ResourceGuardsClientGetDefaultDeleteProtectedItemRequestsObjectResponse{
	// 	DppBaseResource: &armdataprotection.DppBaseResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.DataProtection/resourceGuards/deleteProtectedItemRequests"),
	// 		ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/resourceGuards/swaggerExample/deleteProtectedItemRequests/default"),
	// 	},
	// }
}
