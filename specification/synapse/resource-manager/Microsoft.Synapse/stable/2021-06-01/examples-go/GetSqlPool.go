package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPool.json
func ExampleSQLPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLPoolsClient().Get(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLPool = armsynapse.SQLPool{
	// 	Name: to.Ptr("sqlcrudtest-9187"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187"),
	// 	Location: to.Ptr("Japan East"),
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("TagValue1"),
	// 	},
	// 	Properties: &armsynapse.SQLPoolResourceProperties{
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T00:56:19.200Z"); return t}()),
	// 		MaxSizeBytes: to.Ptr[int64](268435456000),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		Status: to.Ptr("Online"),
	// 		StorageAccountType: to.Ptr(armsynapse.StorageAccountTypeGRS),
	// 	},
	// 	SKU: &armsynapse.SKU{
	// 		Name: to.Ptr("DW100c"),
	// 	},
	// }
}
