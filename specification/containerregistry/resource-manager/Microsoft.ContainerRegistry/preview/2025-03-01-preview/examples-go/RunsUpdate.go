package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/RunsUpdate.json
func ExampleRunsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRunsClient().Update(ctx, "myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", armcontainerregistry.RunUpdateParameters{
		IsArchiveEnabled: to.Ptr(true),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Run = armcontainerregistry.Run{
	// 	Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Properties: &armcontainerregistry.RunProperties{
	// 		AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 			CPU: to.Ptr[int32](2),
	// 		},
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:13:51.618Z"); return t}()),
	// 		ImageUpdateTrigger: &armcontainerregistry.ImageUpdateTrigger{
	// 			ID: to.Ptr("c0c43143-da5d-41ef-b9e1-e7d749272e88"),
	// 			Images: []*armcontainerregistry.ImageDescriptor{
	// 				{
	// 					Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 					Registry: to.Ptr("registry.hub.docker.com"),
	// 					Repository: to.Ptr("mybaseimage"),
	// 					Tag: to.Ptr("latest"),
	// 			}},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		},
	// 		IsArchiveEnabled: to.Ptr(true),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
	// 		LogArtifact: &armcontainerregistry.ImageDescriptor{
	// 			Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 			Registry: to.Ptr("myregistry"),
	// 			Repository: to.Ptr("acr/tasks"),
	// 			Tag: to.Ptr("mytask-0accec26-d6de-4757-8e74-d080f38eaaab-log"),
	// 		},
	// 		OutputImages: []*armcontainerregistry.ImageDescriptor{
	// 			{
	// 				Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 				Registry: to.Ptr("myregistry.azurecr.io"),
	// 				Repository: to.Ptr("myimage"),
	// 				Tag: to.Ptr("latest"),
	// 		}},
	// 		Platform: &armcontainerregistry.PlatformProperties{
	// 			Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 			OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		RunType: to.Ptr(armcontainerregistry.RunTypeAutoBuild),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:50:51.618Z"); return t}()),
	// 		Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 		Task: to.Ptr("myTask"),
	// 	},
	// }
}
