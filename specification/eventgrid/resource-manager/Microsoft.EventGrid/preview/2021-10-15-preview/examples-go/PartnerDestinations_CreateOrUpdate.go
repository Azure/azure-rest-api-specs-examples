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
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerDestinationsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examplerg", "examplePartnerDestinationName1", armeventgrid.PartnerDestination{
		Location: to.Ptr("westus2"),
		Properties: &armeventgrid.PartnerDestinationProperties{
			EndpointBaseURL:                 to.Ptr("https://www.example/endpoint"),
			EndpointServiceContext:          to.Ptr("This is an example"),
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-14T19:33:43.430Z"); return t }()),
			MessageForActivation:            to.Ptr("Sample Activation message"),
			PartnerRegistrationImmutableID:  to.Ptr("0bd70ee2-7d95-447e-ab1f-c4f320019404"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
