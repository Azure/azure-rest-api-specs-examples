package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apps_GetResourceUploadUrl.json
func ExampleAppsClient_GetResourceUploadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().GetResourceUploadURL(ctx, "myResourceGroup", "myservice", "myapp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceUploadDefinition = armappplatform.ResourceUploadDefinition{
	// 	RelativePath: to.Ptr("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-20190801-3ed9f4a2-986b-4bbd-b833-a42dccb2f777"),
	// 	UploadURL: to.Ptr("https://springcloudstorageaccount.file.core.windows.net/bd172614181f42e2853f6fd90029cda8/e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-20190801-3ed9f4a2-986b-4bbd-b833-a42dccb2f777?sv=2018-03-28&sr=f&sig=SampleSignature&se=2019-08-01T10%3A42%3A21Z&sp=w"),
	// }
}
