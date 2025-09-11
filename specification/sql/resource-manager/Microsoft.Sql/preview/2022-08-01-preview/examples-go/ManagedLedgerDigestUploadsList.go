package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedLedgerDigestUploadsList.json
func ExampleManagedLedgerDigestUploadsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedLedgerDigestUploadsClient().NewListByDatabasePager("ledgertestrg", "ledgertestserver", "testdb", nil)
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
		// page.ManagedLedgerDigestUploadsListResult = armsql.ManagedLedgerDigestUploadsListResult{
		// 	Value: []*armsql.ManagedLedgerDigestUploads{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/ledgerDigestUploads"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/ledgertestrg/providers/Microsoft.Sql/managedInstances/ledgertestserver/databases/testdb/current"),
		// 			Properties: &armsql.ManagedLedgerDigestUploadsProperties{
		// 				DigestStorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
		// 				State: to.Ptr(armsql.ManagedLedgerDigestUploadsStateEnabled),
		// 			},
		// 	}},
		// }
	}
}
