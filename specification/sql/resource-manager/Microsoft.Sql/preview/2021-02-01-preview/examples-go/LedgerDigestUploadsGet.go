package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/LedgerDigestUploadsGet.json
func ExampleLedgerDigestUploadsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLedgerDigestUploadsClient().Get(ctx, "ledgertestrg", "ledgertestserver", "testdb", armsql.LedgerDigestUploadsNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LedgerDigestUploads = armsql.LedgerDigestUploads{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/ledgerDigestUploads"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/ledgertestrg/providers/Microsoft.Sql/servers/ledgertestserver/databases/testdb/current"),
	// 	Properties: &armsql.LedgerDigestUploadsProperties{
	// 		DigestStorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		State: to.Ptr(armsql.LedgerDigestUploadsStateEnabled),
	// 	},
	// }
}
