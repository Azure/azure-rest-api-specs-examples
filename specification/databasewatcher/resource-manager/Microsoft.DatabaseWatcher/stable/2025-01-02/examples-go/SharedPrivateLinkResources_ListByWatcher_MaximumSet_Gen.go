package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/SharedPrivateLinkResources_ListByWatcher_MaximumSet_Gen.json
func ExampleSharedPrivateLinkResourcesClient_NewListByWatcherPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedPrivateLinkResourcesClient().NewListByWatcherPager("apiTest-ddat4p", "databasemo3ej9ih", nil)
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
		// page = armdatabasewatcher.SharedPrivateLinkResourcesClientListByWatcherResponse{
		// 	SharedPrivateLinkResourceListResult: armdatabasewatcher.SharedPrivateLinkResourceListResult{
		// 		Value: []*armdatabasewatcher.SharedPrivateLinkResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih/sharedPrivateLinkResources/monitoringh22eed"),
		// 				Name: to.Ptr("monitoringh22eed"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers/sharedPrivateLinkResources"),
		// 				Properties: &armdatabasewatcher.SharedPrivateLinkResourceProperties{
		// 					PrivateLinkResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.KeyVault/vaults/kvmo3ej9ih"),
		// 					GroupID: to.Ptr("vault"),
		// 					RequestMessage: to.Ptr("request message"),
		// 					DNSZone: to.Ptr("ec3ae9d410ba"),
		// 					Status: to.Ptr(armdatabasewatcher.SharedPrivateLinkResourceStatusPending),
		// 				},
		// 				SystemData: &armdatabasewatcher.SystemData{
		// 					CreatedBy: to.Ptr("enbpvlpqbwd"),
		// 					CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("mxp"),
		// 					LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/awfba"),
		// 	},
		// }
	}
}
