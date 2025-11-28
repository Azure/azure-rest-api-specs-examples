package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/NodeTypePutOperation_example_max.json
func ExampleNodeTypesClient_BeginCreateOrUpdate_putANodeTypeWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNodeTypesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "BE-testResourceGroup-testRegion-test", armservicefabricmanagedclusters.NodeType{
		Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
			AdditionalDataDisks: []*armservicefabricmanagedclusters.VmssDataDisk{
				{
					DiskLetter: to.Ptr("F"),
					DiskSizeGB: to.Ptr[int32](256),
					DiskType:   to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
					Lun:        to.Ptr[int32](1),
				},
				{
					DiskLetter: to.Ptr("G"),
					DiskSizeGB: to.Ptr[int32](150),
					DiskType:   to.Ptr(armservicefabricmanagedclusters.DiskTypePremiumLRS),
					Lun:        to.Ptr[int32](2),
				},
			},
			AdditionalNetworkInterfaceConfigurations: []*armservicefabricmanagedclusters.AdditionalNetworkInterfaceConfiguration{
				{
					Name: to.Ptr("nic-1"),
					DscpConfiguration: &armservicefabricmanagedclusters.SubResource{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
					},
					EnableAcceleratedNetworking: to.Ptr(true),
					IPConfigurations: []*armservicefabricmanagedclusters.IPConfiguration{
						{
							Name: to.Ptr("ipconfig-1"),
							ApplicationGatewayBackendAddressPools: []*armservicefabricmanagedclusters.SubResource{
								{
									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
								},
							},
							LoadBalancerBackendAddressPools: []*armservicefabricmanagedclusters.SubResource{
								{
									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
								},
							},
							LoadBalancerInboundNatPools: []*armservicefabricmanagedclusters.SubResource{
								{
									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
								},
							},
							PrivateIPAddressVersion: to.Ptr(armservicefabricmanagedclusters.PrivateIPAddressVersionIPv4),
							PublicIPAddressConfiguration: &armservicefabricmanagedclusters.IPConfigurationPublicIPAddressConfiguration{
								Name: to.Ptr("publicip-1"),
								IPTags: []*armservicefabricmanagedclusters.IPTag{
									{
										IPTagType: to.Ptr("RoutingPreference"),
										Tag:       to.Ptr("Internet"),
									},
								},
								PublicIPAddressVersion: to.Ptr(armservicefabricmanagedclusters.PublicIPAddressVersionIPv4),
							},
							Subnet: &armservicefabricmanagedclusters.SubResource{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
							},
						},
					},
				},
			},
			Capacities: map[string]*string{
				"ClientConnections": to.Ptr("65536"),
			},
			ComputerNamePrefix:          to.Ptr("BE"),
			DataDiskLetter:              to.Ptr("S"),
			DataDiskSizeGB:              to.Ptr[int32](200),
			DataDiskType:                to.Ptr(armservicefabricmanagedclusters.DiskTypePremiumLRS),
			DscpConfigurationID:         to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
			EnableAcceleratedNetworking: to.Ptr(true),
			EnableEncryptionAtHost:      to.Ptr(true),
			EnableNodePublicIP:          to.Ptr(true),
			EnableNodePublicIPv6:        to.Ptr(true),
			EnableOverProvisioning:      to.Ptr(false),
			EvictionPolicy:              to.Ptr(armservicefabricmanagedclusters.EvictionPolicyTypeDeallocate),
			FrontendConfigurations: []*armservicefabricmanagedclusters.FrontendConfiguration{
				{
					ApplicationGatewayBackendAddressPoolID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
					LoadBalancerBackendAddressPoolID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
					LoadBalancerInboundNatPoolID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
				},
			},
			IsPrimary:               to.Ptr(false),
			IsSpotVM:                to.Ptr(true),
			IsStateless:             to.Ptr(true),
			MultiplePlacementGroups: to.Ptr(true),
			NatGatewayID:            to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/natGateways/myNatGateway"),
			PlacementProperties: map[string]*string{
				"HasSSD":       to.Ptr("true"),
				"NodeColor":    to.Ptr("green"),
				"SomeProperty": to.Ptr("5"),
			},
			SecureBootEnabled:            to.Ptr(true),
			SecurityType:                 to.Ptr(armservicefabricmanagedclusters.SecurityTypeConfidentialVM),
			SecurityEncryptionType:       to.Ptr(armservicefabricmanagedclusters.SecurityEncryptionTypeDiskWithVMGuestState),
			ServiceArtifactReferenceID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/myVmArtifactProfile"),
			SpotRestoreTimeout:           to.Ptr("PT30M"),
			SubnetID:                     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			UseDefaultPublicLoadBalancer: to.Ptr(true),
			UseEphemeralOSDisk:           to.Ptr(true),
			VMApplications: []*armservicefabricmanagedclusters.VMApplication{
				{
					ConfigurationReference:          to.Ptr("https://mystorageaccount.blob.core.windows.net/containername/blobname"),
					EnableAutomaticUpgrade:          to.Ptr(true),
					Order:                           to.Ptr[int32](1),
					PackageReferenceID:              to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/applications/myApplication/versions/1.0.0"),
					TreatFailureAsDeploymentFailure: to.Ptr(false),
					VMGalleryTags:                   to.Ptr("{\"Tag1\":\"Value1\",\"Tag2\":\"Value2\"}"),
				},
			},
			VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
				{
					Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
					Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
						Type:                    to.Ptr("GenevaMonitoring"),
						AutoUpgradeMinorVersion: to.Ptr(true),
						EnableAutomaticUpgrade:  to.Ptr(true),
						ForceUpdateTag:          to.Ptr("v.1.0"),
						Publisher:               to.Ptr("Microsoft.Azure.Geneva"),
						Settings:                map[string]any{},
						SetupOrder: []*armservicefabricmanagedclusters.VmssExtensionSetupOrder{
							to.Ptr(armservicefabricmanagedclusters.VmssExtensionSetupOrderBeforeSFRuntime),
						},
						TypeHandlerVersion: to.Ptr("2.0"),
					},
				},
			},
			VMImageOffer:     to.Ptr("WindowsServer"),
			VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
			VMImageSKU:       to.Ptr("2016-Datacenter-Server-Core"),
			VMImageVersion:   to.Ptr("latest"),
			VMInstanceCount:  to.Ptr[int32](10),
			VMManagedIdentity: &armservicefabricmanagedclusters.VMManagedIdentity{
				UserAssignedIdentities: []*string{
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2"),
				},
			},
			VMSecrets: []*armservicefabricmanagedclusters.VaultSecretGroup{
				{
					SourceVault: &armservicefabricmanagedclusters.SubResource{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
					},
					VaultCertificates: []*armservicefabricmanagedclusters.VaultCertificate{
						{
							CertificateStore: to.Ptr("My"),
							CertificateURL:   to.Ptr("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),
						},
					},
				},
			},
			VMSetupActions: []*armservicefabricmanagedclusters.VMSetupAction{
				to.Ptr(armservicefabricmanagedclusters.VMSetupActionEnableContainers),
				to.Ptr(armservicefabricmanagedclusters.VMSetupActionEnableHyperV),
			},
			VMSize:         to.Ptr("Standard_DS3"),
			IsOutboundOnly: to.Ptr(true),
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
	// res = armservicefabricmanagedclusters.NodeTypesClientCreateOrUpdateResponse{
	// 	NodeType: &armservicefabricmanagedclusters.NodeType{
	// 		Name: to.Ptr("BE"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/nodeTypes"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster/nodeTypes/BE"),
	// 		Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
	// 			AdditionalDataDisks: []*armservicefabricmanagedclusters.VmssDataDisk{
	// 				{
	// 					DiskLetter: to.Ptr("F"),
	// 					DiskSizeGB: to.Ptr[int32](256),
	// 					DiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
	// 					Lun: to.Ptr[int32](1),
	// 				},
	// 				{
	// 					DiskLetter: to.Ptr("G"),
	// 					DiskSizeGB: to.Ptr[int32](150),
	// 					DiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypePremiumLRS),
	// 					Lun: to.Ptr[int32](2),
	// 				},
	// 			},
	// 			AdditionalNetworkInterfaceConfigurations: []*armservicefabricmanagedclusters.AdditionalNetworkInterfaceConfiguration{
	// 				{
	// 					Name: to.Ptr("nic-1"),
	// 					DscpConfiguration: &armservicefabricmanagedclusters.SubResource{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
	// 					},
	// 					EnableAcceleratedNetworking: to.Ptr(true),
	// 					IPConfigurations: []*armservicefabricmanagedclusters.IPConfiguration{
	// 						{
	// 							Name: to.Ptr("ipconfig-1"),
	// 							ApplicationGatewayBackendAddressPools: []*armservicefabricmanagedclusters.SubResource{
	// 								{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
	// 								},
	// 							},
	// 							LoadBalancerBackendAddressPools: []*armservicefabricmanagedclusters.SubResource{
	// 								{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
	// 								},
	// 							},
	// 							LoadBalancerInboundNatPools: []*armservicefabricmanagedclusters.SubResource{
	// 								{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
	// 								},
	// 							},
	// 							PrivateIPAddressVersion: to.Ptr(armservicefabricmanagedclusters.PrivateIPAddressVersionIPv4),
	// 							PublicIPAddressConfiguration: &armservicefabricmanagedclusters.IPConfigurationPublicIPAddressConfiguration{
	// 								Name: to.Ptr("publicip-1"),
	// 								IPTags: []*armservicefabricmanagedclusters.IPTag{
	// 									{
	// 										IPTagType: to.Ptr("RoutingPreference"),
	// 										Tag: to.Ptr("Internet"),
	// 									},
	// 								},
	// 								PublicIPAddressVersion: to.Ptr(armservicefabricmanagedclusters.PublicIPAddressVersionIPv4),
	// 							},
	// 							Subnet: &armservicefabricmanagedclusters.SubResource{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Capacities: map[string]*string{
	// 				"ClientConnections": to.Ptr("65536"),
	// 			},
	// 			DataDiskLetter: to.Ptr("S"),
	// 			DataDiskSizeGB: to.Ptr[int32](200),
	// 			DataDiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypePremiumLRS),
	// 			DscpConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
	// 			EnableAcceleratedNetworking: to.Ptr(true),
	// 			EnableEncryptionAtHost: to.Ptr(true),
	// 			EnableNodePublicIP: to.Ptr(true),
	// 			EnableNodePublicIPv6: to.Ptr(true),
	// 			EnableOverProvisioning: to.Ptr(false),
	// 			EvictionPolicy: to.Ptr(armservicefabricmanagedclusters.EvictionPolicyTypeDeallocate),
	// 			FrontendConfigurations: []*armservicefabricmanagedclusters.FrontendConfiguration{
	// 				{
	// 					ApplicationGatewayBackendAddressPoolID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
	// 					LoadBalancerBackendAddressPoolID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
	// 					LoadBalancerInboundNatPoolID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
	// 				},
	// 			},
	// 			IsPrimary: to.Ptr(false),
	// 			IsSpotVM: to.Ptr(true),
	// 			IsStateless: to.Ptr(true),
	// 			MultiplePlacementGroups: to.Ptr(true),
	// 			NatGatewayID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/natGateways/myNatGateway"),
	// 			PlacementProperties: map[string]*string{
	// 				"HasSSD": to.Ptr("true"),
	// 				"NodeColor": to.Ptr("green"),
	// 				"SomeProperty": to.Ptr("5"),
	// 			},
	// 			ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateCreating),
	// 			SecureBootEnabled: to.Ptr(true),
	// 			SecurityType: to.Ptr(armservicefabricmanagedclusters.SecurityTypeConfidentialVM),
	// 			SecurityEncryptionType: to.Ptr(armservicefabricmanagedclusters.SecurityEncryptionTypeDiskWithVMGuestState),
	// 			ServiceArtifactReferenceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/myVmArtifactProfile"),
	// 			SpotRestoreTimeout: to.Ptr("PT30M"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			UseDefaultPublicLoadBalancer: to.Ptr(true),
	// 			UseEphemeralOSDisk: to.Ptr(true),
	// 			VMApplications: []*armservicefabricmanagedclusters.VMApplication{
	// 				{
	// 					ConfigurationReference: to.Ptr("https://mystorageaccount.blob.core.windows.net/containername/blobname"),
	// 					EnableAutomaticUpgrade: to.Ptr(true),
	// 					Order: to.Ptr[int32](1),
	// 					PackageReferenceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/applications/myApplication/versions/1.0.0"),
	// 					TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 					VMGalleryTags: to.Ptr("{\"Tag1\":\"Value1\",\"Tag2\":\"Value2\"}"),
	// 				},
	// 			},
	// 			VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
	// 				{
	// 					Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
	// 					Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
	// 						Type: to.Ptr("GenevaMonitoring"),
	// 						AutoUpgradeMinorVersion: to.Ptr(true),
	// 						EnableAutomaticUpgrade: to.Ptr(true),
	// 						ForceUpdateTag: to.Ptr("v.1.0"),
	// 						Publisher: to.Ptr("Microsoft.Azure.Geneva"),
	// 						Settings: map[string]any{
	// 						},
	// 						SetupOrder: []*armservicefabricmanagedclusters.VmssExtensionSetupOrder{
	// 							to.Ptr(armservicefabricmanagedclusters.VmssExtensionSetupOrderBeforeSFRuntime),
	// 						},
	// 						TypeHandlerVersion: to.Ptr("2.0"),
	// 					},
	// 				},
	// 			},
	// 			VMImageOffer: to.Ptr("WindowsServer"),
	// 			VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
	// 			VMImageSKU: to.Ptr("2016-Datacenter-Server-Core"),
	// 			VMImageVersion: to.Ptr("latest"),
	// 			VMInstanceCount: to.Ptr[int32](10),
	// 			VMManagedIdentity: &armservicefabricmanagedclusters.VMManagedIdentity{
	// 				UserAssignedIdentities: []*string{
	// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
	// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2"),
	// 				},
	// 			},
	// 			VMSecrets: []*armservicefabricmanagedclusters.VaultSecretGroup{
	// 				{
	// 					SourceVault: &armservicefabricmanagedclusters.SubResource{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
	// 					},
	// 					VaultCertificates: []*armservicefabricmanagedclusters.VaultCertificate{
	// 						{
	// 							CertificateStore: to.Ptr("My"),
	// 							CertificateURL: to.Ptr("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			VMSetupActions: []*armservicefabricmanagedclusters.VMSetupAction{
	// 				to.Ptr(armservicefabricmanagedclusters.VMSetupActionEnableContainers),
	// 				to.Ptr(armservicefabricmanagedclusters.VMSetupActionEnableHyperV),
	// 			},
	// 			VMSize: to.Ptr("Standard_DS3"),
	// 			Zones: []*string{
	// 				to.Ptr("1"),
	// 				to.Ptr("2"),
	// 				to.Ptr("3"),
	// 			},
	// 			IsOutboundOnly: to.Ptr(true),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
