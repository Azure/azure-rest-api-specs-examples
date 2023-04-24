package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_testallroutes.json
func ExampleResourceClient_TestAllRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().TestAllRoutes(ctx, "testHub", "myResourceGroup", armiothub.TestAllRoutesInput{
		Message: &armiothub.RoutingMessage{
			AppProperties: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Body: to.Ptr("Body of message"),
			SystemProperties: map[string]*string{
				"key1": to.Ptr("value1"),
			},
		},
		RoutingSource: to.Ptr(armiothub.RoutingSourceDeviceMessages),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestAllRoutesResult = armiothub.TestAllRoutesResult{
	// 	Routes: []*armiothub.MatchedRoute{
	// 		{
	// 			Properties: &armiothub.RouteProperties{
	// 				Name: to.Ptr("Routeid"),
	// 				EndpointNames: []*string{
	// 					to.Ptr("id1")},
	// 					IsEnabled: to.Ptr(true),
	// 					Source: to.Ptr(armiothub.RoutingSourceDeviceMessages),
	// 				},
	// 		}},
	// 	}
}
