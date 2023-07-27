package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAttachedDatabaseConfigurationsClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster2", "attachedDatabaseConfigurationsTest", armkusto.AttachedDatabaseConfiguration{
		Location: to.Ptr("westus"),
		Properties: &armkusto.AttachedDatabaseConfigurationProperties{
			ClusterResourceID:                 to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
			DatabaseName:                      to.Ptr("kustodatabase"),
			DatabaseNameOverride:              to.Ptr("overridekustodatabase"),
			DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
			TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
				ExternalTablesToExclude: []*string{
					to.Ptr("ExternalTable2")},
				ExternalTablesToInclude: []*string{
					to.Ptr("ExternalTable1")},
				MaterializedViewsToExclude: []*string{
					to.Ptr("MaterializedViewTable2")},
				MaterializedViewsToInclude: []*string{
					to.Ptr("MaterializedViewTable1")},
				TablesToExclude: []*string{
					to.Ptr("Table2")},
				TablesToInclude: []*string{
					to.Ptr("Table1")},
			},
		},
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
	// res.AttachedDatabaseConfiguration = armkusto.AttachedDatabaseConfiguration{
	// 	Name: to.Ptr("kustoCluster2/attachedDatabaseConfigurationsTest"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/attachedDatabaseConfigurations/attachedDatabaseConfigurationsTest"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkusto.AttachedDatabaseConfigurationProperties{
	// 		ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
	// 		DatabaseName: to.Ptr("db1"),
	// 		DatabaseNameOverride: to.Ptr("overridekustodatabase"),
	// 		DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
	// 			ExternalTablesToExclude: []*string{
	// 				to.Ptr("ExternalTable2")},
	// 				ExternalTablesToInclude: []*string{
	// 					to.Ptr("ExternalTable1")},
	// 					MaterializedViewsToExclude: []*string{
	// 						to.Ptr("MaterializedViewTable2")},
	// 						MaterializedViewsToInclude: []*string{
	// 							to.Ptr("MaterializedViewTable1")},
	// 							TablesToExclude: []*string{
	// 								to.Ptr("Table2")},
	// 								TablesToInclude: []*string{
	// 									to.Ptr("Table1")},
	// 								},
	// 							},
	// 						}
}
