package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_ApproveOrRejectPrivateEndpointConnection.json
func ExampleEnvironmentsClient_BeginApproveOrRejectPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEnvironmentsClient().BeginApproveOrRejectPrivateEndpointConnection(ctx, "test-rg", "test-ase", "fa38656c-034e-43d8-adce-fe06ce039c98", armappservice.RemotePrivateEndpointConnectionARMResource{
		Properties: &armappservice.RemotePrivateEndpointConnectionARMResourceProperties{
			PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
				Description: to.Ptr("Approved by johndoe@company.com"),
				Status:      to.Ptr("Approved"),
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
	// 	Name: to.Ptr("fa38656c-034e-43d8-adce-fe06ce039c98"),
	// 	Type: to.Ptr("Microsoft.Web/hostingEnvironments/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/privateEndpointConnections/fa38656c-034e-43d8-adce-fe06ce039c98"),
	// 	Properties: &armappservice.RemotePrivateEndpointConnectionARMResourceProperties{
	// 		IPAddresses: []*string{
	// 		},
	// 		PrivateEndpoint: &armappservice.ArmIDWrapper{
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/test-privateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
	// 			Description: to.Ptr("Approved by johndoe@company.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
