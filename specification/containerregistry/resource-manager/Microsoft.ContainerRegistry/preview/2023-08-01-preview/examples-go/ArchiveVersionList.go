package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ArchiveVersionList.json
func ExampleArchiveVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArchiveVersionsClient().NewListPager("myResourceGroup", "myRegistry", "myPackageType", "myArchiveName", nil)
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
		// page.ArchiveVersionListResult = armcontainerregistry.ArchiveVersionListResult{
		// 	Value: []*armcontainerregistry.ArchiveVersion{
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
		// 			Properties: &armcontainerregistry.ArchiveVersionProperties{
		// 				ArchiveVersionErrorMessage: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
