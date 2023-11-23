package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveGet.json
func ExampleArchivesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.Archive = armcontainerregistry.Archive{
	// 	Name: to.Ptr("myArchiveName"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/packages/archives"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/myPackageType/archives/myArchiveName"),
	// 	SystemData: &armcontainerregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armcontainerregistry.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armcontainerregistry.ArchiveProperties{
	// 		PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
	// 			Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
	// 			URL: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		PublishedVersion: to.Ptr("string"),
	// 	},
	// }
}
