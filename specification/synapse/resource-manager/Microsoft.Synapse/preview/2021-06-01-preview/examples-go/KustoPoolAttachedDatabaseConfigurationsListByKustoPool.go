package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsListByKustoPool.json
func ExampleKustoPoolAttachedDatabaseConfigurationsClient_NewListByKustoPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolAttachedDatabaseConfigurationsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByKustoPoolPager("kustorptest", "kustoclusterrptest4", "kustorptest", nil)
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
		// page.AttachedDatabaseConfigurationListResult = armsynapse.AttachedDatabaseConfigurationListResult{
		// 	Value: []*armsynapse.AttachedDatabaseConfiguration{
		// 		{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Synapse/Workspaces/KustoPools/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/kustoclusterrptest4/AttachedDatabaseConfigurations/KustoDatabase8"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.AttachedDatabaseConfigurationProperties{
		// 				KustoPoolResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armsynapse.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Synapse/Workspaces/KustoPools/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/kustoclusterrptest4/AttachedDatabaseConfigurations/KustoDatabase9"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.AttachedDatabaseConfigurationProperties{
		// 				KustoPoolResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armsynapse.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 				TableLevelSharingProperties: &armsynapse.TableLevelSharingProperties{
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
