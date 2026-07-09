package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/NetworkConnections_ListOutboundNetworkDependenciesEndpoints.json
func ExampleNetworkConnectionsClient_NewListOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkConnectionsClient().NewListOutboundNetworkDependenciesEndpointsPager("rg1", "uswest3network", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse{
		// 	OutboundEnvironmentEndpointCollection: armdevcenter.OutboundEnvironmentEndpointCollection{
		// 		Value: []*armdevcenter.OutboundEnvironmentEndpoint{
		// 			{
		// 				Category: to.Ptr("Dev Box Service"),
		// 				Endpoints: []*armdevcenter.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("devbox.azure.com"),
		// 						EndpointDetails: []*armdevcenter.EndpointDetail{
		// 							{
		// 								Port: to.Ptr[int32](443),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Category: to.Ptr("Intune"),
		// 				Endpoints: []*armdevcenter.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("login.microsoftonline.com"),
		// 						EndpointDetails: []*armdevcenter.EndpointDetail{
		// 							{
		// 								Port: to.Ptr[int32](443),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Category: to.Ptr("Cloud PC"),
		// 				Endpoints: []*armdevcenter.EndpointDependency{
		// 					{
		// 						DomainName: to.Ptr("rdweb.wvd.microsoft.com"),
		// 						EndpointDetails: []*armdevcenter.EndpointDetail{
		// 							{
		// 								Port: to.Ptr[int32](443),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
