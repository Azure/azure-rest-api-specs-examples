package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_ProxyGet.json
func ExampleJobsClient_ProxyGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().ProxyGet(ctx, "rg", "testcontainerappsjob0", "rootApi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armappcontainers.Job{
	// 	Name: to.Ptr("testcontainerappsjob0"),
	// 	Type: to.Ptr("Microsoft.App/jobs"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerappsjob0/detectorproperties/rootApi"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.JobProperties{
	// 		Configuration: &armappcontainers.JobConfiguration{
	// 			ManualTriggerConfig: &armappcontainers.JobConfigurationManualTriggerConfig{
	// 				Parallelism: to.Ptr[int32](4),
	// 				ReplicaCompletionCount: to.Ptr[int32](1),
	// 			},
	// 			ReplicaRetryLimit: to.Ptr[int32](10),
	// 			ReplicaTimeout: to.Ptr[int32](10),
	// 			TriggerType: to.Ptr(armappcontainers.TriggerTypeManual),
	// 		},
	// 		EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
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
	// 					Image: to.Ptr("repo/testcontainerappsjob0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.5),
	// 						Memory: to.Ptr("1Gi"),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}
