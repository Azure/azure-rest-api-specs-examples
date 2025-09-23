package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/Jobs_Resume.json
func ExampleJobsClient_BeginResume() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginResume(ctx, "rg", "testcontainerAppsJob0", nil)
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
	// 	Name: to.Ptr("testcontainerAppsJob0"),
	// 	Type: to.Ptr("Microsoft.App/jobs"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/testcontainerAppsJob0"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.JobProperties{
	// 		Configuration: &armappcontainers.JobConfiguration{
	// 			IdentitySettings: []*armappcontainers.IdentitySettings{
	// 				{
	// 					Identity: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 					Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleAll),
	// 				},
	// 				{
	// 					Identity: to.Ptr("system"),
	// 					Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleInit),
	// 			}},
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
	// 		RunningState: to.Ptr(armappcontainers.JobRunningStateReady),
	// 		Template: &armappcontainers.JobTemplate{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerAppsJob0"),
	// 					Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
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
	// 			}},
	// 			InitContainers: []*armappcontainers.InitContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainerAppsJob0"),
	// 					Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 			}},
	// 			Volumes: []*armappcontainers.Volume{
	// 				{
	// 					Name: to.Ptr("azurefile"),
	// 					StorageName: to.Ptr("storage"),
	// 					StorageType: to.Ptr(armappcontainers.StorageTypeAzureFile),
	// 				},
	// 				{
	// 					Name: to.Ptr("nfsazurefile"),
	// 					StorageName: to.Ptr("nfsStorage"),
	// 					StorageType: to.Ptr(armappcontainers.StorageTypeNfsAzureFile),
	// 			}},
	// 		},
	// 	},
	// }
}
