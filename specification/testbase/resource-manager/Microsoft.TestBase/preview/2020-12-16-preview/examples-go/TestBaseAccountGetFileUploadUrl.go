package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountGetFileUploadUrl.json
func ExampleAccountsClient_GetFileUploadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().GetFileUploadURL(ctx, "contoso-rg1", "contoso-testBaseAccount1", &armtestbase.AccountsClientGetFileUploadURLOptions{Parameters: &armtestbase.GetFileUploadURLParameters{
		BlobName: to.Ptr("package.zip"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileUploadURLResponse = armtestbase.FileUploadURLResponse{
	// 	BlobPath: to.Ptr("https://uslwestusdevsa.blob.core.windows.net/usltest/temp/20c0d7e0-1bb6-477f-bc04-57c734453000/package.zip"),
	// 	UploadURL: to.Ptr("https://uslwestusdevsa.blob.core.windows.net/usltest/temp/20c0d7e0-1bb6-477f-bc04-57c734453000/token"),
	// }
}
