package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerConfigurations_UnauthorizePartner.json
func ExamplePartnerConfigurationsClient_UnauthorizePartner() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerConfigurationsClient().UnauthorizePartner(ctx, "examplerg", armeventgrid.Partner{
		AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-28T01:20:55.142Z"); return t }()),
		PartnerName:                      to.Ptr("Contoso.Finance"),
		PartnerRegistrationImmutableID:   to.Ptr("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerConfiguration = armeventgrid.PartnerConfiguration{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerConfigurations"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerConfigurations/default"),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armeventgrid.PartnerConfigurationProperties{
	// 		PartnerAuthorization: &armeventgrid.PartnerAuthorization{
	// 			AuthorizedPartnersList: []*armeventgrid.Partner{
	// 				{
	// 					AuthorizationExpirationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-20T01:00:00.142Z"); return t}()),
	// 					PartnerName: to.Ptr("fabrikam.HR"),
	// 					PartnerRegistrationImmutableID: to.Ptr("5362bdb6-ce3e-4d0d-9a5b-3eb92c8aab38"),
	// 			}},
	// 			DefaultMaximumExpirationTimeInDays: to.Ptr[int32](10),
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// }
}
