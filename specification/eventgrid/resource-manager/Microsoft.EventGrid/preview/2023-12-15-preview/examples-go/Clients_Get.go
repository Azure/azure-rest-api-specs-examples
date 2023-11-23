package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Clients_Get.json
func ExampleClientsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClientsClient().Get(ctx, "examplerg", "exampleNamespaceName1", "exampleClientName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Client = armeventgrid.Client{
	// 	Name: to.Ptr("exampleClientName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/clients"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/clients/exampleClientName1"),
	// 	Properties: &armeventgrid.ClientProperties{
	// 		Description: to.Ptr("This is a test client"),
	// 		Attributes: map[string]any{
	// 			"deviceTypes": []any{
	// 				"Light",
	// 				"1",
	// 			},
	// 			"floor": float64(3),
	// 			"room": "345a",
	// 		},
	// 		ClientCertificateAuthentication: &armeventgrid.ClientCertificateAuthentication{
	// 			ValidationScheme: to.Ptr(armeventgrid.ClientCertificateValidationSchemeSubjectMatchesAuthenticationName),
	// 		},
	// 		ProvisioningState: to.Ptr(armeventgrid.ClientProvisioningStateSucceeded),
	// 		State: to.Ptr(armeventgrid.ClientStateEnabled),
	// 	},
	// }
}
