package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ApproveRejectSitePrivateEndpointConnection.json
func ExampleStaticSitesClient_BeginApproveOrRejectPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStaticSitesClient().BeginApproveOrRejectPrivateEndpointConnection(ctx, "rg", "testSite", "connection", armappservice.RemotePrivateEndpointConnectionARMResource{
		Properties: &armappservice.RemotePrivateEndpointConnectionARMResourceProperties{
			PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
				Description:     to.Ptr("Approved by admin."),
				ActionsRequired: to.Ptr(""),
				Status:          to.Ptr("Approved"),
			},
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
	// res.RemotePrivateEndpointConnectionARMResource = armappservice.RemotePrivateEndpointConnectionARMResource{
	// 	Name: to.Ptr("connection"),
	// 	Type: to.Ptr("Microsoft.Web/sites/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/sites/testSite/privateEndpointConnections/connection"),
	// 	Properties: &armappservice.RemotePrivateEndpointConnectionARMResourceProperties{
	// 		PrivateEndpoint: &armappservice.ArmIDWrapper{
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
	// 			Description: to.Ptr("Approved by admin."),
	// 			ActionsRequired: to.Ptr(""),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
