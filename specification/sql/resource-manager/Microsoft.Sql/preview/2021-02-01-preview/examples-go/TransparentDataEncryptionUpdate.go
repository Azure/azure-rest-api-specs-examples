package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/TransparentDataEncryptionUpdate.json
func ExampleTransparentDataEncryptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransparentDataEncryptionsClient().CreateOrUpdate(ctx, "securitytde-42-rg", "securitytde-42", "testdb", armsql.TransparentDataEncryptionNameCurrent, armsql.LogicalDatabaseTransparentDataEncryption{
		Properties: &armsql.TransparentDataEncryptionProperties{
			State: to.Ptr(armsql.TransparentDataEncryptionStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogicalDatabaseTransparentDataEncryption = armsql.LogicalDatabaseTransparentDataEncryption{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/transparentDataEncryption"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securitytde-42-rg/providers/Microsoft.Sql/servers/securitytde-42/databases/testdb/transparentDataEncryption"),
	// 	Properties: &armsql.TransparentDataEncryptionProperties{
	// 		State: to.Ptr(armsql.TransparentDataEncryptionStateEnabled),
	// 	},
	// }
}
