package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2023-04-01/examples/Deployments_List.json
func ExampleDeploymentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeploymentsClient().NewListPager(nil)
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
		// page.DeploymentListResponse = armnginx.DeploymentListResponse{
		// 	Value: []*armnginx.Deployment{
		// 		{
		// 			Name: to.Ptr("myDeployment"),
		// 			Type: to.Ptr("nginx.nginxplus/deployments"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnginx.DeploymentProperties{
		// 				IPAddress: to.Ptr("1.1.1.1"),
		// 				ManagedResourceGroup: to.Ptr("myManagedResourceGroup"),
		// 				NetworkProfile: &armnginx.NetworkProfile{
		// 					FrontEndIPConfiguration: &armnginx.FrontendIPConfiguration{
		// 						PrivateIPAddresses: []*armnginx.PrivateIPAddress{
		// 							{
		// 								PrivateIPAddress: to.Ptr("1.1.1.1"),
		// 								PrivateIPAllocationMethod: to.Ptr(armnginx.NginxPrivateIPAllocationMethodStatic),
		// 								SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 						}},
		// 						PublicIPAddresses: []*armnginx.PublicIPAddress{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/publicIPAddresses/myPublicIPAddress"),
		// 						}},
		// 					},
		// 					NetworkInterfaceConfiguration: &armnginx.NetworkInterfaceConfiguration{
		// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 					},
		// 				},
		// 				NginxVersion: to.Ptr("nginx-1.19.6"),
		// 				ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
