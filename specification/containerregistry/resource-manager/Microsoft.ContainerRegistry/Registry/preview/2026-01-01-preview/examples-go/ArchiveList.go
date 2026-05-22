package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/ArchiveList.json
func ExampleArchivesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArchivesClient().NewListPager("myResourceGroup", "myRegistry", "myPackageType", nil)
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
		// page = armcontainerregistry.ArchivesClientListResponse{
		// 	ArchiveListResult: armcontainerregistry.ArchiveListResult{
		// 		Value: []*armcontainerregistry.Archive{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/myPackageType/archives/myArchiveName"),
		// 				Name: to.Ptr("myArchiveName"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/archives"),
		// 				SystemData: &armcontainerregistry.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 				},
		// 				Properties: &armcontainerregistry.ArchiveProperties{
		// 					PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
		// 						Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
		// 						URL: to.Ptr("string"),
		// 					},
		// 					PublishedVersion: to.Ptr("string"),
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateCreating),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("string"),
		// 	},
		// }
	}
}
