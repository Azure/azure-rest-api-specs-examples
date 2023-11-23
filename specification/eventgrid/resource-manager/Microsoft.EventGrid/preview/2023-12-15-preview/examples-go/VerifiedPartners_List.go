package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/VerifiedPartners_List.json
func ExampleVerifiedPartnersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVerifiedPartnersClient().NewListPager(&armeventgrid.VerifiedPartnersClientListOptions{Filter: nil,
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VerifiedPartnersListResult = armeventgrid.VerifiedPartnersListResult{
		// 	Value: []*armeventgrid.VerifiedPartner{
		// 		{
		// 			Name: to.Ptr("Contoso.Finance"),
		// 			Type: to.Ptr("Microsoft.EventGrid/verifiedPartners"),
		// 			ID: to.Ptr("/providers/Microsoft.EventGrid/verifiedPartners/Contoso.Finance"),
		// 			Properties: &armeventgrid.VerifiedPartnerProperties{
		// 				OrganizationName: to.Ptr("Contoso"),
		// 				PartnerDestinationDetails: &armeventgrid.PartnerDetails{
		// 					Description: to.Ptr("This is custom description"),
		// 					LongDescription: to.Ptr("This is long custom description"),
		// 					SetupURI: to.Ptr("https://www.example.com/"),
		// 				},
		// 				PartnerDisplayName: to.Ptr("Contoso - Finance Department"),
		// 				PartnerRegistrationImmutableID: to.Ptr("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
		// 				PartnerTopicDetails: &armeventgrid.PartnerDetails{
		// 					Description: to.Ptr("This is short description"),
		// 					LongDescription: to.Ptr("This is really long long... long description"),
		// 					SetupURI: to.Ptr("https://www.example.com/"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
