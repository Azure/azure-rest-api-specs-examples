package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceRolesInstance_List.json
func ExampleCloudServiceRoleInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudServiceRoleInstancesClient().NewListPager("ConstosoRG", "{cs-name}", &armcompute.CloudServiceRoleInstancesClientListOptions{Expand: nil})
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
		// page.RoleInstanceListResult = armcompute.RoleInstanceListResult{
		// 	Value: []*armcompute.RoleInstance{
		// 		{
		// 			Name: to.Ptr("ContosoFrontend_IN_0"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_0"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_0/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoFrontend_IN_1"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_1"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_1/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoBackend_IN_0"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_0"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_0/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoBackend_IN_1"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_1"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_1/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
