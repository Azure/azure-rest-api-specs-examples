package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_GetLogFileUrl.json
func ExampleDeploymentsClient_GetLogFileURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().GetLogFileURL(ctx, "myResourceGroup", "myservice", "myapp", "mydeployment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogFileURLResponse = armappplatform.LogFileURLResponse{
	// 	URL: to.Ptr("https://spring.blob.core.windows.net/logs/110ec0c337154d45b1f01daf2196c0bf/b58b0cb4ecdea3c65311b4ca8833fe47b6ae0a7500f87a8eb31e8379d3fe48f1-2019081312-42b7b90c-f108-4c09-b33d-1ea134f57f23?sv=2018-03-28&sr=b&sig=example-signature&se=2019-08-14T09%3A43%3A52Z&sp=r"),
	// }
}
