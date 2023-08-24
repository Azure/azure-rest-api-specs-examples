package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Jobs_ListBySubscription.json
func ExampleJobsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListBySubscriptionPager(nil)
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
		// page.JobsCollection = armappcontainers.JobsCollection{
		// 	Value: []*armappcontainers.Job{
		// 		{
		// 			Name: to.Ptr("testcontainerAppsJob0"),
		// 			Type: to.Ptr("Microsoft.App/jobs"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerAppsJob0"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.JobProperties{
		// 				Configuration: &armappcontainers.JobConfiguration{
		// 					ManualTriggerConfig: &armappcontainers.JobConfigurationManualTriggerConfig{
		// 						Parallelism: to.Ptr[int32](4),
		// 						ReplicaCompletionCount: to.Ptr[int32](1),
		// 					},
		// 					ReplicaRetryLimit: to.Ptr[int32](10),
		// 					ReplicaTimeout: to.Ptr[int32](10),
		// 					TriggerType: to.Ptr(armappcontainers.TriggerTypeManual),
		// 				},
		// 				EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
		// 				ProvisioningState: to.Ptr(armappcontainers.JobProvisioningStateSucceeded),
		// 				Template: &armappcontainers.JobTemplate{
		// 					Containers: []*armappcontainers.Container{
		// 						{
		// 							Name: to.Ptr("testcontainerAppsJob0"),
		// 							Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					InitContainers: []*armappcontainers.InitContainer{
		// 						{
		// 							Name: to.Ptr("testinitcontainerAppsJob0"),
		// 							Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testcontainerAppsJob1"),
		// 			Type: to.Ptr("Microsoft.App/jobs"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerAppsJob1"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.JobProperties{
		// 				Configuration: &armappcontainers.JobConfiguration{
		// 					ReplicaRetryLimit: to.Ptr[int32](10),
		// 					ReplicaTimeout: to.Ptr[int32](10),
		// 					ScheduleTriggerConfig: &armappcontainers.JobConfigurationScheduleTriggerConfig{
		// 						CronExpression: to.Ptr("* * * * 5"),
		// 						Parallelism: to.Ptr[int32](5),
		// 						ReplicaCompletionCount: to.Ptr[int32](1),
		// 					},
		// 					TriggerType: to.Ptr(armappcontainers.TriggerType("Scheduled")),
		// 				},
		// 				EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
		// 				ProvisioningState: to.Ptr(armappcontainers.JobProvisioningStateSucceeded),
		// 				Template: &armappcontainers.JobTemplate{
		// 					Containers: []*armappcontainers.Container{
		// 						{
		// 							Name: to.Ptr("testcontainerAppsJob1"),
		// 							Image: to.Ptr("repo/testcontainerAppsJob1:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					InitContainers: []*armappcontainers.InitContainer{
		// 						{
		// 							Name: to.Ptr("testinitcontainerAppsJob1"),
		// 							Image: to.Ptr("repo/testcontainerAppsJob1:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
