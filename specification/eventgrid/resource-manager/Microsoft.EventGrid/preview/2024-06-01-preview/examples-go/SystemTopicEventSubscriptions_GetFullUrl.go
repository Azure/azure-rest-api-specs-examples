package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/SystemTopicEventSubscriptions_GetFullUrl.json
func ExampleSystemTopicEventSubscriptionsClient_GetFullURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSystemTopicEventSubscriptionsClient().GetFullURL(ctx, "examplerg", "exampleSystemTopic1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscriptionFullURL = armeventgrid.EventSubscriptionFullURL{
	// 	EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
	// }
}
