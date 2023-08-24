package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPerformConnectivityCheckHttpConnect.json
func ExampleClient_BeginPerformConnectivityCheckAsync_httpConnectivityCheck() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginPerformConnectivityCheckAsync(ctx, "rg1", "apimService1", armapimanagement.ConnectivityCheckRequest{
		Destination: &armapimanagement.ConnectivityCheckRequestDestination{
			Address: to.Ptr("https://microsoft.com"),
			Port:    to.Ptr[int64](3306),
		},
		ProtocolConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfiguration{
			HTTPConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfigurationHTTPConfiguration{
				Method: to.Ptr(armapimanagement.MethodGET),
				Headers: []*armapimanagement.HTTPHeader{
					{
						Name:  to.Ptr("Authorization"),
						Value: to.Ptr("******"),
					}},
				ValidStatusCodes: []*int64{
					to.Ptr[int64](200),
					to.Ptr[int64](204)},
			},
		},
		Source: &armapimanagement.ConnectivityCheckRequestSource{
			Region: to.Ptr("northeurope"),
		},
		Protocol: to.Ptr(armapimanagement.ConnectivityCheckProtocolHTTPS),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectivityCheckResponse = armapimanagement.ConnectivityCheckResponse{
	// 	AvgLatencyInMs: to.Ptr[int64](260),
	// 	ConnectionStatus: to.Ptr(armapimanagement.ConnectionStatus("Reachable")),
	// 	Hops: []*armapimanagement.ConnectivityHop{
	// 		{
	// 			Type: to.Ptr("Source"),
	// 			Address: to.Ptr("20.82.216.48"),
	// 			ID: to.Ptr("c60e2296-5ebc-48cc-80e8-7e6d2981e7b2"),
	// 			Issues: []*armapimanagement.ConnectivityIssue{
	// 			},
	// 			NextHopIDs: []*string{
	// 				to.Ptr("26aa44e7-04f1-462f-aa5d-5951957b5650")},
	// 				ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 			},
	// 			{
	// 				Type: to.Ptr("Internet"),
	// 				Address: to.Ptr("40.113.200.201"),
	// 				ID: to.Ptr("26aa44e7-04f1-462f-aa5d-5951957b5650"),
	// 				Issues: []*armapimanagement.ConnectivityIssue{
	// 				},
	// 				NextHopIDs: []*string{
	// 				},
	// 		}},
	// 		MaxLatencyInMs: to.Ptr[int64](281),
	// 		MinLatencyInMs: to.Ptr[int64](250),
	// 		ProbesFailed: to.Ptr[int64](0),
	// 		ProbesSent: to.Ptr[int64](3),
	// 	}
}
