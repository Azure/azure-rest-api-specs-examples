package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsInWorkspaceWithFilter.json
func ExampleSQLPoolsClient_NewListByWorkspacePager_listSqlAnalyticsPoolsInAWorkspaceWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLPoolsClient().NewListByWorkspacePager("sqlcrudtest-6845", "sqlcrudtest-7177", nil)
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
		// page.SQLPoolInfoListResult = armsynapse.SQLPoolInfoListResult{
		// 	Value: []*armsynapse.SQLPool{
		// 		{
		// 			Name: to.Ptr("master"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6845/providers/Microsoft.Synapse/workspaces/sqlcrudtest-7177/sqlPools/master"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsynapse.SQLPoolResourceProperties{
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T00:59:33.483Z"); return t}()),
		// 				MaxSizeBytes: to.Ptr[int64](32212254720),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				Status: to.Ptr("Online"),
		// 				StorageAccountType: to.Ptr(armsynapse.StorageAccountTypeGRS),
		// 			},
		// 			SKU: &armsynapse.SKU{
		// 				Name: to.Ptr("GP_Gen5_2"),
		// 			},
		// 	}},
		// }
	}
}
