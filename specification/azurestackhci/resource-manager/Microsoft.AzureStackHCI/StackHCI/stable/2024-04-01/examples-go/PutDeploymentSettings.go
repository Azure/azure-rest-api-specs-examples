package armazurestackhci_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutDeploymentSettings.json
func ExampleDeploymentSettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentSettingsClient().BeginCreateOrUpdate(ctx, "test-rg", "myCluster", "default", armazurestackhci.DeploymentSetting{
		Properties: &armazurestackhci.DeploymentSettingsProperties{
			ArcNodeResourceIDs: []*string{
				to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
				to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2")},
			DeploymentConfiguration: &armazurestackhci.DeploymentConfiguration{
				ScaleUnits: []*armazurestackhci.ScaleUnits{
					{
						DeploymentData: &armazurestackhci.DeploymentData{
							AdouPath: to.Ptr("OU=ms169,DC=ASZ1PLab8,DC=nttest,DC=microsoft,DC=com"),
							Cluster: &armazurestackhci.DeploymentCluster{
								Name:                 to.Ptr("testHCICluster"),
								AzureServiceEndpoint: to.Ptr("core.windows.net"),
								CloudAccountName:     to.Ptr("myasestoragacct"),
								WitnessPath:          to.Ptr("Cloud"),
								WitnessType:          to.Ptr("Cloud"),
							},
							DomainFqdn: to.Ptr("ASZ1PLab8.nttest.microsoft.com"),
							HostNetwork: &armazurestackhci.DeploymentSettingHostNetwork{
								EnableStorageAutoIP: to.Ptr(false),
								Intents: []*armazurestackhci.DeploymentSettingIntents{
									{
										Name: to.Ptr("Compute_Management"),
										Adapter: []*string{
											to.Ptr("Port2")},
										AdapterPropertyOverrides: &armazurestackhci.DeploymentSettingAdapterPropertyOverrides{
											JumboPacket:             to.Ptr("1514"),
											NetworkDirect:           to.Ptr("Enabled"),
											NetworkDirectTechnology: to.Ptr("iWARP"),
										},
										OverrideAdapterProperty:            to.Ptr(false),
										OverrideQosPolicy:                  to.Ptr(false),
										OverrideVirtualSwitchConfiguration: to.Ptr(false),
										QosPolicyOverrides: &armazurestackhci.QosPolicyOverrides{
											BandwidthPercentageSMB:         to.Ptr("50"),
											PriorityValue8021ActionCluster: to.Ptr("7"),
											PriorityValue8021ActionSMB:     to.Ptr("3"),
										},
										TrafficType: []*string{
											to.Ptr("Compute"),
											to.Ptr("Management")},
										VirtualSwitchConfigurationOverrides: &armazurestackhci.DeploymentSettingVirtualSwitchConfigurationOverrides{
											EnableIov:              to.Ptr("True"),
											LoadBalancingAlgorithm: to.Ptr("HyperVPort"),
										},
									}},
								StorageConnectivitySwitchless: to.Ptr(true),
								StorageNetworks: []*armazurestackhci.DeploymentSettingStorageNetworks{
									{
										Name:               to.Ptr("Storage1Network"),
										NetworkAdapterName: to.Ptr("Port3"),
										StorageAdapterIPInfo: []*armazurestackhci.DeploymentSettingStorageAdapterIPInfo{
											{
												IPv4Address:  to.Ptr("10.57.48.60"),
												PhysicalNode: to.Ptr("string"),
												SubnetMask:   to.Ptr("255.255.248.0"),
											}},
										VlanID: to.Ptr("5"),
									}},
							},
							InfrastructureNetwork: []*armazurestackhci.InfrastructureNetwork{
								{
									DNSServers: []*string{
										to.Ptr("10.57.50.90")},
									Gateway: to.Ptr("255.255.248.0"),
									IPPools: []*armazurestackhci.IPPools{
										{
											EndingAddress:   to.Ptr("10.57.48.66"),
											StartingAddress: to.Ptr("10.57.48.60"),
										}},
									SubnetMask: to.Ptr("255.255.248.0"),
								}},
							NamingPrefix: to.Ptr("ms169"),
							Observability: &armazurestackhci.Observability{
								EpisodicDataUpload:  to.Ptr(true),
								EuLocation:          to.Ptr(false),
								StreamingDataClient: to.Ptr(true),
							},
							OptionalServices: &armazurestackhci.OptionalServices{
								CustomLocation: to.Ptr("customLocationName"),
							},
							PhysicalNodes: []*armazurestackhci.PhysicalNodes{
								{
									Name:        to.Ptr("ms169host"),
									IPv4Address: to.Ptr("10.57.51.224"),
								},
								{
									Name:        to.Ptr("ms154host"),
									IPv4Address: to.Ptr("10.57.53.236"),
								}},
							SdnIntegration: &armazurestackhci.SdnIntegration{
								NetworkController: &armazurestackhci.NetworkController{
									MacAddressPoolStart:          to.Ptr("00-0D-3A-1B-C7-21"),
									MacAddressPoolStop:           to.Ptr("00-0D-3A-1B-C7-29"),
									NetworkVirtualizationEnabled: to.Ptr(true),
								},
							},
							Secrets: []*armazurestackhci.EceDeploymentSecrets{
								{
									EceSecretName:  to.Ptr(armazurestackhci.EceSecrets("BMCAdminUserCred")),
									SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
									SecretName:     to.Ptr("cluster1-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
								},
								{
									EceSecretName:  to.Ptr(armazurestackhci.EceSecretsAzureStackLCMUserCredential),
									SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4c"),
									SecretName:     to.Ptr("cluster2-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
								}},
							SecretsLocation: to.Ptr("/subscriptions/db4e2fdb-6d80-4e6e-b7cd-xxxxxxx/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/abcd123"),
							SecuritySettings: &armazurestackhci.DeploymentSecuritySettings{
								BitlockerBootVolume:           to.Ptr(true),
								BitlockerDataVolumes:          to.Ptr(true),
								CredentialGuardEnforced:       to.Ptr(false),
								DriftControlEnforced:          to.Ptr(true),
								DrtmProtection:                to.Ptr(true),
								HvciProtection:                to.Ptr(true),
								SideChannelMitigationEnforced: to.Ptr(true),
								SmbClusterEncryption:          to.Ptr(false),
								SmbSigningEnforced:            to.Ptr(true),
								WdacEnforced:                  to.Ptr(true),
							},
							Storage: &armazurestackhci.Storage{
								ConfigurationMode: to.Ptr("Express"),
							},
						},
						SbePartnerInfo: &armazurestackhci.SbePartnerInfo{
							CredentialList: []*armazurestackhci.SbeCredentials{
								{
									EceSecretName:  to.Ptr("DownloadConnectorCred"),
									SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
									SecretName:     to.Ptr("cluster1-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
								}},
							PartnerProperties: []*armazurestackhci.SbePartnerProperties{
								{
									Name:  to.Ptr("EnableBMCIpV6"),
									Value: to.Ptr("false"),
								},
								{
									Name:  to.Ptr("PhoneHomePort"),
									Value: to.Ptr("1653"),
								},
								{
									Name:  to.Ptr("BMCSecurityState"),
									Value: to.Ptr("HighSecurity"),
								}},
							SbeDeploymentInfo: &armazurestackhci.SbeDeploymentInfo{
								Family:                  to.Ptr("Gen5"),
								Publisher:               to.Ptr("Contoso"),
								SbeManifestCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T02:40:33.000Z"); return t }()),
								SbeManifestSource:       to.Ptr("default"),
								Version:                 to.Ptr("4.0.2309.13"),
							},
						},
					}},
				Version: to.Ptr("string"),
			},
			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
			OperationType:  to.Ptr(armazurestackhci.OperationTypeClusterProvisioning),
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
	// res.DeploymentSetting = armazurestackhci.DeploymentSetting{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/clusters/deploymentSettings"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/deploymentSettings/default"),
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// 	Properties: &armazurestackhci.DeploymentSettingsProperties{
	// 		ArcNodeResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
	// 			to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2")},
	// 			DeploymentConfiguration: &armazurestackhci.DeploymentConfiguration{
	// 				ScaleUnits: []*armazurestackhci.ScaleUnits{
	// 					{
	// 						DeploymentData: &armazurestackhci.DeploymentData{
	// 							AdouPath: to.Ptr("OU=ms169,DC=ASZ1PLab8,DC=nttest,DC=microsoft,DC=com"),
	// 							Cluster: &armazurestackhci.DeploymentCluster{
	// 								Name: to.Ptr("testHCICluster"),
	// 								AzureServiceEndpoint: to.Ptr("core.windows.net"),
	// 								CloudAccountName: to.Ptr("myasestoragacct"),
	// 								WitnessPath: to.Ptr("Cloud"),
	// 								WitnessType: to.Ptr("Cloud"),
	// 							},
	// 							DomainFqdn: to.Ptr("ASZ1PLab8.nttest.microsoft.com"),
	// 							HostNetwork: &armazurestackhci.DeploymentSettingHostNetwork{
	// 								EnableStorageAutoIP: to.Ptr(false),
	// 								Intents: []*armazurestackhci.DeploymentSettingIntents{
	// 									{
	// 										Name: to.Ptr("Compute_Management"),
	// 										Adapter: []*string{
	// 											to.Ptr("Port2")},
	// 											AdapterPropertyOverrides: &armazurestackhci.DeploymentSettingAdapterPropertyOverrides{
	// 												JumboPacket: to.Ptr("1514"),
	// 												NetworkDirect: to.Ptr("Enabled"),
	// 												NetworkDirectTechnology: to.Ptr("iWARP"),
	// 											},
	// 											OverrideAdapterProperty: to.Ptr(false),
	// 											OverrideQosPolicy: to.Ptr(false),
	// 											OverrideVirtualSwitchConfiguration: to.Ptr(false),
	// 											QosPolicyOverrides: &armazurestackhci.QosPolicyOverrides{
	// 												BandwidthPercentageSMB: to.Ptr("50"),
	// 												PriorityValue8021ActionCluster: to.Ptr("7"),
	// 												PriorityValue8021ActionSMB: to.Ptr("3"),
	// 											},
	// 											TrafficType: []*string{
	// 												to.Ptr("Compute"),
	// 												to.Ptr("Management")},
	// 												VirtualSwitchConfigurationOverrides: &armazurestackhci.DeploymentSettingVirtualSwitchConfigurationOverrides{
	// 													EnableIov: to.Ptr("True"),
	// 													LoadBalancingAlgorithm: to.Ptr("HyperVPort"),
	// 												},
	// 										}},
	// 										StorageConnectivitySwitchless: to.Ptr(true),
	// 										StorageNetworks: []*armazurestackhci.DeploymentSettingStorageNetworks{
	// 											{
	// 												Name: to.Ptr("Storage1Network"),
	// 												NetworkAdapterName: to.Ptr("Port3"),
	// 												StorageAdapterIPInfo: []*armazurestackhci.DeploymentSettingStorageAdapterIPInfo{
	// 													{
	// 														IPv4Address: to.Ptr("10.57.48.60"),
	// 														PhysicalNode: to.Ptr("string"),
	// 														SubnetMask: to.Ptr("255.255.248.0"),
	// 												}},
	// 												VlanID: to.Ptr("5"),
	// 										}},
	// 									},
	// 									InfrastructureNetwork: []*armazurestackhci.InfrastructureNetwork{
	// 										{
	// 											DNSServers: []*string{
	// 												to.Ptr("10.57.50.90")},
	// 												Gateway: to.Ptr("255.255.248.0"),
	// 												IPPools: []*armazurestackhci.IPPools{
	// 													{
	// 														EndingAddress: to.Ptr("10.57.48.66"),
	// 														StartingAddress: to.Ptr("10.57.48.60"),
	// 												}},
	// 												SubnetMask: to.Ptr("255.255.248.0"),
	// 										}},
	// 										NamingPrefix: to.Ptr("ms169"),
	// 										Observability: &armazurestackhci.Observability{
	// 											EpisodicDataUpload: to.Ptr(true),
	// 											EuLocation: to.Ptr(false),
	// 											StreamingDataClient: to.Ptr(true),
	// 										},
	// 										OptionalServices: &armazurestackhci.OptionalServices{
	// 											CustomLocation: to.Ptr("customLocationName"),
	// 										},
	// 										PhysicalNodes: []*armazurestackhci.PhysicalNodes{
	// 											{
	// 												Name: to.Ptr("ms169host"),
	// 												IPv4Address: to.Ptr("10.57.51.224"),
	// 											},
	// 											{
	// 												Name: to.Ptr("ms154host"),
	// 												IPv4Address: to.Ptr("10.57.53.236"),
	// 										}},
	// 										SdnIntegration: &armazurestackhci.SdnIntegration{
	// 											NetworkController: &armazurestackhci.NetworkController{
	// 												MacAddressPoolStart: to.Ptr("00-0D-3A-1B-C7-21"),
	// 												MacAddressPoolStop: to.Ptr("00-0D-3A-1B-C7-29"),
	// 												NetworkVirtualizationEnabled: to.Ptr(true),
	// 											},
	// 										},
	// 										Secrets: []*armazurestackhci.EceDeploymentSecrets{
	// 											{
	// 												EceSecretName: to.Ptr(armazurestackhci.EceSecrets("BMCAdminUserCred")),
	// 												SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
	// 												SecretName: to.Ptr("cluster1-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
	// 											},
	// 											{
	// 												EceSecretName: to.Ptr(armazurestackhci.EceSecretsAzureStackLCMUserCredential),
	// 												SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4c"),
	// 												SecretName: to.Ptr("cluster2-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
	// 										}},
	// 										SecretsLocation: to.Ptr("/subscriptions/db4e2fdb-6d80-4e6e-b7cd-xxxxxxx/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/abcd123"),
	// 										SecuritySettings: &armazurestackhci.DeploymentSecuritySettings{
	// 											BitlockerBootVolume: to.Ptr(true),
	// 											BitlockerDataVolumes: to.Ptr(true),
	// 											CredentialGuardEnforced: to.Ptr(false),
	// 											DriftControlEnforced: to.Ptr(true),
	// 											DrtmProtection: to.Ptr(true),
	// 											HvciProtection: to.Ptr(true),
	// 											SideChannelMitigationEnforced: to.Ptr(true),
	// 											SmbClusterEncryption: to.Ptr(false),
	// 											SmbSigningEnforced: to.Ptr(true),
	// 											WdacEnforced: to.Ptr(true),
	// 										},
	// 										Storage: &armazurestackhci.Storage{
	// 											ConfigurationMode: to.Ptr("Express"),
	// 										},
	// 									},
	// 									SbePartnerInfo: &armazurestackhci.SbePartnerInfo{
	// 										CredentialList: []*armazurestackhci.SbeCredentials{
	// 											{
	// 												EceSecretName: to.Ptr("DownloadConnectorCred"),
	// 												SecretLocation: to.Ptr("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
	// 												SecretName: to.Ptr("cluster1-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63"),
	// 										}},
	// 										PartnerProperties: []*armazurestackhci.SbePartnerProperties{
	// 											{
	// 												Name: to.Ptr("EnableBMCIpV6"),
	// 												Value: to.Ptr("false"),
	// 											},
	// 											{
	// 												Name: to.Ptr("PhoneHomePort"),
	// 												Value: to.Ptr("1653"),
	// 											},
	// 											{
	// 												Name: to.Ptr("BMCSecurityState"),
	// 												Value: to.Ptr("HighSecurity"),
	// 										}},
	// 										SbeDeploymentInfo: &armazurestackhci.SbeDeploymentInfo{
	// 											Family: to.Ptr("Gen5"),
	// 											Publisher: to.Ptr("Contoso"),
	// 											SbeManifestCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T02:40:33.000Z"); return t}()),
	// 											SbeManifestSource: to.Ptr("default"),
	// 											Version: to.Ptr("4.0.2309.13"),
	// 										},
	// 									},
	// 							}},
	// 							Version: to.Ptr("string"),
	// 						},
	// 						DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 						ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 						ReportedProperties: &armazurestackhci.EceReportedProperties{
	// 							DeploymentStatus: &armazurestackhci.EceActionStatus{
	// 								Status: to.Ptr("Error"),
	// 								Steps: []*armazurestackhci.DeploymentStep{
	// 									{
	// 										Name: to.Ptr("Cloud Deployment"),
	// 										Description: to.Ptr("Deploy Cloud."),
	// 										EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 										Exception: []*string{
	// 											to.Ptr("exception1"),
	// 											to.Ptr("exception2")},
	// 											FullStepIndex: to.Ptr("0"),
	// 											StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 											Status: to.Ptr("Error"),
	// 											Steps: []*armazurestackhci.DeploymentStep{
	// 												{
	// 													Name: to.Ptr("Before Cloud Deployment"),
	// 													EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 													Exception: []*string{
	// 														to.Ptr("exception1"),
	// 														to.Ptr("exception2")},
	// 														FullStepIndex: to.Ptr("0.1"),
	// 														StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 														Steps: []*armazurestackhci.DeploymentStep{
	// 														},
	// 													},
	// 													{
	// 														Name: to.Ptr("Clean up temporary content"),
	// 														EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 														Exception: []*string{
	// 															to.Ptr("exception1"),
	// 															to.Ptr("exception2")},
	// 															FullStepIndex: to.Ptr("0.36"),
	// 															StartTimeUTC: to.Ptr("2023-06-09T03:58:37"),
	// 															Status: to.Ptr("Error"),
	// 															Steps: []*armazurestackhci.DeploymentStep{
	// 															},
	// 													}},
	// 											}},
	// 										},
	// 										ValidationStatus: &armazurestackhci.EceActionStatus{
	// 											Status: to.Ptr("Error"),
	// 											Steps: []*armazurestackhci.DeploymentStep{
	// 												{
	// 													Name: to.Ptr("Cloud Deployment"),
	// 													Description: to.Ptr("Deploy Cloud."),
	// 													EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 													Exception: []*string{
	// 														to.Ptr("exception1"),
	// 														to.Ptr("exception2")},
	// 														FullStepIndex: to.Ptr("0"),
	// 														StartTimeUTC: to.Ptr("2023-06-09T00:08:19"),
	// 														Status: to.Ptr("Error"),
	// 														Steps: []*armazurestackhci.DeploymentStep{
	// 															{
	// 																Name: to.Ptr("Before Cloud Deployment"),
	// 																Description: to.Ptr("Before Cloud Deployment"),
	// 																EndTimeUTC: to.Ptr("2023-06-09T01:10:10"),
	// 																Exception: []*string{
	// 																	to.Ptr("exception1"),
	// 																	to.Ptr("exception2")},
	// 																	FullStepIndex: to.Ptr("0.1"),
	// 																	StartTimeUTC: to.Ptr("2023-06-09T00:08:23"),
	// 																	Steps: []*armazurestackhci.DeploymentStep{
	// 																	},
	// 																},
	// 																{
	// 																	Name: to.Ptr("Clean up temporary content"),
	// 																	Description: to.Ptr("Clean up temporary content"),
	// 																	EndTimeUTC: to.Ptr("2023-06-09T04:01:47"),
	// 																	Exception: []*string{
	// 																		to.Ptr("exception1"),
	// 																		to.Ptr("exception2")},
	// 																		FullStepIndex: to.Ptr("0.36"),
	// 																		StartTimeUTC: to.Ptr("2023-06-09T03:58:37"),
	// 																		Status: to.Ptr("Error"),
	// 																		Steps: []*armazurestackhci.DeploymentStep{
	// 																		},
	// 																}},
	// 														}},
	// 													},
	// 												},
	// 											},
	// 										}
}
