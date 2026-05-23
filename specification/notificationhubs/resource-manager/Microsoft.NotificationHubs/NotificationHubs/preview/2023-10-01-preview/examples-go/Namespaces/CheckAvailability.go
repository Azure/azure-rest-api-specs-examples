package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: 2023-10-01-preview/Namespaces/CheckAvailability.json
func ExampleNamespacesClient_CheckAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().CheckAvailability(ctx, armnotificationhubs.CheckAvailabilityParameters{
		Name: to.Ptr("sdk-Namespace-2924"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnotificationhubs.NamespacesClientCheckAvailabilityResponse{
	// 	CheckAvailabilityResult: &armnotificationhubs.CheckAvailabilityResult{
	// 		Name: to.Ptr("sdk-Namespace-2924"),
	// 		Type: to.Ptr("Microsoft.NotificationHubs/namespaces/checkNamespaceAvailability"),
	// 		ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/providers/Microsoft.NotificationHubs/namespaces/sdk-Namespace-2924"),
	// 		IsAvailiable: to.Ptr(true),
	// 	},
	// }
}
