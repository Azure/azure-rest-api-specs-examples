package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/listKeys.json
func ExampleWorkspaceClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquantum.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceClient().ListKeys(ctx, "quantumResourcegroup", "quantumworkspace1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListKeysResult = armquantum.ListKeysResult{
	// 	APIKeyEnabled: to.Ptr(true),
	// 	PrimaryConnectionString: to.Ptr("<primaryConnectionString>"),
	// 	PrimaryKey: &armquantum.APIKey{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-13T23:06:09.147Z"); return t}()),
	// 		Key: to.Ptr("<primaryKey>"),
	// 	},
	// 	SecondaryConnectionString: to.Ptr("<secondaryConnectionString>"),
	// 	SecondaryKey: &armquantum.APIKey{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-13T23:06:09.147Z"); return t}()),
	// 		Key: to.Ptr("<secondaryKey>"),
	// 	},
	// }
}
