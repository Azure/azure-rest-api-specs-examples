package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoDatabasesListByCluster.json
func ExampleDatabasesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByClusterPager("kustorptest", "kustoCluster", nil)
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
		// page.DatabaseListResult = armkusto.DatabaseListResult{
		// 	Value: []armkusto.DatabaseClassification{
		// 		&armkusto.ReadWriteDatabase{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
		// 			Kind: to.Ptr(armkusto.KindReadWrite),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.ReadWriteDatabaseProperties{
		// 				HotCachePeriod: to.Ptr("P1D"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 			},
		// 		},
		// 		&armkusto.ReadOnlyFollowingDatabase{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase9"),
		// 			Kind: to.Ptr(armkusto.KindReadOnlyFollowing),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.ReadOnlyFollowingDatabaseProperties{
		// 				AttachedDatabaseConfigurationName: to.Ptr("attachedDatabaseConfigurationsTest"),
		// 				DatabaseShareOrigin: to.Ptr(armkusto.DatabaseShareOriginDirect),
		// 				HotCachePeriod: to.Ptr("P1D"),
		// 				LeaderClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
		// 				OriginalDatabaseName: to.Ptr("KustoDatabase9"),
		// 				PrincipalsModificationKind: to.Ptr(armkusto.PrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 				Statistics: &armkusto.DatabaseStatistics{
		// 					Size: to.Ptr[float32](1024),
		// 				},
		// 				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
		// 					ExternalTablesToExclude: []*string{
		// 						to.Ptr("ExternalTable2")},
		// 						ExternalTablesToInclude: []*string{
		// 							to.Ptr("ExternalTable1")},
		// 							MaterializedViewsToExclude: []*string{
		// 								to.Ptr("MaterializedViewTable2")},
		// 								MaterializedViewsToInclude: []*string{
		// 									to.Ptr("MaterializedViewTable1")},
		// 									TablesToExclude: []*string{
		// 										to.Ptr("Table2")},
		// 										TablesToInclude: []*string{
		// 											to.Ptr("Table1")},
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}
