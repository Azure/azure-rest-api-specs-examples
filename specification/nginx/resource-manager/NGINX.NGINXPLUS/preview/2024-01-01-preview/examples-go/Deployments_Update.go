package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Deployments_Update.json
func ExampleDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginUpdate(ctx, "myResourceGroup", "myDeployment", &armnginx.DeploymentsClientBeginUpdateOptions{Body: nil})
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
	// res.Deployment = armnginx.Deployment{
	// 	Name: to.Ptr("myDeployment"),
	// 	Type: to.Ptr("nginx.nginxplus/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnginx.DeploymentProperties{
	// 		AutoUpgradeProfile: &armnginx.AutoUpgradeProfile{
	// 			UpgradeChannel: to.Ptr("stable"),
	// 		},
	// 		IPAddress: to.Ptr("1.1.1.1"),
	// 		ManagedResourceGroup: to.Ptr("myManagedResourceGroup"),
	// 		NetworkProfile: &armnginx.NetworkProfile{
	// 			FrontEndIPConfiguration: &armnginx.FrontendIPConfiguration{
	// 				PrivateIPAddresses: []*armnginx.PrivateIPAddress{
	// 					{
	// 						PrivateIPAddress: to.Ptr("1.1.1.1"),
	// 						PrivateIPAllocationMethod: to.Ptr(armnginx.NginxPrivateIPAllocationMethodStatic),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 				}},
	// 				PublicIPAddresses: []*armnginx.PublicIPAddress{
	// 					{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/publicIPAddresses/myPublicIPAddress"),
	// 				}},
	// 			},
	// 			NetworkInterfaceConfiguration: &armnginx.NetworkInterfaceConfiguration{
	// 				SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 			},
	// 		},
	// 		NginxVersion: to.Ptr("nginx-1.19.6"),
	// 		ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 		ScalingProperties: &armnginx.DeploymentScalingProperties{
	// 			Capacity: to.Ptr[int32](10),
	// 		},
	// 		UserProfile: &armnginx.DeploymentUserProfile{
	// 			PreferredEmail: to.Ptr("example@example.email"),
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}
