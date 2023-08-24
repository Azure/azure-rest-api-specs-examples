package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPerformConnectivityCheck.json
func ExampleClient_BeginPerformConnectivityCheckAsync_tcpConnectivityCheck() {
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
			Address: to.Ptr("8.8.8.8"),
			Port:    to.Ptr[int64](53),
		},
		PreferredIPVersion: to.Ptr(armapimanagement.PreferredIPVersionIPv4),
		Source: &armapimanagement.ConnectivityCheckRequestSource{
			Region: to.Ptr("northeurope"),
		},
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
	// 	AvgLatencyInMs: to.Ptr[int64](1),
	// 	ConnectionStatus: to.Ptr(armapimanagement.ConnectionStatusConnected),
	// 	Hops: []*armapimanagement.ConnectivityHop{
	// 		{
	// 			Type: to.Ptr("Source"),
	// 			Address: to.Ptr("10.1.1.4"),
	// 			ID: to.Ptr("7dbbe7aa-60ba-4650-831e-63d775d38e9e"),
	// 			Issues: []*armapimanagement.ConnectivityIssue{
	// 			},
	// 			NextHopIDs: []*string{
	// 				to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9")},
	// 				ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 			},
	// 			{
	// 				Type: to.Ptr("Internet"),
	// 				Address: to.Ptr("8.8.8.8"),
	// 				ID: to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9"),
	// 				Issues: []*armapimanagement.ConnectivityIssue{
	// 				},
	// 				NextHopIDs: []*string{
	// 				},
	// 		}},
	// 		MaxLatencyInMs: to.Ptr[int64](4),
	// 		MinLatencyInMs: to.Ptr[int64](1),
	// 		ProbesFailed: to.Ptr[int64](0),
	// 		ProbesSent: to.Ptr[int64](100),
	// 	}
}
