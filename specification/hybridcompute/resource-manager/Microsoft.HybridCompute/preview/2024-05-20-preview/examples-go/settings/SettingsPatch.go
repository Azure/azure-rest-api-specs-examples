package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/settings/SettingsPatch.json
func ExampleSettingsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSettingsClient().Patch(ctx, "hybridRG", "Microsoft.HybridCompute", "machines", "testMachine", "default", armhybridcompute.Settings{
		Properties: &armhybridcompute.SettingsProperties{
			GatewayProperties: &armhybridcompute.SettingsGatewayProperties{
				GatewayResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/gateways/newGateway"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Settings = armhybridcompute.Settings{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/settings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridCompute/settings/default"),
	// 	Properties: &armhybridcompute.SettingsProperties{
	// 		GatewayProperties: &armhybridcompute.SettingsGatewayProperties{
	// 			GatewayResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/gateways/newGateway"),
	// 		},
	// 		TenantID: to.Ptr("00000000-1111-2222-5555-444444444444"),
	// 	},
	// }
}
