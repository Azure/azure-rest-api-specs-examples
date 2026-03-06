package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/ApplicationList.json
func ExampleApplicationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationClient().NewListPager("default-azurebatch-japaneast", "sampleacct", nil)
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
		// page = armbatch.ApplicationClientListResponse{
		// 	ListApplicationsResult: armbatch.ListApplicationsResult{
		// 		Value: []*armbatch.Application{
		// 			{
		// 				Name: to.Ptr("app1"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1"),
		// 				Properties: &armbatch.ApplicationProperties{
		// 					AllowUpdates: to.Ptr(false),
		// 					DefaultVersion: to.Ptr("1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("app1"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts/applications"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app2"),
		// 				Properties: &armbatch.ApplicationProperties{
		// 					AllowUpdates: to.Ptr(false),
		// 					DefaultVersion: to.Ptr("2.0"),
		// 					DisplayName: to.Ptr("myAppName"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
