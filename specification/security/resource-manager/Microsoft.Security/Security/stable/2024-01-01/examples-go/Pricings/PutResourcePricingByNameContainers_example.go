package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2024-01-01/Pricings/PutResourcePricingByNameContainers_example.json
func ExamplePricingsClient_Update_updatePricingOnResourceExampleForContainersPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().Update(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/demo-containers-rg/providers/Microsoft.ContainerService/managedClusters/demo-aks-cluster", "Containers", armsecurity.Pricing{
		Properties: &armsecurity.PricingProperties{
			PricingTier: to.Ptr(armsecurity.PricingTierStandard),
			Extensions: []*armsecurity.Extension{
				{
					Name:      to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
				},
				{
					Name:      to.Ptr("ContainerSensor"),
					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
				},
				{
					Name:      to.Ptr("AgentlessDiscoveryForKubernetes"),
					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
				},
				{
					Name: to.Ptr("AgentlessVmScanning"),
					AdditionalExtensionProperties: map[string]any{
						"ExclusionTags": "[]",
					},
					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
				},
				{
					Name:      to.Ptr("ContainerIntegrityContribution"),
					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.PricingsClientUpdateResponse{
	// 	Pricing: armsecurity.Pricing{
	// 		Name: to.Ptr("Containers"),
	// 		Type: to.Ptr("Microsoft.Security/pricings"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/demo-containers-rg/providers/Microsoft.ContainerService/managedClusters/demo-aks-cluster/providers/Microsoft.Security/pricings/Containers"),
	// 		Properties: &armsecurity.PricingProperties{
	// 			EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-14T16:35:12.4567890Z"); return t}()),
	// 			FreeTrialRemainingTime: to.Ptr("P29DT23H55M"),
	// 			Inherited: to.Ptr(armsecurity.InheritedFalse),
	// 			PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 			Extensions: []*armsecurity.Extension{
	// 				{
	// 					Name: to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
	// 					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				},
	// 				{
	// 					Name: to.Ptr("ContainerSensor"),
	// 					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				},
	// 				{
	// 					Name: to.Ptr("AgentlessDiscoveryForKubernetes"),
	// 					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				},
	// 				{
	// 					Name: to.Ptr("AgentlessVmScanning"),
	// 					AdditionalExtensionProperties: map[string]any{
	// 						"ExclusionTags": "[]",
	// 					},
	// 					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				},
	// 				{
	// 					Name: to.Ptr("ContainerIntegrityContribution"),
	// 					IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
