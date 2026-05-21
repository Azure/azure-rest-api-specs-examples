package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/Job_CreateorUpdate_ConnectedEnvironment.json
func ExampleJobsClient_BeginCreateOrUpdate_createOrUpdateContainerAppsJobOnAConnectedEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerAppsJob0", armappcontainers.Job{
		ExtendedLocation: &armappcontainers.ExtendedLocation{
			Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
			Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
		},
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.JobProperties{
			Configuration: &armappcontainers.JobConfiguration{
				ManualTriggerConfig: &armappcontainers.JobConfigurationManualTriggerConfig{
					Parallelism:            to.Ptr[int32](4),
					ReplicaCompletionCount: to.Ptr[int32](1),
				},
				ReplicaRetryLimit: to.Ptr[int32](10),
				ReplicaTimeout:    to.Ptr[int32](10),
				TriggerType:       to.Ptr(armappcontainers.TriggerTypeManual),
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube"),
			Template: &armappcontainers.JobTemplate{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("testcontainerAppsJob0"),
						Image: to.Ptr("repo/testcontainerAppsJob0:v1"),
						Probes: []*armappcontainers.ContainerAppProbe{
							{
								Type: to.Ptr(armappcontainers.TypeLiveness),
								HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
									Path: to.Ptr("/health"),
									HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
										{
											Name:  to.Ptr("Custom-Header"),
											Value: to.Ptr("Awesome"),
										},
									},
									Port: to.Ptr[int32](8080),
								},
								InitialDelaySeconds: to.Ptr[int32](5),
								PeriodSeconds:       to.Ptr[int32](3),
							},
						},
					},
				},
				InitContainers: []*armappcontainers.InitContainer{
					{
						Name: to.Ptr("testinitcontainerAppsJob0"),
						Args: []*string{
							to.Ptr("-c"),
							to.Ptr("while true; do echo hello; sleep 10;done"),
						},
						Command: []*string{
							to.Ptr("/bin/sh"),
						},
						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](0.2),
							Memory: to.Ptr("100Mi"),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.JobsClientCreateOrUpdateResponse{
	// 	Job: armappcontainers.Job{
	// 		Name: to.Ptr("testcontainerAppsJob0"),
	// 		Type: to.Ptr("Microsoft.App/jobs"),
	// 		ExtendedLocation: &armappcontainers.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 			Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerAppsJob0"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.JobProperties{
	// 			Configuration: &armappcontainers.JobConfiguration{
	// 				ManualTriggerConfig: &armappcontainers.JobConfigurationManualTriggerConfig{
	// 					Parallelism: to.Ptr[int32](4),
	// 					ReplicaCompletionCount: to.Ptr[int32](1),
	// 				},
	// 				ReplicaRetryLimit: to.Ptr[int32](10),
	// 				ReplicaTimeout: to.Ptr[int32](10),
	// 				TriggerType: to.Ptr(armappcontainers.TriggerTypeManual),
	// 			},
	// 			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube"),
	// 			EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 			ProvisioningState: to.Ptr(armappcontainers.JobProvisioningStateInProgress),
	// 			Template: &armappcontainers.JobTemplate{
	// 				Containers: []*armappcontainers.Container{
	// 					{
	// 						Name: to.Ptr("testcontainerAppsJob0"),
	// 						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
	// 						Probes: []*armappcontainers.ContainerAppProbe{
	// 							{
	// 								Type: to.Ptr(armappcontainers.TypeLiveness),
	// 								HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
	// 									Path: to.Ptr("/health"),
	// 									HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
	// 										{
	// 											Name: to.Ptr("Custom-Header"),
	// 											Value: to.Ptr("Awesome"),
	// 										},
	// 									},
	// 									Port: to.Ptr[int32](8080),
	// 								},
	// 								InitialDelaySeconds: to.Ptr[int32](3),
	// 								PeriodSeconds: to.Ptr[int32](3),
	// 							},
	// 						},
	// 						Resources: &armappcontainers.ContainerResources{
	// 							CPU: to.Ptr[float64](0.2),
	// 							Memory: to.Ptr("100Mi"),
	// 						},
	// 					},
	// 				},
	// 				InitContainers: []*armappcontainers.InitContainer{
	// 					{
	// 						Name: to.Ptr("testinitcontainerAppsJob0"),
	// 						Args: []*string{
	// 							to.Ptr("-c"),
	// 							to.Ptr("while true; do echo hello; sleep 10;done"),
	// 						},
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh"),
	// 						},
	// 						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
	// 						Resources: &armappcontainers.ContainerResources{
	// 							CPU: to.Ptr[float64](0.2),
	// 							Memory: to.Ptr("100Mi"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
