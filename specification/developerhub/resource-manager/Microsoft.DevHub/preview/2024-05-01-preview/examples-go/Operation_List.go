package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/34f5146bc945549d087d38a8a593c0a5f475ad7f/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2024-05-01-preview/examples/Operation_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armdevhub.OperationListResult{
	// 	Value: []*armdevhub.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.DevHub/workflows/read"),
	// 			Display: &armdevhub.OperationDisplay{
	// 				Description: to.Ptr("Gets workflow"),
	// 				Operation: to.Ptr("Get workflow"),
	// 				Provider: to.Ptr("Microsoft Developer Hub"),
	// 				Resource: to.Ptr("Workflows"),
	// 			},
	// 	}},
	// }
}
