package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/DeleteSqlServer.json
func ExampleSQLServersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazuredata.NewSQLServersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"testrg",
		"testsqlregistration",
		"testsqlserver",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
