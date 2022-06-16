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
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewPartnerConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		armeventgrid.PartnerConfiguration{
			Properties: &armeventgrid.PartnerConfigurationProperties{
				PartnerAuthorization: &armeventgrid.PartnerAuthorization{
					AuthorizedPartnersList: []*armeventgrid.Partner{
						{
							AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-28T01:20:55.142Z"); return t }()),
							PartnerName:                      to.Ptr("<partner-name>"),
							PartnerRegistrationImmutableID:   to.Ptr("<partner-registration-immutable-id>"),
						},
						{
							AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-20T01:00:00.142Z"); return t }()),
							PartnerName:                      to.Ptr("<partner-name>"),
							PartnerRegistrationImmutableID:   to.Ptr("<partner-registration-immutable-id>"),
						}},
					DefaultMaximumExpirationTimeInDays: to.Ptr[int32](10),
				},
			},
		},
		&armeventgrid.PartnerConfigurationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
