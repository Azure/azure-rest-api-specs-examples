package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/LabelHistory_List.json
func ExampleContainerAppsLabelHistoryClient_NewListLabelHistoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsLabelHistoryClient().NewListLabelHistoryPager("rg", "testContainerApp", &armappcontainers.ContainerAppsLabelHistoryClientListLabelHistoryOptions{Filter: nil})
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
		// page.LabelHistoryCollection = armappcontainers.LabelHistoryCollection{
		// 	Value: []*armappcontainers.LabelHistory{
		// 		{
		// 			Name: to.Ptr("dev"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/labelHistory"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testContainerApp/labelHistory/dev"),
		// 			Properties: &armappcontainers.LabelHistoryProperties{
		// 				Records: []*armappcontainers.LabelHistoryRecordItem{
		// 					{
		// 						Revision: to.Ptr("testContainerApp--2"),
		// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-15T05:49:47.000Z"); return t}()),
		// 						Status: to.Ptr(armappcontainers.StatusStarting),
		// 					},
		// 					{
		// 						Revision: to.Ptr("testContainerApp--1"),
		// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-15T01:12:56.000Z"); return t}()),
		// 						Status: to.Ptr(armappcontainers.StatusFailed),
		// 						Stop: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-15T05:45:08.000Z"); return t}()),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("prod"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/labelHistory"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testContainerApp/labelHistory/prod"),
		// 			Properties: &armappcontainers.LabelHistoryProperties{
		// 				Records: []*armappcontainers.LabelHistoryRecordItem{
		// 					{
		// 						Revision: to.Ptr("testContainerApp--1"),
		// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-15T05:45:08.000Z"); return t}()),
		// 						Status: to.Ptr(armappcontainers.StatusSucceeded),
		// 						Stop: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-15T05:49:47.000Z"); return t}()),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
