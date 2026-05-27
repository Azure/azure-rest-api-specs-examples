package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/GetDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("a3473687-7581-41e1-ac24-6bcca5843f07", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseExtensionsClient().Get(ctx, "rg_a1f9d6f8-30d5-4228-9504-8a364361bca3", "srv_65858e0f-b1d1-4bdc-8351-a7da86ca4939", "11aa6c5e-58ed-4693-b303-3b8e3131deaa", "polybaseimport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.DatabaseExtensionsClientGetResponse{
	// }
}
