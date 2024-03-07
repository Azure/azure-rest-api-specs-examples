package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/ListResourcePricings_example.json
func ExamplePricingsClient_List_getPricingsOnResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().List(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1", &armsecurity.PricingsClientListOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PricingList = armsecurity.PricingList{
	// 	Value: []*armsecurity.Pricing{
	// 		{
	// 			Name: to.Ptr("VirtualMachines"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1/providers/Microsoft.Security/pricings/VirtualMachines"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				Inherited: to.Ptr(armsecurity.InheritedTrue),
	// 				InheritedFrom: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				SubPlan: to.Ptr("P2"),
	// 				Extensions: []*armsecurity.Extension{
	// 					{
	// 						Name: to.Ptr("AgentlessVmScanning"),
	// 						AdditionalExtensionProperties: map[string]any{
	// 							"ExclusionTags": "[{\"Key\":\"TestKey1\",\"Value\":\"TestValue1\"},{\"Key\":\"TestKey2\",\"Value\":\"TestValue2\"}]",
	// 						},
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 					},
	// 					{
	// 						Name: to.Ptr("MdeDesignatedSubscription"),
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				}},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("CloudPosture"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1/providers/Microsoft.Security/pricings/CloudPosture"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				Inherited: to.Ptr(armsecurity.InheritedFalse),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 			},
	// 	}},
	// }
}
