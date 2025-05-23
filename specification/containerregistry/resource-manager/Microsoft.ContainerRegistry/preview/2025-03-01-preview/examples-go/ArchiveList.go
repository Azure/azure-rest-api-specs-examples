package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ArchiveList.json
func ExampleArchivesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ArchiveListResult = armcontainerregistry.ArchiveListResult{
		// 	Value: []*armcontainerregistry.Archive{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armcontainerregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armcontainerregistry.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armcontainerregistry.ArchiveProperties{
		// 				PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
		// 					Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
		// 					URL: to.Ptr("string"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				PublishedVersion: to.Ptr("string"),
		// 			},
		// 	}},
		// }
	}
}
