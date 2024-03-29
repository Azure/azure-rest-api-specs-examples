package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlpoolList.json
func ExampleRestorableDroppedSQLPoolsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableDroppedSQLPoolsClient().NewListByWorkspacePager("restorabledroppeddatabasetest-1349", "restorabledroppeddatabasetest-1840", nil)
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
		// page.RestorableDroppedSQLPoolListResult = armsynapse.RestorableDroppedSQLPoolListResult{
		// 	Value: []*armsynapse.RestorableDroppedSQLPool{
		// 		{
		// 			Name: to.Ptr("restorabledroppeddatabasetest-1345,131403269876900000"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/restorableDroppedSqlPools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/restorabledroppeddatabasetest-1349/providers/Microsoft.Synapse/workspaces/restorabledroppeddatabasetest-1840/restorableDroppedSqlPools/restorabledroppeddatabasetest-1345,131403269876900000"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsynapse.RestorableDroppedSQLPoolProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T00:56:19.200Z"); return t}()),
		// 				DatabaseName: to.Ptr("restorabledroppeddatabasetest-1345"),
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T02:49:47.690Z"); return t}()),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-20T02:49:47.690Z"); return t}()),
		// 				Edition: to.Ptr("Basic"),
		// 				MaxSizeBytes: to.Ptr("2147483648"),
		// 				ServiceLevelObjective: to.Ptr("Basic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("restorabledroppeddatabasetest-2764,131403249684030000"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/restorableDroppedSqlPools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/restorabledroppeddatabasetest-1349/providers/Microsoft.Synapse/workspaces/restorabledroppeddatabasetest-1840/restorableDroppedSqlPools/restorabledroppeddatabasetest-2764,131403249684030000"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsynapse.RestorableDroppedSQLPoolProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-10T00:56:19.200Z"); return t}()),
		// 				DatabaseName: to.Ptr("restorabledroppeddatabasetest-2764"),
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T02:16:08.403Z"); return t}()),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-21T02:49:47.690Z"); return t}()),
		// 				Edition: to.Ptr("Standard"),
		// 				MaxSizeBytes: to.Ptr("268435456000"),
		// 				ServiceLevelObjective: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}
