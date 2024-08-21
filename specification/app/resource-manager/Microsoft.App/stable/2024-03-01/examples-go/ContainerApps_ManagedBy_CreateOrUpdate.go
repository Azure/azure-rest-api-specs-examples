package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerApps_ManagedBy_CreateOrUpdate.json
func ExampleContainerAppsClient_BeginCreateOrUpdate_createOrUpdateManagedByApp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerappmanagedby", armappcontainers.ContainerApp{
		Location:  to.Ptr("East US"),
		ManagedBy: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp"),
		Properties: &armappcontainers.ContainerAppProperties{
			Configuration: &armappcontainers.Configuration{
				Ingress: &armappcontainers.Ingress{
					ExposedPort: to.Ptr[int32](4000),
					External:    to.Ptr(true),
					TargetPort:  to.Ptr[int32](3000),
					Traffic: []*armappcontainers.TrafficWeight{
						{
							RevisionName: to.Ptr("testcontainerappmanagedby-ab1234"),
							Weight:       to.Ptr[int32](100),
						}},
					Transport: to.Ptr(armappcontainers.IngressTransportMethodTCP),
				},
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			Template: &armappcontainers.Template{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("testcontainerappmanagedby"),
						Image: to.Ptr("repo/testcontainerappmanagedby:v1"),
						Probes: []*armappcontainers.ContainerAppProbe{
							{
								Type:                to.Ptr(armappcontainers.TypeLiveness),
								InitialDelaySeconds: to.Ptr[int32](3),
								PeriodSeconds:       to.Ptr[int32](3),
								TCPSocket: &armappcontainers.ContainerAppProbeTCPSocket{
									Port: to.Ptr[int32](8080),
								},
							}},
					}},
				Scale: &armappcontainers.Scale{
					MaxReplicas: to.Ptr[int32](5),
					MinReplicas: to.Ptr[int32](1),
					Rules: []*armappcontainers.ScaleRule{
						{
							Name: to.Ptr("tcpscalingrule"),
							TCP: &armappcontainers.TCPScaleRule{
								Metadata: map[string]*string{
									"concurrentConnections": to.Ptr("50"),
								},
							},
						}},
				},
			},
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
	// 	Name: to.Ptr("testcontainerappmanagedby"),
	// 	Type: to.Ptr("Microsoft.App/containerApps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerappmanagedby"),
	// 	Location: to.Ptr("East US"),
	// 	ManagedBy: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp"),
	// 	Properties: &armappcontainers.ContainerAppProperties{
	// 		Configuration: &armappcontainers.Configuration{
	// 			Ingress: &armappcontainers.Ingress{
	// 				ExposedPort: to.Ptr[int32](4000),
	// 				External: to.Ptr(true),
	// 				Fqdn: to.Ptr("testcontainerappmanagedby.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 				TargetPort: to.Ptr[int32](3000),
	// 				Traffic: []*armappcontainers.TrafficWeight{
	// 					{
	// 						RevisionName: to.Ptr("testcontainerappmanagedby-ab1234"),
	// 						Weight: to.Ptr[int32](80),
	// 					},
	// 					{
	// 						RevisionName: to.Ptr("testcontainerappmanagedby-ab4321"),
	// 						Weight: to.Ptr[int32](20),
	// 				}},
	// 				Transport: to.Ptr(armappcontainers.IngressTransportMethodTCP),
	// 			},
	// 		},
	// 		LatestReadyRevisionName: to.Ptr("testcontainerappmanagedby-pjxhsye"),
	// 		LatestRevisionFqdn: to.Ptr("testcontainerappmanagedby-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 		ManagedEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerappmanagedby"),
	// 					Image: to.Ptr("repo/testcontainerappmanagedby:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.5),
	// 						Memory: to.Ptr("1Gi"),
	// 					},
	// 					Probes: []*armappcontainers.ContainerAppProbe{
	// 						{
	// 							Type: to.Ptr(armappcontainers.TypeLiveness),
	// 							InitialDelaySeconds: to.Ptr[int32](3),
	// 							PeriodSeconds: to.Ptr[int32](3),
	// 							TCPSocket: &armappcontainers.ContainerAppProbeTCPSocket{
	// 								Port: to.Ptr[int32](8080),
	// 							},
	// 					}},
	// 			}},
	// 			Scale: &armappcontainers.Scale{
	// 				MaxReplicas: to.Ptr[int32](5),
	// 				MinReplicas: to.Ptr[int32](1),
	// 				Rules: []*armappcontainers.ScaleRule{
	// 					{
	// 						Name: to.Ptr("tcpscalingrule"),
	// 						TCP: &armappcontainers.TCPScaleRule{
	// 							Metadata: map[string]*string{
	// 								"concurrentConnections": to.Ptr("50"),
	// 							},
	// 						},
	// 				}},
	// 			},
	// 		},
	// 	},
	// }
}
