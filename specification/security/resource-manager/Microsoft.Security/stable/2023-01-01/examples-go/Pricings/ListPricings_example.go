package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/ListPricings_example.json
func ExamplePricingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPricingsClient().List(ctx, nil)
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
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/VirtualMachines"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				SubPlan: to.Ptr("P2"),
	// 				Extensions: []*armsecurity.Extension{
	// 					{
	// 						Name: to.Ptr("AgentlessVmScanning"),
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				}},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("SqlServers"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/SqlServers"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("AppServices"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/AppServices"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("StorageAccounts"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/StorageAccounts"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 				SubPlan: to.Ptr("PerStorageAccount"),
	// 				Extensions: []*armsecurity.Extension{
	// 					{
	// 						Name: to.Ptr("OnUploadMalwareScanning"),
	// 						AdditionalExtensionProperties: map[string]any{
	// 							"capGBPerMonthPerStorageAccount": float64(10),
	// 						},
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 					},
	// 					{
	// 						Name: to.Ptr("SensitiveDataDiscovery"),
	// 						IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 				}},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("SqlServerVirtualMachines"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/SqlServerVirtualMachines"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("KubernetesService"),
	// 			Type: to.Ptr("Microsoft.Security/pricings"),
	// 			ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/KubernetesService"),
	// 			Properties: &armsecurity.PricingProperties{
	// 				Deprecated: to.Ptr(true),
	// 				FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 				PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 				ReplacedBy: []*string{
	// 					to.Ptr("Containers")},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("ContainerRegistry"),
	// 				Type: to.Ptr("Microsoft.Security/pricings"),
	// 				ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/ContainerRegistry"),
	// 				Properties: &armsecurity.PricingProperties{
	// 					Deprecated: to.Ptr(true),
	// 					FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 					PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 					ReplacedBy: []*string{
	// 						to.Ptr("Containers")},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("KeyVaults"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/KeyVaults"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Dns"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Dns"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Arm"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Arm"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("OpenSourceRelationalDatabases"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/OpenSourceRelationalDatabases"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierFree),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Containers"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/Containers"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 						Extensions: []*armsecurity.Extension{
	// 							{
	// 								Name: to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
	// 								IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 						}},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("CloudPosture"),
	// 					Type: to.Ptr("Microsoft.Security/pricings"),
	// 					ID: to.Ptr("/subscriptions/d34fd44c-ebfa-4a9c-bceb-9eeafe72ac15/providers/Microsoft.Security/pricings/CloudPosture"),
	// 					Properties: &armsecurity.PricingProperties{
	// 						EnablementTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T12:42:42.192Z"); return t}()),
	// 						FreeTrialRemainingTime: to.Ptr("PT0S"),
	// 						PricingTier: to.Ptr(armsecurity.PricingTierStandard),
	// 						Extensions: []*armsecurity.Extension{
	// 							{
	// 								Name: to.Ptr("AgentlessVmScanning"),
	// 								IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 							},
	// 							{
	// 								Name: to.Ptr("AgentlessDiscoveryForKubernetes"),
	// 								IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 							},
	// 							{
	// 								Name: to.Ptr("SensitiveDataDiscovery"),
	// 								IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 							},
	// 							{
	// 								Name: to.Ptr("ContainerRegistriesVulnerabilityAssessments"),
	// 								IsEnabled: to.Ptr(armsecurity.IsEnabledTrue),
	// 						}},
	// 					},
	// 			}},
	// 		}
}
