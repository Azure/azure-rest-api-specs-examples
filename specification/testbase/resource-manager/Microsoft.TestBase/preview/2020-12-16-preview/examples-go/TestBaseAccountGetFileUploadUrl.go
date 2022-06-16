package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountGetFileUploadUrl.json
func ExampleTestBaseAccountsClient_GetFileUploadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestBaseAccountsClient("<subscription-id>", cred, nil)
	_, err = client.GetFileUploadURL(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		&armtestbase.TestBaseAccountsGetFileUploadURLOptions{Parameters: &armtestbase.GetFileUploadURLParameters{
			BlobName: to.StringPtr("<blob-name>"),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
