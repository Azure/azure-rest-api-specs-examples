package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoAttachedDatabaseConfigurationsListByCluster.json
func ExampleAttachedDatabaseConfigurationsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAttachedDatabaseConfigurationsClient().NewListByClusterPager("kustorptest", "kustoCluster2", nil)
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
		// page.AttachedDatabaseConfigurationListResult = armkusto.AttachedDatabaseConfigurationListResult{
		// 	Value: []*armkusto.AttachedDatabaseConfiguration{
		// 		{
		// 			Name: to.Ptr("kustoCluster2/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/AttachedDatabaseConfigurations/KustoDatabase8"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster2/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/AttachedDatabaseConfigurations/KustoDatabase9"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
		// 					ExternalTablesToExclude: []*string{
		// 						to.Ptr("ExternalTable2")},
		// 						ExternalTablesToInclude: []*string{
		// 							to.Ptr("ExternalTable1")},
		// 							FunctionsToExclude: []*string{
		// 								to.Ptr("functionsToExclude2")},
		// 								FunctionsToInclude: []*string{
		// 									to.Ptr("functionsToInclude1")},
		// 									MaterializedViewsToExclude: []*string{
		// 										to.Ptr("MaterializedViewTable2")},
		// 										MaterializedViewsToInclude: []*string{
		// 											to.Ptr("MaterializedViewTable1")},
		// 											TablesToExclude: []*string{
		// 												to.Ptr("Table2")},
		// 												TablesToInclude: []*string{
		// 													to.Ptr("Table1")},
		// 												},
		// 											},
		// 									}},
		// 								}
	}
}
