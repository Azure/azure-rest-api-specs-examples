package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/Job_ProxyGet.json
func ExampleJobsClient_ProxyGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().ProxyGet(ctx, "rg", "testcontainerAppsJob0", "rootApi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.JobsClientProxyGetResponse{
	// 	Job: armappcontainers.Job{
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerAppsJob0/detectorproperties/rootApi"),
	// 		Name: to.Ptr("testcontainerAppsJob0"),
	// 		Type: to.Ptr("Microsoft.App/jobs"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.JobProperties{
	// 			ProvisioningState: to.Ptr(armappcontainers.JobProvisioningStateSucceeded),
	// 			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 			Configuration: &armappcontainers.JobConfiguration{
	// 				ReplicaTimeout: to.Ptr[int32](10),
	// 				ReplicaRetryLimit: to.Ptr[int32](10),
	// 				ManualTriggerConfig: &armappcontainers.JobConfigurationManualTriggerConfig{
	// 					ReplicaCompletionCount: to.Ptr[int32](1),
	// 					Parallelism: to.Ptr[int32](4),
	// 				},
	// 				TriggerType: to.Ptr(armappcontainers.TriggerTypeManual),
	// 			},
	// 			Template: &armappcontainers.JobTemplate{
	// 				Containers: []*armappcontainers.Container{
	// 					{
	// 						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
	// 						Name: to.Ptr("testcontainerAppsJob0"),
	// 						Resources: &armappcontainers.ContainerResources{
	// 							CPU: to.Ptr[float64](0.2),
	// 							Memory: to.Ptr("100Mi"),
	// 						},
	// 					},
	// 				},
	// 				InitContainers: []*armappcontainers.InitContainer{
	// 					{
	// 						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
	// 						Name: to.Ptr("testinitcontainerAppsJob0"),
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
