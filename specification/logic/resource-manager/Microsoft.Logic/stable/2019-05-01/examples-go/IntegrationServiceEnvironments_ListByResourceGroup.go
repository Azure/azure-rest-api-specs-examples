package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ListByResourceGroup.json
func ExampleIntegrationServiceEnvironmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationServiceEnvironmentsClient().NewListByResourceGroupPager("testResourceGroup", &armlogic.IntegrationServiceEnvironmentsClientListByResourceGroupOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IntegrationServiceEnvironmentListResult = armlogic.IntegrationServiceEnvironmentListResult{
		// 	Value: []*armlogic.IntegrationServiceEnvironment{
		// 		{
		// 			Name: to.Ptr("ISE-ILB-NU"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 			ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/ISE-ILB-NU"),
		// 			Location: to.Ptr("northeurope"),
		// 			Properties: &armlogic.IntegrationServiceEnvironmentProperties{
		// 				EncryptionConfiguration: &armlogic.IntegrationServiceEnvironmenEncryptionConfiguration{
		// 					EncryptionKeyReference: &armlogic.IntegrationServiceEnvironmenEncryptionKeyReference{
		// 						KeyName: to.Ptr("testKeyName"),
		// 						KeyVault: &armlogic.ResourceReference{
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
		// 						},
		// 						KeyVersion: to.Ptr("13b261d30b984753869902d7f47f4d55"),
		// 					},
		// 				},
		// 				EndpointsConfiguration: &armlogic.FlowEndpointsConfiguration{
		// 					Connector: &armlogic.FlowEndpoints{
		// 						OutgoingIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("40.127.188.117"),
		// 							},
		// 							{
		// 								Address: to.Ptr("40.85.114.29"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.43.2.0/24"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.43.3.0/24"),
		// 						}},
		// 					},
		// 					Workflow: &armlogic.FlowEndpoints{
		// 						AccessEndpointIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("10.43.1.6"),
		// 						}},
		// 						OutgoingIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("40.69.195.162"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.43.1.0/24"),
		// 						}},
		// 					},
		// 				},
		// 				IntegrationServiceEnvironmentID: to.Ptr("13b261d30b984753869902d7f47f4d55"),
		// 				NetworkConfiguration: &armlogic.NetworkConfiguration{
		// 					AccessEndpoint: &armlogic.IntegrationServiceEnvironmentAccessEndpoint{
		// 						Type: to.Ptr(armlogic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
		// 					},
		// 					Subnets: []*armlogic.ResourceReference{
		// 						{
		// 							Name: to.Ptr("VNET-ILB-NorthEurope/s1"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-NorthEurope/s2"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-NorthEurope/s3"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s3"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-NorthEurope/s4"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s4"),
		// 					}},
		// 				},
		// 				ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 				State: to.Ptr(armlogic.WorkflowStateEnabled),
		// 			},
		// 			SKU: &armlogic.IntegrationServiceEnvironmentSKU{
		// 				Name: to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNameDeveloper),
		// 				Capacity: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ISE-ILB-WCentralUS"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 			ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/ISE-ILB-WCentralUS"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armlogic.IntegrationServiceEnvironmentProperties{
		// 				EncryptionConfiguration: &armlogic.IntegrationServiceEnvironmenEncryptionConfiguration{
		// 					EncryptionKeyReference: &armlogic.IntegrationServiceEnvironmenEncryptionKeyReference{
		// 						KeyName: to.Ptr("testKeyName"),
		// 						KeyVault: &armlogic.ResourceReference{
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
		// 						},
		// 						KeyVersion: to.Ptr("13b261d30b984753869902d7f47f4d55"),
		// 					},
		// 				},
		// 				EndpointsConfiguration: &armlogic.FlowEndpointsConfiguration{
		// 					Connector: &armlogic.FlowEndpoints{
		// 						OutgoingIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("13.78.134.201"),
		// 							},
		// 							{
		// 								Address: to.Ptr("13.77.206.166"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.42.2.0/24"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.42.3.0/24"),
		// 						}},
		// 					},
		// 					Workflow: &armlogic.FlowEndpoints{
		// 						AccessEndpointIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("10.42.1.5"),
		// 						}},
		// 						OutgoingIPAddresses: []*armlogic.IPAddress{
		// 							{
		// 								Address: to.Ptr("13.78.237.166"),
		// 							},
		// 							{
		// 								Address: to.Ptr("10.42.1.0/24"),
		// 						}},
		// 					},
		// 				},
		// 				IntegrationServiceEnvironmentID: to.Ptr("08bdba07c6b34ad6a263fc0152ff1735"),
		// 				NetworkConfiguration: &armlogic.NetworkConfiguration{
		// 					AccessEndpoint: &armlogic.IntegrationServiceEnvironmentAccessEndpoint{
		// 						Type: to.Ptr(armlogic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
		// 					},
		// 					Subnets: []*armlogic.ResourceReference{
		// 						{
		// 							Name: to.Ptr("VNET-ILB-WCentralUS/s1"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-WCentralUS/subnets/s1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-WCentralUS/s2"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-WCentralUS/subnets/s2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-WCentralUS/s3"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-WCentralUS/subnets/s3"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VNET-ILB-WCentralUS/s4"),
		// 							Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-WCentralUS/subnets/s4"),
		// 					}},
		// 				},
		// 				ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 				State: to.Ptr(armlogic.WorkflowStateEnabled),
		// 			},
		// 			SKU: &armlogic.IntegrationServiceEnvironmentSKU{
		// 				Name: to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNameDeveloper),
		// 				Capacity: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
