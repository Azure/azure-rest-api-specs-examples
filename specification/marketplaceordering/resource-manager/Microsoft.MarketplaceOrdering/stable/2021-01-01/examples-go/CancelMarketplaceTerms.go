package armmarketplaceordering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/CancelMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplaceordering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceAgreementsClient().Cancel(ctx, "pubid", "offid", "planid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgreementTerms = armmarketplaceordering.AgreementTerms{
	// 	Name: to.Ptr("planid"),
	// 	Type: to.Ptr("Microsoft.MarketplaceOrdering/offertypes"),
	// 	ID: to.Ptr("id"),
	// 	Properties: &armmarketplaceordering.AgreementProperties{
	// 		Accepted: to.Ptr(false),
	// 		LicenseTextLink: to.Ptr("test.licenseLink"),
	// 		MarketplaceTermsLink: to.Ptr("test.marketplaceTermsLink"),
	// 		Plan: to.Ptr("planid"),
	// 		PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink"),
	// 		Product: to.Ptr("offid"),
	// 		Publisher: to.Ptr("pubid"),
	// 		RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.121Z"); return t}()),
	// 		Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
	// 	},
	// }
}
