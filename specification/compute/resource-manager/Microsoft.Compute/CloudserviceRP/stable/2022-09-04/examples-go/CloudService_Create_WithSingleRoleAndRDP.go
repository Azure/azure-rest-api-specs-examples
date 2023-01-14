package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudService_Create_WithSingleRoleAndRDP.json
func ExampleCloudServicesClient_BeginCreateOrUpdate_createNewCloudServiceWithSingleRoleAndRdpExtension() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewCloudServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "ConstosoRG", "{cs-name}", armcompute.CloudService{
		Location: to.Ptr("westus"),
		Properties: &armcompute.CloudServiceProperties{
			Configuration: to.Ptr("{ServiceConfiguration}"),
			ExtensionProfile: &armcompute.CloudServiceExtensionProfile{
				Extensions: []*armcompute.Extension{
					{
						Name: to.Ptr("RDPExtension"),
						Properties: &armcompute.CloudServiceExtensionProperties{
							Type:                    to.Ptr("RDP"),
							AutoUpgradeMinorVersion: to.Ptr(false),
							ProtectedSettings:       "<PrivateConfig><Password>{password}</Password></PrivateConfig>",
							Publisher:               to.Ptr("Microsoft.Windows.Azure.Extensions"),
							Settings:                "<PublicConfig><UserName>UserAzure</UserName><Expiration>10/22/2021 15:05:45</Expiration></PublicConfig>",
							TypeHandlerVersion:      to.Ptr("1.2"),
						},
					}},
			},
			NetworkProfile: &armcompute.CloudServiceNetworkProfile{
				LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
					{
						Name: to.Ptr("contosolb"),
						Properties: &armcompute.LoadBalancerConfigurationProperties{
							FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
								{
									Name: to.Ptr("contosofe"),
									Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
										PublicIPAddress: &armcompute.SubResource{
											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
										},
									},
								}},
						},
					}},
			},
			PackageURL: to.Ptr("{PackageUrl}"),
			RoleProfile: &armcompute.CloudServiceRoleProfile{
				Roles: []*armcompute.CloudServiceRoleProfileProperties{
					{
						Name: to.Ptr("ContosoFrontend"),
						SKU: &armcompute.CloudServiceRoleSKU{
							Name:     to.Ptr("Standard_D1_v2"),
							Capacity: to.Ptr[int64](1),
							Tier:     to.Ptr("Standard"),
						},
					}},
			},
			UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
