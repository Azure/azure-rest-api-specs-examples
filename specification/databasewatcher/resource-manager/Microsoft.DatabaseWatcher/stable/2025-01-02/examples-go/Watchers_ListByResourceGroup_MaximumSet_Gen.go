package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/Watchers_ListByResourceGroup_MaximumSet_Gen.json
func ExampleWatchersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("A76F9850-996B-40B3-94D4-C98110A0EEC9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchersClient().NewListByResourceGroupPager("rgWatcher", nil)
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
		// page = armdatabasewatcher.WatchersClientListByResourceGroupResponse{
		// 	WatcherListResult: armdatabasewatcher.WatcherListResult{
		// 		Value: []*armdatabasewatcher.Watcher{
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/dummyrg/providers/Microsoft.DatabaseWatcher/watchers/myWatcher"),
		// 				Name: to.Ptr("myWatcher"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
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
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-4sz1jg/providers/Microsoft.DatabaseWatcher/watchers/databasemo4o4zdf"),
		// 				Name: to.Ptr("databasemo4o4zdf"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
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
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-d77ftn/providers/Microsoft.DatabaseWatcher/watchers/databasemosn3h6l"),
		// 				Name: to.Ptr("databasemosn3h6l"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
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
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-nyb4gm/providers/Microsoft.DatabaseWatcher/watchers/databasemoyb6iar"),
		// 				Name: to.Ptr("databasemoyb6iar"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
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
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-llis4j/providers/Microsoft.DatabaseWatcher/watchers/databasemoi04xst"),
		// 				Name: to.Ptr("databasemoi04xst"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
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
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-thy6zd/providers/Microsoft.DatabaseWatcher/watchers/databasemonpyl24"),
		// 				Name: to.Ptr("databasemonpyl24"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				SystemData: &armdatabasewatcher.SystemData{
		// 					CreatedBy: to.Ptr("ysoqerxnmxqsvhmvjojoyzotc"),
		// 					CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T18:01:18.690Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("yrilzsg"),
		// 					LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T18:01:18.690Z"); return t}()),
		// 				},
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.DatabaseWatcher/watchers/databasemo3ej9ih"),
		// 				Name: to.Ptr("databasemo3ej9ih"),
		// 				Type: to.Ptr("microsoft.databasewatcher/watchers"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armdatabasewatcher.WatcherProperties{
		// 					Status: to.Ptr(armdatabasewatcher.WatcherStatusStarting),
		// 					ProvisioningState: to.Ptr(armdatabasewatcher.ProvisioningStateSucceeded),
		// 					DefaultAlertRuleIdentityResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/3pmtest"),
		// 					Datastore: &armdatabasewatcher.Datastore{
		// 						AdxClusterResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto"),
		// 						KustoClusterURI: to.Ptr("https://kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoClusterDisplayName: to.Ptr("kustoUri-adx"),
		// 						KustoDataIngestionURI: to.Ptr("https://ingest-kustouri-adx.eastus.kusto.windows.net"),
		// 						KustoDatabaseName: to.Ptr("kustoDatabaseName1"),
		// 						KustoManagementURL: to.Ptr("https://portal.azure.com/"),
		// 						KustoOfferingType: to.Ptr(armdatabasewatcher.KustoOfferingTypeAdx),
		// 					},
		// 				},
		// 				Identity: &armdatabasewatcher.ManagedServiceIdentityV4{
		// 					UserAssignedIdentities: map[string]*armdatabasewatcher.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 					TenantID: to.Ptr("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d"),
		// 				},
		// 				Tags: map[string]*string{
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
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
