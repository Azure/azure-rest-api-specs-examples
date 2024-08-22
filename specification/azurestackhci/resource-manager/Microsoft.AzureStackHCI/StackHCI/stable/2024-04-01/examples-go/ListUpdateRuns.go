package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ListUpdateRuns.json
func ExampleUpdateRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUpdateRunsClient().NewListPager("testrg", "testcluster", "Microsoft4.2203.2.32", nil)
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
		// page.UpdateRunList = armazurestackhci.UpdateRunList{
		// 	Value: []*armazurestackhci.UpdateRun{
		// 		{
		// 			Name: to.Ptr("Microsoft4.2203.2.32/23b779ba-0d52-4a80-8571-45ca74664ec3"),
		// 			Type: to.Ptr("Microsoft.AzureStackHCI/updates/updateRuns"),
		// 			ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/clusters/testcluster/updates/Microsoft4.2203.2.32/updateRuns/23b779ba-0d52-4a80-8571-45ca74664ec3"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armazurestackhci.UpdateRunProperties{
		// 				Progress: &armazurestackhci.Step{
		// 					Name: to.Ptr("Unnamed step"),
		// 					Description: to.Ptr("Update Azure Stack."),
		// 					EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T13:58:42.969Z"); return t}()),
		// 					ErrorMessage: to.Ptr(""),
		// 					LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T13:58:42.969Z"); return t}()),
		// 					StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:36:33.387Z"); return t}()),
		// 					Status: to.Ptr("Success"),
		// 					Steps: []*armazurestackhci.Step{
		// 						{
		// 							Name: to.Ptr("PreUpdate Cloud"),
		// 							Description: to.Ptr("Prepare for SSU update"),
		// 							EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:37:16.872Z"); return t}()),
		// 							ErrorMessage: to.Ptr(""),
		// 							LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:37:16.872Z"); return t}()),
		// 							StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:36:33.387Z"); return t}()),
		// 							Status: to.Ptr("Success"),
		// 							Steps: []*armazurestackhci.Step{
		// 							},
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
