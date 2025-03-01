package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/Watchers_CreateOrUpdate_MaximumSet_Gen.json
func ExampleWatchersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("A76F9850-996B-40B3-94D4-C98110A0EEC9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginCreateOrUpdate(ctx, "rgWatcher", "testWatcher", armdatabasewatcher.Watcher{
		Properties: &armdatabasewatcher.WatcherProperties{
			Status:                             to.Ptr(armdatabasewatcher.WatcherStatusStarting),
			DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
			Datastore: &armdatabasewatcher.Datastore{
				AdxClusterResourceID:    to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
				KustoClusterURI:         to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
				KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
				KustoDataIngestionURI:   to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
				KustoDatabaseName:       to.Ptr("kustoDatabaseName1"),
				KustoManagementURL:      to.Ptr("https://portal.azure.com/"),
				KustoOfferingType:       to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
			},
		},
		Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
			Type: to.Ptr(armdatabasewatcher.ManagedServiceIdentityTypeSystemAssigned),
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("eastus2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.WatchersClientCreateOrUpdateResponse{
	// 	Watcher: &armdatabasewatcher.Watcher{
	// 		Properties: &armdatabasewatcher.WatcherProperties{
	// 			Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
	// 			DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
	// 			Datastore: &armdatabasewatcher.Datastore{
	// 				AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
	// 				KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
	// 				KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
	// 				KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
	// 				KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
	// 				KustoManagementURL: to.Ptr("https://portal.azure.com/"),
	// 				KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
	// 			},
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
	// 		},
	// 		Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
	// 			Type: to.Ptr(armdatabasewatcher.ManagedServiceIdentityTypeSystemAssigned),
	// 			UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
	// 			TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/dummyrg/providers/Microsoft.DatabaseWatcher/watchers/myWatcher"),
	// 		Name: to.Ptr("myWatcher"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 			CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 			CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mxp"),
	// 			LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		},
	// 	},
	// }
}
