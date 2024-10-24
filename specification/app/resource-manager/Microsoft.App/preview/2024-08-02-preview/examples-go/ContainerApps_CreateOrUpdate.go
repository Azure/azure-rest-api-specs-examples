package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_CreateOrUpdate.json
func ExampleContainerAppsClient_BeginCreateOrUpdate_createOrUpdateContainerApp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerApp0", armappcontainers.ContainerApp{
		Location: to.Ptr("East US"),
		Identity: &armappcontainers.ManagedServiceIdentity{
			Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
				"/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
			},
		},
		Properties: &armappcontainers.ContainerAppProperties{
			Configuration: &armappcontainers.Configuration{
				Dapr: &armappcontainers.Dapr{
					AppPort:            to.Ptr[int32](3000),
					AppProtocol:        to.Ptr(armappcontainers.AppProtocolHTTP),
					EnableAPILogging:   to.Ptr(true),
					Enabled:            to.Ptr(true),
					HTTPMaxRequestSize: to.Ptr[int32](10),
					HTTPReadBufferSize: to.Ptr[int32](30),
					LogLevel:           to.Ptr(armappcontainers.LogLevelDebug),
				},
				IdentitySettings: []*armappcontainers.IdentitySettings{
					{
						Identity:  to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
						Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleAll),
					},
					{
						Identity:  to.Ptr("system"),
						Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleInit),
					}},
				Ingress: &armappcontainers.Ingress{
					AdditionalPortMappings: []*armappcontainers.IngressPortMapping{
						{
							External:   to.Ptr(true),
							TargetPort: to.Ptr[int32](1234),
						},
						{
							ExposedPort: to.Ptr[int32](3456),
							External:    to.Ptr(false),
							TargetPort:  to.Ptr[int32](2345),
						}},
					ClientCertificateMode: to.Ptr(armappcontainers.IngressClientCertificateModeAccept),
					CorsPolicy: &armappcontainers.CorsPolicy{
						AllowCredentials: to.Ptr(true),
						AllowedHeaders: []*string{
							to.Ptr("HEADER1"),
							to.Ptr("HEADER2")},
						AllowedMethods: []*string{
							to.Ptr("GET"),
							to.Ptr("POST")},
						AllowedOrigins: []*string{
							to.Ptr("https://a.test.com"),
							to.Ptr("https://b.test.com")},
						ExposeHeaders: []*string{
							to.Ptr("HEADER3"),
							to.Ptr("HEADER4")},
						MaxAge: to.Ptr[int32](1234),
					},
					CustomDomains: []*armappcontainers.CustomDomain{
						{
							Name:          to.Ptr("www.my-name.com"),
							BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
							CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
						},
						{
							Name:          to.Ptr("www.my-other-name.com"),
							BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
							CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
						}},
					External: to.Ptr(true),
					IPSecurityRestrictions: []*armappcontainers.IPSecurityRestrictionRule{
						{
							Name:           to.Ptr("Allow work IP A subnet"),
							Description:    to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
							Action:         to.Ptr(armappcontainers.ActionAllow),
							IPAddressRange: to.Ptr("192.168.1.1/32"),
						},
						{
							Name:           to.Ptr("Allow work IP B subnet"),
							Description:    to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
							Action:         to.Ptr(armappcontainers.ActionAllow),
							IPAddressRange: to.Ptr("192.168.1.1/8"),
						}},
					StickySessions: &armappcontainers.IngressStickySessions{
						Affinity: to.Ptr(armappcontainers.AffinitySticky),
					},
					TargetPort: to.Ptr[int32](3000),
					Traffic: []*armappcontainers.TrafficWeight{
						{
							Label:        to.Ptr("production"),
							RevisionName: to.Ptr("testcontainerApp0-ab1234"),
							Weight:       to.Ptr[int32](100),
						}},
				},
				MaxInactiveRevisions: to.Ptr[int32](10),
				Runtime: &armappcontainers.Runtime{
					Dotnet: &armappcontainers.RuntimeDotnet{
						AutoConfigureDataProtection: to.Ptr(true),
					},
					Java: &armappcontainers.RuntimeJava{
						EnableMetrics: to.Ptr(true),
						JavaAgent: &armappcontainers.RuntimeJavaAgent{
							Enabled: to.Ptr(true),
							Logging: &armappcontainers.RuntimeJavaAgentLogging{
								LoggerSettings: []*armappcontainers.LoggerSetting{
									{
										Level:  to.Ptr(armappcontainers.LevelDebug),
										Logger: to.Ptr("org.springframework.boot"),
									}},
							},
						},
					},
				},
				Service: &armappcontainers.Service{
					Type: to.Ptr("redis"),
				},
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			Template: &armappcontainers.Template{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("testcontainerApp0"),
						Image: to.Ptr("repo/testcontainerApp0:v1"),
						VolumeMounts: []*armappcontainers.VolumeMount{
							{
								MountPath:  to.Ptr("/mnt/path1"),
								SubPath:    to.Ptr("subPath1"),
								VolumeName: to.Ptr("azurefile"),
							},
							{
								MountPath:  to.Ptr("/mnt/path2"),
								SubPath:    to.Ptr("subPath2"),
								VolumeName: to.Ptr("nfsazurefile"),
							}},
						Probes: []*armappcontainers.ContainerAppProbe{
							{
								Type: to.Ptr(armappcontainers.TypeLiveness),
								HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
									Path: to.Ptr("/health"),
									HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
										{
											Name:  to.Ptr("Custom-Header"),
											Value: to.Ptr("Awesome"),
										}},
									Port: to.Ptr[int32](8080),
								},
								InitialDelaySeconds: to.Ptr[int32](3),
								PeriodSeconds:       to.Ptr[int32](3),
							}},
					}},
				InitContainers: []*armappcontainers.InitContainer{
					{
						Name: to.Ptr("testinitcontainerApp0"),
						Args: []*string{
							to.Ptr("-c"),
							to.Ptr("while true; do echo hello; sleep 10;done")},
						Command: []*string{
							to.Ptr("/bin/sh")},
						Image: to.Ptr("repo/testcontainerApp0:v4"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](0.2),
							Memory: to.Ptr("100Mi"),
						},
					}},
				Scale: &armappcontainers.Scale{
					CooldownPeriod:  to.Ptr[int32](350),
					MaxReplicas:     to.Ptr[int32](5),
					MinReplicas:     to.Ptr[int32](1),
					PollingInterval: to.Ptr[int32](35),
					Rules: []*armappcontainers.ScaleRule{
						{
							Name: to.Ptr("httpscalingrule"),
							Custom: &armappcontainers.CustomScaleRule{
								Type: to.Ptr("http"),
								Metadata: map[string]*string{
									"concurrentRequests": to.Ptr("50"),
								},
							},
						},
						{
							Name: to.Ptr("servicebus"),
							Custom: &armappcontainers.CustomScaleRule{
								Type:     to.Ptr("azure-servicebus"),
								Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
								Metadata: map[string]*string{
									"messageCount": to.Ptr("5"),
									"namespace":    to.Ptr("mynamespace"),
									"queueName":    to.Ptr("myqueue"),
								},
							},
						},
						{
							Name: to.Ptr("azure-queue"),
							AzureQueue: &armappcontainers.QueueScaleRule{
								AccountName: to.Ptr("account1"),
								Identity:    to.Ptr("system"),
								QueueLength: to.Ptr[int32](1),
								QueueName:   to.Ptr("queue1"),
							},
						}},
				},
				ServiceBinds: []*armappcontainers.ServiceBind{
					{
						Name:       to.Ptr("redisService"),
						ClientType: to.Ptr("dotnet"),
						CustomizedKeys: map[string]*string{
							"DesiredKey": to.Ptr("defaultKey"),
						},
						ServiceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/redisService"),
					}},
				Volumes: []*armappcontainers.Volume{
					{
						Name:        to.Ptr("azurefile"),
						StorageName: to.Ptr("storage"),
						StorageType: to.Ptr(armappcontainers.StorageTypeAzureFile),
					},
					{
						Name:        to.Ptr("nfsazurefile"),
						StorageName: to.Ptr("nfsStorage"),
						StorageType: to.Ptr(armappcontainers.StorageTypeNfsAzureFile),
					}},
			},
			WorkloadProfileName: to.Ptr("My-GP-01"),
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
	// res.ContainerApp = armappcontainers.ContainerApp{
	// 	Name: to.Ptr("testcontainerApp0"),
	// 	Type: to.Ptr("Microsoft.App/containerApps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0"),
	// 	Location: to.Ptr("East US"),
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("24adfa4f-dedf-8dc0-ca29-b6d1a69ab319"),
	// 		TenantID: to.Ptr("23adfa4f-eedf-1dc0-ba29-a6d1a69ab3d0"),
	// 		UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
	// 			"/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": &armappcontainers.UserAssignedIdentity{
	// 				ClientID: to.Ptr("14adfa4f-eedf-1dc0-ba29-a6d1a69ab3df"),
	// 				PrincipalID: to.Ptr("74adfa4f-dedf-8dc0-ca29-b6d1a69ab312"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armappcontainers.ContainerAppProperties{
	// 		Configuration: &armappcontainers.Configuration{
	// 			Dapr: &armappcontainers.Dapr{
	// 				AppPort: to.Ptr[int32](3000),
	// 				AppProtocol: to.Ptr(armappcontainers.AppProtocolHTTP),
	// 				Enabled: to.Ptr(true),
	// 				HTTPReadBufferSize: to.Ptr[int32](30),
	// 			},
	// 			IdentitySettings: []*armappcontainers.IdentitySettings{
	// 				{
	// 					Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 					Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleAll),
	// 				},
	// 				{
	// 					Identity: to.Ptr("system"),
	// 					Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleInit),
	// 			}},
	// 			Ingress: &armappcontainers.Ingress{
	// 				CustomDomains: []*armappcontainers.CustomDomain{
	// 					{
	// 						Name: to.Ptr("www.my-name.com"),
	// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
	// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
	// 					},
	// 					{
	// 						Name: to.Ptr("www.my-other-name.com"),
	// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
	// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
	// 				}},
	// 				External: to.Ptr(true),
	// 				Fqdn: to.Ptr("testcontainerApp0.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 				IPSecurityRestrictions: []*armappcontainers.IPSecurityRestrictionRule{
	// 					{
	// 						Name: to.Ptr("Allow work IP A subnet"),
	// 						Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
	// 						Action: to.Ptr(armappcontainers.ActionAllow),
	// 						IPAddressRange: to.Ptr("192.168.1.1/32"),
	// 					},
	// 					{
	// 						Name: to.Ptr("Allow work IP B subnet"),
	// 						Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
	// 						Action: to.Ptr(armappcontainers.ActionAllow),
	// 						IPAddressRange: to.Ptr("192.168.1.1/8"),
	// 				}},
	// 				StickySessions: &armappcontainers.IngressStickySessions{
	// 					Affinity: to.Ptr(armappcontainers.AffinitySticky),
	// 				},
	// 				TargetPort: to.Ptr[int32](3000),
	// 				TargetPortHTTPScheme: to.Ptr(armappcontainers.IngressTargetPortHTTPSchemeHTTP),
	// 				Traffic: []*armappcontainers.TrafficWeight{
	// 					{
	// 						RevisionName: to.Ptr("testcontainerApp0-ab1234"),
	// 						Weight: to.Ptr[int32](80),
	// 					},
	// 					{
	// 						Label: to.Ptr("staging"),
	// 						RevisionName: to.Ptr("testcontainerApp0-ab4321"),
	// 						Weight: to.Ptr[int32](20),
	// 				}},
	// 				Transport: to.Ptr(armappcontainers.IngressTransportMethodAuto),
	// 			},
	// 			Runtime: &armappcontainers.Runtime{
	// 				Dotnet: &armappcontainers.RuntimeDotnet{
	// 					AutoConfigureDataProtection: to.Ptr(true),
	// 				},
	// 				Java: &armappcontainers.RuntimeJava{
	// 					EnableMetrics: to.Ptr(true),
	// 					JavaAgent: &armappcontainers.RuntimeJavaAgent{
	// 						Enabled: to.Ptr(true),
	// 						Logging: &armappcontainers.RuntimeJavaAgentLogging{
	// 							LoggerSettings: []*armappcontainers.LoggerSetting{
	// 								{
	// 									Level: to.Ptr(armappcontainers.LevelDebug),
	// 									Logger: to.Ptr("org.springframework.boot"),
	// 							}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 		LatestReadyRevisionName: to.Ptr("testcontainerApp0-pjxhsye"),
	// 		LatestRevisionFqdn: to.Ptr("testcontainerApp0-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 		ManagedEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerApp0"),
	// 					Image: to.Ptr("repo/testcontainerApp0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 					VolumeMounts: []*armappcontainers.VolumeMount{
	// 						{
	// 							MountPath: to.Ptr("/mnt/path1"),
	// 							SubPath: to.Ptr("subPath1"),
	// 							VolumeName: to.Ptr("azurefile"),
	// 						},
	// 						{
	// 							MountPath: to.Ptr("/mnt/path2"),
	// 							SubPath: to.Ptr("subPath2"),
	// 							VolumeName: to.Ptr("nfsazurefile"),
	// 					}},
	// 					Probes: []*armappcontainers.ContainerAppProbe{
	// 						{
	// 							Type: to.Ptr(armappcontainers.TypeLiveness),
	// 							HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
	// 								Path: to.Ptr("/health"),
	// 								HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
	// 									{
	// 										Name: to.Ptr("Custom-Header"),
	// 										Value: to.Ptr("Awesome"),
	// 								}},
	// 								Port: to.Ptr[int32](8080),
	// 							},
	// 							InitialDelaySeconds: to.Ptr[int32](3),
	// 							PeriodSeconds: to.Ptr[int32](3),
	// 					}},
	// 			}},
	// 			InitContainers: []*armappcontainers.InitContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainerApp0"),
	// 					Args: []*string{
	// 						to.Ptr("-c"),
	// 						to.Ptr("while true; do echo hello; sleep 10;done")},
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh")},
	// 							Image: to.Ptr("repo/testcontainerApp0:v4"),
	// 							Resources: &armappcontainers.ContainerResources{
	// 								CPU: to.Ptr[float64](0.2),
	// 								Memory: to.Ptr("100Mi"),
	// 							},
	// 					}},
	// 					Scale: &armappcontainers.Scale{
	// 						CooldownPeriod: to.Ptr[int32](350),
	// 						MaxReplicas: to.Ptr[int32](5),
	// 						MinReplicas: to.Ptr[int32](1),
	// 						PollingInterval: to.Ptr[int32](35),
	// 						Rules: []*armappcontainers.ScaleRule{
	// 							{
	// 								Name: to.Ptr("httpscalingrule"),
	// 								HTTP: &armappcontainers.HTTPScaleRule{
	// 									Metadata: map[string]*string{
	// 										"concurrentRequests": to.Ptr("50"),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("servicebus"),
	// 								Custom: &armappcontainers.CustomScaleRule{
	// 									Type: to.Ptr("azure-servicebus"),
	// 									Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 									Metadata: map[string]*string{
	// 										"messageCount": to.Ptr("5"),
	// 										"namespace": to.Ptr("mynamespace"),
	// 										"queueName": to.Ptr("myqueue"),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("azure-queue"),
	// 								AzureQueue: &armappcontainers.QueueScaleRule{
	// 									AccountName: to.Ptr("account1"),
	// 									Identity: to.Ptr("system"),
	// 									QueueLength: to.Ptr[int32](1),
	// 									QueueName: to.Ptr("queue1"),
	// 								},
	// 						}},
	// 					},
	// 					Volumes: []*armappcontainers.Volume{
	// 						{
	// 							Name: to.Ptr("azurefile"),
	// 							StorageName: to.Ptr("storage"),
	// 							StorageType: to.Ptr(armappcontainers.StorageTypeAzureFile),
	// 						},
	// 						{
	// 							Name: to.Ptr("nfsazurefile"),
	// 							StorageName: to.Ptr("nfsStorage"),
	// 							StorageType: to.Ptr(armappcontainers.StorageTypeNfsAzureFile),
	// 					}},
	// 				},
	// 				WorkloadProfileName: to.Ptr("My-GP-01"),
	// 			},
	// 		}
}
