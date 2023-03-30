package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Patch.json
func ExampleIntegrationServiceEnvironmentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationServiceEnvironmentsClient().BeginUpdate(ctx, "testResourceGroup", "testIntegrationServiceEnvironment", armlogic.IntegrationServiceEnvironment{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
		SKU: &armlogic.IntegrationServiceEnvironmentSKU{
			Name:     to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNameDeveloper),
			Capacity: to.Ptr[int32](0),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationServiceEnvironment = armlogic.IntegrationServiceEnvironment{
	// 	Name: to.Ptr("testIntegrationServiceEnvironment"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
	// 	ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armlogic.IntegrationServiceEnvironmentProperties{
	// 		EncryptionConfiguration: &armlogic.IntegrationServiceEnvironmenEncryptionConfiguration{
	// 			EncryptionKeyReference: &armlogic.IntegrationServiceEnvironmenEncryptionKeyReference{
	// 				KeyName: to.Ptr("testKeyName"),
	// 				KeyVault: &armlogic.ResourceReference{
	// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
	// 				},
	// 				KeyVersion: to.Ptr("13b261d30b984753869902d7f47f4d55"),
	// 			},
	// 		},
	// 		EndpointsConfiguration: &armlogic.FlowEndpointsConfiguration{
	// 			Connector: &armlogic.FlowEndpoints{
	// 				OutgoingIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("40.127.188.117"),
	// 					},
	// 					{
	// 						Address: to.Ptr("40.85.114.29"),
	// 					},
	// 					{
	// 						Address: to.Ptr("10.43.2.0/24"),
	// 					},
	// 					{
	// 						Address: to.Ptr("10.43.3.0/24"),
	// 				}},
	// 			},
	// 			Workflow: &armlogic.FlowEndpoints{
	// 				AccessEndpointIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("10.43.1.6"),
	// 				}},
	// 				OutgoingIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("40.69.195.162"),
	// 					},
	// 					{
	// 						Address: to.Ptr("10.43.1.0/24"),
	// 				}},
	// 			},
	// 		},
	// 		IntegrationServiceEnvironmentID: to.Ptr("13b261d30b984753869902d7f47f4d55"),
	// 		NetworkConfiguration: &armlogic.NetworkConfiguration{
	// 			AccessEndpoint: &armlogic.IntegrationServiceEnvironmentAccessEndpoint{
	// 				Type: to.Ptr(armlogic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
	// 			},
	// 			Subnets: []*armlogic.ResourceReference{
	// 				{
	// 					Name: to.Ptr("VNET-ILB-NorthEurope/s1"),
	// 					Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("VNET-ILB-NorthEurope/s2"),
	// 					Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s2"),
	// 				},
	// 				{
	// 					Name: to.Ptr("VNET-ILB-NorthEurope/s3"),
	// 					Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s3"),
	// 				},
	// 				{
	// 					Name: to.Ptr("VNET-ILB-NorthEurope/s4"),
	// 					Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/VNET-ILB-NorthEurope/subnets/s4"),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
	// 		State: to.Ptr(armlogic.WorkflowStateEnabled),
	// 	},
	// 	SKU: &armlogic.IntegrationServiceEnvironmentSKU{
	// 		Name: to.Ptr(armlogic.IntegrationServiceEnvironmentSKUNameDeveloper),
	// 		Capacity: to.Ptr[int32](0),
	// 	},
	// }
}
