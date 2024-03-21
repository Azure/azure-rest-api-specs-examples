package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/MarketplaceAgreements_List.json
func ExampleMarketplaceAgreementsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplaceAgreementsClient().NewListPager(nil)
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
		// page.AgreementResourceListResponse = armconfluent.AgreementResourceListResponse{
		// 	Value: []*armconfluent.AgreementResource{
		// 		{
		// 			Name: to.Ptr("planid1"),
		// 			Type: to.Ptr("Microsoft.Confluent/agreements"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Confluent/agreements/default"),
		// 			Properties: &armconfluent.AgreementProperties{
		// 				Accepted: to.Ptr(true),
		// 				LicenseTextLink: to.Ptr("test.licenseLink1"),
		// 				Plan: to.Ptr("planid1"),
		// 				PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink1"),
		// 				Product: to.Ptr("offid1"),
		// 				Publisher: to.Ptr("pubid1"),
		// 				RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.121Z"); return t}()),
		// 				Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("planid2"),
		// 			Type: to.Ptr("Microsoft.MarketplaceOrdering/offertypes"),
		// 			ID: to.Ptr("id2"),
		// 			Properties: &armconfluent.AgreementProperties{
		// 				Accepted: to.Ptr(true),
		// 				LicenseTextLink: to.Ptr("test.licenseLin2k"),
		// 				Plan: to.Ptr("planid2"),
		// 				PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink2"),
		// 				Product: to.Ptr("offid2"),
		// 				Publisher: to.Ptr("pubid2"),
		// 				RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-14T11:33:07.121Z"); return t}()),
		// 				Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
		// 			},
		// 	}},
		// }
	}
}
