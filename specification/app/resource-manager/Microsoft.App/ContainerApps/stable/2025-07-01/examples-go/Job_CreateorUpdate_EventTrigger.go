package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/Job_CreateorUpdate_EventTrigger.json
func ExampleJobsClient_BeginCreateOrUpdate_createOrUpdateContainerAppsJobWithEventDrivenTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerappsjob0", armappcontainers.Job{
		Location: to.Ptr("East US"),
		Identity: &armappcontainers.ManagedServiceIdentity{
			Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
				"/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
			},
		},
		Properties: &armappcontainers.JobProperties{
			Configuration: &armappcontainers.JobConfiguration{
				EventTriggerConfig: &armappcontainers.JobConfigurationEventTriggerConfig{
					Parallelism:            to.Ptr[int32](4),
					ReplicaCompletionCount: to.Ptr[int32](1),
					Scale: &armappcontainers.JobScale{
						MaxExecutions:   to.Ptr[int32](5),
						MinExecutions:   to.Ptr[int32](1),
						PollingInterval: to.Ptr[int32](40),
						Rules: []*armappcontainers.JobScaleRule{
							{
								Name:     to.Ptr("servicebuscalingrule"),
								Type:     to.Ptr("azure-servicebus"),
								Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
								Metadata: map[string]any{
									"topicName": "my-topic",
								},
							}},
					},
				},
				ReplicaRetryLimit: to.Ptr[int32](10),
				ReplicaTimeout:    to.Ptr[int32](10),
				TriggerType:       to.Ptr(armappcontainers.TriggerTypeEvent),
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			Template: &armappcontainers.JobTemplate{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("testcontainerappsjob0"),
						Image: to.Ptr("repo/testcontainerappsjob0:v1"),
					}},
				InitContainers: []*armappcontainers.InitContainer{
					{
						Name: to.Ptr("testinitcontainerAppsJob0"),
						Args: []*string{
							to.Ptr("-c"),
							to.Ptr("while true; do echo hello; sleep 10;done")},
						Command: []*string{
							to.Ptr("/bin/sh")},
						Image: to.Ptr("repo/testcontainerappsjob0:v4"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](0.5),
							Memory: to.Ptr("1Gi"),
						},
					}},
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
	// res.Job = armappcontainers.Job{
	// 	Name: to.Ptr("testcontainerappsjob0"),
	// 	Type: to.Ptr("Microsoft.App/jobs"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerappsjob0"),
	// 	Location: to.Ptr("East US"),
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
	// 			"/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": &armappcontainers.UserAssignedIdentity{
	// 				ClientID: to.Ptr("14adfa4f-eedf-1dc0-ba29-a6d1a69ab3df"),
	// 				PrincipalID: to.Ptr("74adfa4f-dedf-8dc0-ca29-b6d1a69ab312"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armappcontainers.JobProperties{
	// 		Configuration: &armappcontainers.JobConfiguration{
	// 			EventTriggerConfig: &armappcontainers.JobConfigurationEventTriggerConfig{
	// 				Parallelism: to.Ptr[int32](4),
	// 				ReplicaCompletionCount: to.Ptr[int32](1),
	// 				Scale: &armappcontainers.JobScale{
	// 					MaxExecutions: to.Ptr[int32](5),
	// 					MinExecutions: to.Ptr[int32](1),
	// 					PollingInterval: to.Ptr[int32](20),
	// 					Rules: []*armappcontainers.JobScaleRule{
	// 						{
	// 							Name: to.Ptr("servicebuscalingrule"),
	// 							Type: to.Ptr("azure-servicebus"),
	// 							Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 							Metadata: map[string]any{
	// 								"topicName": "my-topic",
	// 							},
	// 					}},
	// 				},
	// 			},
	// 			ReplicaRetryLimit: to.Ptr[int32](10),
	// 			ReplicaTimeout: to.Ptr[int32](10),
	// 			TriggerType: to.Ptr(armappcontainers.TriggerTypeEvent),
	// 		},
	// 		EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 		EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 		ProvisioningState: to.Ptr(armappcontainers.JobProvisioningStateSucceeded),
	// 		Template: &armappcontainers.JobTemplate{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerappsjob0"),
	// 					Image: to.Ptr("repo/testcontainerappsjob0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.5),
	// 						Memory: to.Ptr("1Gi"),
	// 					},
	// 			}},
	// 			InitContainers: []*armappcontainers.InitContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainerAppsJob0"),
	// 					Args: []*string{
	// 						to.Ptr("-c"),
	// 						to.Ptr("while true; do echo hello; sleep 10;done")},
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh")},
	// 							Image: to.Ptr("repo/testcontainerappsjob0:v4"),
	// 							Resources: &armappcontainers.ContainerResources{
	// 								CPU: to.Ptr[float64](0.5),
	// 								Memory: to.Ptr("1Gi"),
	// 							},
	// 					}},
	// 				},
	// 			},
	// 		}
}
