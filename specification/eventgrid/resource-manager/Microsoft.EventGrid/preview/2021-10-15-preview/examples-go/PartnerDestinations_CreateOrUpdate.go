package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerDestinations_CreateOrUpdate.json
func ExamplePartnerDestinationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerDestinationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<partner-destination-name>",
		armeventgrid.PartnerDestination{
			Location: to.Ptr("<location>"),
			Properties: &armeventgrid.PartnerDestinationProperties{
				EndpointBaseURL:                 to.Ptr("<endpoint-base-url>"),
				EndpointServiceContext:          to.Ptr("<endpoint-service-context>"),
				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-14T19:33:43.430Z"); return t }()),
				MessageForActivation:            to.Ptr("<message-for-activation>"),
				PartnerRegistrationImmutableID:  to.Ptr("<partner-registration-immutable-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
