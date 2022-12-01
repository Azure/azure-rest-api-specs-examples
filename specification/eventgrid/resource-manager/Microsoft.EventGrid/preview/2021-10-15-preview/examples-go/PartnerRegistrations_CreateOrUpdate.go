package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerRegistrations_CreateOrUpdate.json
func ExamplePartnerRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerRegistrationsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "examplerg", "examplePartnerRegistrationName1", armeventgrid.PartnerRegistration{
		Location: to.Ptr("global"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("Value2"),
			"key3": to.Ptr("Value3"),
		},
		Properties: &armeventgrid.PartnerRegistrationProperties{
			AuthorizedAzureSubscriptionIDs: []*string{
				to.Ptr("d48566a8-2428-4a6c-8347-9675d09fb851")},
			LogoURI:                        to.Ptr("https://www.example.com/logo.png"),
			PartnerName:                    to.Ptr("ContosoCorp"),
			PartnerResourceTypeDescription: to.Ptr("ContocoCorp Accounts Description Text"),
			PartnerResourceTypeDisplayName: to.Ptr("ContocoCorp Accounts DisplayName Text"),
			PartnerResourceTypeName:        to.Ptr("ContosoCorp.Accounts"),
			SetupURI:                       to.Ptr("https://www.example.com/setup.html"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
