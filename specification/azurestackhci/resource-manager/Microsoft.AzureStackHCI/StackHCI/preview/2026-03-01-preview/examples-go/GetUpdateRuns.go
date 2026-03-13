package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/GetUpdateRuns.json
func ExampleUpdateRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("b8d594e5-51f3-4c11-9c54-a7771b81c712", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpdateRunsClient().Get(ctx, "testrg", "testcluster", "Microsoft4.2203.2.32", "23b779ba-0d52-4a80-8571-45ca74664ec3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.UpdateRunsClientGetResponse{
	// 	UpdateRun: &armazurestackhci.UpdateRun{
	// 		Name: to.Ptr("Microsoft4.2203.2.32/23b779ba-0d52-4a80-8571-45ca74664ec3"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/updates/updateRuns"),
	// 		ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/clusters/testcluster/updates/Microsoft4.2203.2.32/updateRuns/23b779ba-0d52-4a80-8571-45ca74664ec3"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armazurestackhci.UpdateRunProperties{
	// 			Progress: &armazurestackhci.Step{
	// 				Name: to.Ptr("Unnamed step"),
	// 				Description: to.Ptr("Update Azure Stack."),
	// 				EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T13:58:42.969006+00:00"); return t}()),
	// 				ErrorMessage: to.Ptr(""),
	// 				LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T13:58:42.969006+00:00"); return t}()),
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:36:33.3876751+00:00"); return t}()),
	// 				Status: to.Ptr("Success"),
	// 				Steps: []*armazurestackhci.Step{
	// 					{
	// 						Name: to.Ptr("PreUpdate Cloud"),
	// 						Description: to.Ptr("Prepare for SSU update"),
	// 						EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:37:16.8728314+00:00"); return t}()),
	// 						ErrorMessage: to.Ptr(""),
	// 						LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:37:16.8728314+00:00"); return t}()),
	// 						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T01:36:33.3876751+00:00"); return t}()),
	// 						Status: to.Ptr("Success"),
	// 						Steps: []*armazurestackhci.Step{
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
