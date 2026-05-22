package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/Services_ListSkus.json
func ExampleServicesClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListSKUsPager("DmsSdkRg", "DmsSdkService", nil)
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
		// page = armdatamigration.ServicesClientListSKUsResponse{
		// 	ServiceSKUList: armdatamigration.ServiceSKUList{
		// 		Value: []*armdatamigration.AvailableServiceSKU{
		// 			{
		// 				ResourceType: to.Ptr("Microsoft.DataMigration/services"),
		// 				SKU: &armdatamigration.AvailableServiceSKUSKU{
		// 					Name: to.Ptr("Basic_1vCore"),
		// 					Size: to.Ptr("1 vCore"),
		// 					Tier: to.Ptr("Basic"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
