package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub/v2"
)

// Generated from example definition: 2026-03-01-preview/iothub_testnewroute.json
func ExampleResourceClient_TestRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().TestRoute(ctx, "testHub", "myResourceGroup", armiothub.TestRouteInput{
		Message: &armiothub.RoutingMessage{
			AppProperties: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Body: to.Ptr("Body of message"),
			SystemProperties: map[string]*string{
				"key1": to.Ptr("value1"),
			},
		},
		Route: &armiothub.RouteProperties{
			Name: to.Ptr("Routeid"),
			EndpointNames: []*string{
				to.Ptr("id1"),
			},
			IsEnabled: to.Ptr(true),
			Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiothub.ResourceClientTestRouteResponse{
	// 	TestRouteResult: armiothub.TestRouteResult{
	// 		Result: to.Ptr(armiothub.TestResultStatusFalse),
	// 		Details: &armiothub.TestRouteResultDetails{
	// 			CompilationErrors: []*armiothub.RouteCompilationError{
	// 				{
	// 					Location: &armiothub.RouteErrorRange{
	// 						End: &armiothub.RouteErrorPosition{
	// 							Column: to.Ptr[int32](24),
	// 							Line: to.Ptr[int32](12),
	// 						},
	// 						Start: &armiothub.RouteErrorPosition{
	// 							Column: to.Ptr[int32](12),
	// 							Line: to.Ptr[int32](12),
	// 						},
	// 					},
	// 					Message: to.Ptr("string response"),
	// 					Severity: to.Ptr(armiothub.RouteErrorSeverityError),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
