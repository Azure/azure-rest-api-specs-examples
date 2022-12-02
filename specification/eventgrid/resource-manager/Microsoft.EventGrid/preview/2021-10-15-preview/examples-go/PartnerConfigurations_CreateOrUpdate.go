package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerConfigurations_CreateOrUpdate.json
func ExamplePartnerConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerConfigurationsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "examplerg", armeventgrid.PartnerConfiguration{
		Properties: &armeventgrid.PartnerConfigurationProperties{
			PartnerAuthorization: &armeventgrid.PartnerAuthorization{
				AuthorizedPartnersList: []*armeventgrid.Partner{
					{
						AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-28T01:20:55.142Z"); return t }()),
						PartnerName:                      to.Ptr("Contoso.Finance"),
						PartnerRegistrationImmutableID:   to.Ptr("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
					},
					{
						AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-20T01:00:00.142Z"); return t }()),
						PartnerName:                      to.Ptr("fabrikam.HR"),
						PartnerRegistrationImmutableID:   to.Ptr("5362bdb6-ce3e-4d0d-9a5b-3eb92c8aab38"),
					}},
				DefaultMaximumExpirationTimeInDays: to.Ptr[int32](10),
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
	// TODO: use response item
	_ = res
}
