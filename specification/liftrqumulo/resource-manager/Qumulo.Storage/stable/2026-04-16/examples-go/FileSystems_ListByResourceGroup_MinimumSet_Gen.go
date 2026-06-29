package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo/v2"
)

// Generated from example definition: 2026-04-16/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMaximumSetGenGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("53BA951C-DA09-400A-AB3A-F8E98F317423", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgQumulo", nil)
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
		// page = armqumulo.FileSystemsClientListByResourceGroupResponse{
		// 	FileSystemResourceListResult: armqumulo.FileSystemResourceListResult{
		// 		Value: []*armqumulo.FileSystemResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/53BA951C-DA09-400A-AB3A-F8E98F317423/resourceGroups/rgQumulo/providers/Qumulo.Storage/fileSystems/qumulo-fs-01"),
		// 				Name: to.Ptr("qumulo-fs-01"),
		// 				Type: to.Ptr("Qumulo.Storage/fileSystems"),
		// 				Location: to.Ptr("ninyfyhmsedp"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
