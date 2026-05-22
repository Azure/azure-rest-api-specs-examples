package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/ArchiveGet.json
func ExampleArchivesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArchivesClient().Get(ctx, "myResourceGroup", "myRegistry", "myPackageType", "myArchiveName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.ArchivesClientGetResponse{
	// 	Archive: &armcontainerregistry.Archive{
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/packages/archives"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/myPackageType/archives/myArchiveName"),
	// 		Name: to.Ptr("myArchiveName"),
	// 		SystemData: &armcontainerregistry.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 		},
	// 		Properties: &armcontainerregistry.ArchiveProperties{
	// 			PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
	// 				Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
	// 				URL: to.Ptr("string"),
	// 			},
	// 			PublishedVersion: to.Ptr("string"),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateCreating),
	// 		},
	// 	},
	// }
}
