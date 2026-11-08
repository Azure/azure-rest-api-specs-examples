package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShields_Get.json
func ExampleNeuroShieldsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNeuroShieldsClient().Get(ctx, "myResourceGroup", "myNeuroShield", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworksecurity.NeuroShieldsClientGetResponse{
	// 	NeuroShieldResource: armnetworksecurity.NeuroShieldResource{
	// 		Name: to.Ptr("myNeuroShield"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield"),
	// 		Type: to.Ptr("microsoft.networksecurity/neuroShields"),
	// 		Location: to.Ptr("eastus2"),
	// 		Tags: map[string]*string{
	// 			"env": to.Ptr("test"),
	// 		},
	// 		Properties: &armnetworksecurity.NeuroShieldProperties{
	// 			ProvisioningState: to.Ptr(armnetworksecurity.ProvisioningStateNeuroShieldSucceeded),
	// 			ResourceGUID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			ResponseMode: to.Ptr(armnetworksecurity.NeuroShieldResponseModeManual),
	// 			SeverityThreshold: to.Ptr(armnetworksecurity.NeuroShieldSeverityThresholdMedium),
	// 			Description: to.Ptr("Test NeuroShield resource"),
	// 		},
	// 	},
	// }
}
