package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/GetPricingByNameVirtualMachines_example.json
func ExamplePricingsClient_Get_getPricingsOnSubscriptionVirtualMachinesPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().Get(ctx, "VirtualMachines", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pricing = armsecurity.Pricing{
	// 	Name: to.Ptr("VirtualMachines"),
	// 	Type: to.Ptr("Microsoft.Security/pricings"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/VirtualMachines"),
	// 	Properties: &armsecurity.PricingProperties{
	// 		EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.1921106Z"); return t}()),
	// 		FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 		PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 		SubPlan: to.Ptr("P2"),
	// 		Extensions: []*armsecurity.Extension{
	// 			{
	// 				Name: to.Ptr("AgentlessVmScanning"),
	// 				IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 		}},
	// 	},
	// }
}
