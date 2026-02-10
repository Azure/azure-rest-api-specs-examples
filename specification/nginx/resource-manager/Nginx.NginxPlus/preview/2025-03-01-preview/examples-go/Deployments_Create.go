package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/Deployments_Create.json
func ExampleDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDeployment", armnginx.Deployment{}, nil)
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
	// res = armnginx.DeploymentsClientCreateOrUpdateResponse{
	// 	Deployment: &armnginx.Deployment{
	// 		Name: to.Ptr("myDeployment"),
	// 		Type: to.Ptr("nginx.nginxplus/deployments"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnginx.DeploymentProperties{
	// 			AutoUpgradeProfile: &armnginx.AutoUpgradeProfile{
	// 				UpgradeChannel: to.Ptr("stable"),
	// 			},
	// 			DataplaneAPIEndpoint: to.Ptr("mynginx-75b3bf22a555.eastus2.nginxaas.net"),
	// 			IPAddress: to.Ptr("1.1.1.1"),
	// 			NetworkProfile: &armnginx.NetworkProfile{
	// 				FrontEndIPConfiguration: &armnginx.FrontendIPConfiguration{
	// 					PrivateIPAddresses: []*armnginx.PrivateIPAddress{
	// 						{
	// 							PrivateIPAddress: to.Ptr("1.1.1.1"),
	// 							PrivateIPAllocationMethod: to.Ptr(armnginx.NginxPrivateIPAllocationMethodStatic),
	// 							SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 						},
	// 					},
	// 					PublicIPAddresses: []*armnginx.PublicIPAddress{
	// 						{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/publicIPAddresses/myPublicIPAddress"),
	// 						},
	// 					},
	// 				},
	// 				NetworkInterfaceConfiguration: &armnginx.NetworkInterfaceConfiguration{
	// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 				},
	// 			},
	// 			NginxAppProtect: &armnginx.DeploymentPropertiesNginxAppProtect{
	// 				WebApplicationFirewallSettings: &armnginx.WebApplicationFirewallSettings{
	// 					ActivationState: to.Ptr(armnginx.ActivationStateEnabled),
	// 				},
	// 				WebApplicationFirewallStatus: &armnginx.WebApplicationFirewallStatus{
	// 					AttackSignaturesPackage: &armnginx.WebApplicationFirewallPackage{
	// 						RevisionDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-21T15:50:53Z"); return t}()),
	// 						Version: to.Ptr("2024.02.21"),
	// 					},
	// 					BotSignaturesPackage: &armnginx.WebApplicationFirewallPackage{
	// 						RevisionDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-22T15:50:53Z"); return t}()),
	// 						Version: to.Ptr("2024.02.22"),
	// 					},
	// 					ComponentVersions: &armnginx.WebApplicationFirewallComponentVersions{
	// 						WafEngineVersion: to.Ptr("10.624.0"),
	// 						WafNginxVersion: to.Ptr("4.815.0"),
	// 					},
	// 					ThreatCampaignsPackage: &armnginx.WebApplicationFirewallPackage{
	// 						RevisionDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-23T15:50:53Z"); return t}()),
	// 						Version: to.Ptr("2024.02.23"),
	// 					},
	// 				},
	// 			},
	// 			NginxVersion: to.Ptr("nginx-1.19.6"),
	// 			ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 			ScalingProperties: &armnginx.DeploymentScalingProperties{
	// 				Capacity: to.Ptr[int32](10),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"Environment": to.Ptr("Dev"),
	// 		},
	// 	},
	// }
}
