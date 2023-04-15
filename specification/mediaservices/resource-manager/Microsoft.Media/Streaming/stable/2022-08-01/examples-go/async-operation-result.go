package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/async-operation-result.json
func ExampleStreamingEndpointsClient_AsyncOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingEndpointsClient().AsyncOperation(ctx, "mediaresources", "slitestmedia10", "62e4d893-d233-4005-988e-a428d9f77076", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AsyncOperationResult = armmediaservices.AsyncOperationResult{
	// 	Name: to.Ptr("62e4d893-d233-4005-988e-a428d9f77076"),
	// 	Error: &armmediaservices.ErrorDetail{
	// 		Code: to.Ptr("None"),
	// 		Target: to.Ptr("d7aea790-8acb-40b9-8f9f-21cc37c9e519"),
	// 	},
	// 	Status: to.Ptr(armmediaservices.AsyncOperationStatusInProgress),
	// }
}
