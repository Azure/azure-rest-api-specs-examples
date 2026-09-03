package armredhatopenshifthcp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
)

// Generated from example definition: 2026-09-01-preview/HcpOpenShiftClusters_ListByResourceGroup_MaximumSet_Gen.json
func ExampleHcpOpenShiftClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshifthcp.NewClientFactory("FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHcpOpenShiftClustersClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armredhatopenshifthcp.HcpOpenShiftClustersClientListByResourceGroupResponse{
		// 	HcpOpenShiftClusterListResult: armredhatopenshifthcp.HcpOpenShiftClusterListResult{
		// 		Value: []*armredhatopenshifthcp.HcpOpenShiftCluster{
		// 			{
		// 				Properties: &armredhatopenshifthcp.HcpOpenShiftClusterProperties{
		// 					ProvisioningState: to.Ptr(armredhatopenshifthcp.ProvisioningStateSucceeded),
		// 					Version: &armredhatopenshifthcp.VersionProfile{
		// 						ChannelGroup: to.Ptr("stable"),
		// 						ID: to.Ptr("4.12"),
		// 					},
		// 					DNS: &armredhatopenshifthcp.DNSProfile{
		// 						BaseDomain: to.Ptr("yubrqcgqdhgqfkobjqm"),
		// 						BaseDomainPrefix: to.Ptr("jcldjrtyebhrlxs"),
		// 					},
		// 					Network: &armredhatopenshifthcp.NetworkProfile{
		// 						NetworkType: to.Ptr(armredhatopenshifthcp.NetworkTypeOVNKubernetes),
		// 						PodCIDR: to.Ptr("10.128.0.0/14"),
		// 						ServiceCIDR: to.Ptr("172.30.0.0/16"),
		// 						MachineCIDR: to.Ptr("10.0.0.0/16"),
		// 						HostPrefix: to.Ptr[int32](26),
		// 					},
		// 					Console: &armredhatopenshifthcp.ConsoleProfile{
		// 						URL: to.Ptr("https://console.test.shrd.usw3test.hcp.osadev.cloud"),
		// 					},
		// 					API: &armredhatopenshifthcp.APIProfile{
		// 						URL: to.Ptr("https://api.test.shrd.usw3test.hcp.osadev.cloud:443"),
		// 						Visibility: to.Ptr(armredhatopenshifthcp.VisibilityPublic),
		// 					},
		// 					Ingress: &armredhatopenshifthcp.IngressProfile{
		// 						Type: to.Ptr(armredhatopenshifthcp.IngressTypePublic),
		// 					},
		// 					Platform: &armredhatopenshifthcp.PlatformProfile{
		// 						ManagedResourceGroup: to.Ptr("nhyhywrxupo"),
		// 						SubnetID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/hcp-network-example/subnets/example-subnet"),
		// 						VnetIntegrationSubnetID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/hcp-network-example/subnets/vnet-integration-subnet"),
		// 						OutboundType: to.Ptr(armredhatopenshifthcp.OutboundTypeLoadBalancer),
		// 						NetworkSecurityGroupID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.Network/networkSecurityGroups/nsg-example"),
		// 						OperatorsAuthentication: &armredhatopenshifthcp.OperatorsAuthenticationProfile{
		// 							UserAssignedIdentities: &armredhatopenshifthcp.UserAssignedIdentitiesProfile{
		// 								ControlPlaneOperators: map[string]*string{
		// 								},
		// 								DataPlaneOperators: map[string]*string{
		// 								},
		// 								ServiceManagedIdentity: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/serviceMI"),
		// 							},
		// 						},
		// 						IssuerURL: to.Ptr("https://oidc.contoso.com"),
		// 					},
		// 					Autoscaling: &armredhatopenshifthcp.ClusterAutoscalingProfile{
		// 						MaxNodesTotal: to.Ptr[int32](0),
		// 						MaxPodGracePeriodSeconds: to.Ptr[int32](0),
		// 						MaxNodeProvisionTimeSeconds: to.Ptr[int32](0),
		// 						PodPriorityThreshold: to.Ptr[int32](1),
		// 					},
		// 					Etcd: &armredhatopenshifthcp.EtcdProfile{
		// 						DataEncryption: &armredhatopenshifthcp.EtcdDataEncryptionProfile{
		// 							KeyManagementMode: to.Ptr(armredhatopenshifthcp.EtcdDataEncryptionKeyManagementModeTypeCustomerManaged),
		// 							CustomerManaged: &armredhatopenshifthcp.CustomerManagedEncryptionProfile{
		// 								EncryptionType: to.Ptr(armredhatopenshifthcp.CustomerManagedEncryptionTypeKms),
		// 								Kms: &armredhatopenshifthcp.KmsEncryptionProfile{
		// 									VaultName: to.Ptr("my-cool-vault"),
		// 									Visibility: to.Ptr(armredhatopenshifthcp.KeyVaultVisibilityPublic),
		// 									ActiveKey: &armredhatopenshifthcp.KmsKey{
		// 										Name: to.Ptr("my-cool-key"),
		// 										Version: to.Ptr("8e73e7d1fd7d4a87b730f676fc77d3a6"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ImageDigestMirrors: []*armredhatopenshifthcp.ImageDigestMirror{
		// 						{
		// 							Source: to.Ptr("registry.example.com/image1"),
		// 							Mirrors: []*string{
		// 								to.Ptr("mirror1.example.com/image1"),
		// 								to.Ptr("mirror2.example.com/image1"),
		// 							},
		// 						},
		// 						{
		// 							Source: to.Ptr("registry.example.com/image2"),
		// 							Mirrors: []*string{
		// 								to.Ptr("mirror1.example.com/image2"),
		// 							},
		// 						},
		// 					},
		// 					NodeDrainTimeoutMinutes: to.Ptr[int32](20),
		// 					ClusterImageRegistry: &armredhatopenshifthcp.ClusterImageRegistryProfile{
		// 						State: to.Ptr(armredhatopenshifthcp.ClusterImageRegistryStateEnabled),
		// 					},
		// 					CryptoRestrictions: to.Ptr(armredhatopenshifthcp.CryptoRestrictionsNone),
		// 				},
		// 				Identity: &armredhatopenshifthcp.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("e5867472-f0ed-4fc1-80e1-59b4c0256adb"),
		// 					TenantID: to.Ptr("2a58de3b-8a38-44cd-8f33-4bd5b91479c1"),
		// 					Type: to.Ptr(armredhatopenshifthcp.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armredhatopenshifthcp.UserAssignedIdentity{
		// 						"/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/serviceMI": &armredhatopenshifthcp.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("15a1e2d8-41ae-4068-8ea9-a80f2cdd94c3"),
		// 							ClientID: to.Ptr("a60f3156-367f-4303-bb21-7a41b3c41cb9"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4181": to.Ptr("leaswtidajsjtgmqawhdl"),
		// 				},
		// 				Location: to.Ptr("ayecbdqonsqfowbq"),
		// 				ID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.RedHatOpenShift/resourceType/resourceName"),
		// 				Name: to.Ptr("vuwzuwooutjavgdhoatz"),
		// 				Type: to.Ptr("utiyj"),
		// 				SystemData: &armredhatopenshifthcp.SystemData{
		// 					CreatedBy: to.Ptr("lsrkqcuijqfp"),
		// 					CreatedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.March, 27, 14, 57, 32, 578000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("tgpmwu"),
		// 					LastModifiedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.March, 27, 14, 57, 32, 578000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
