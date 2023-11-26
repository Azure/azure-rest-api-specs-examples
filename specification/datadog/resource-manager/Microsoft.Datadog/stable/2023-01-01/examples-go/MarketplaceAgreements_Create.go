package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MarketplaceAgreements_Create.json
func ExampleMarketplaceAgreementsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceAgreementsClient().CreateOrUpdate(ctx, &armdatadog.MarketplaceAgreementsClientCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgreementResource = armdatadog.AgreementResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Datadog/agreements"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Datadog/agreements/default"),
	// 	Properties: &armdatadog.AgreementProperties{
	// 		Accepted: to.Ptr(true),
	// 		LicenseTextLink: to.Ptr("test.licenseLink1"),
	// 		Plan: to.Ptr("planid1"),
	// 		PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink1"),
	// 		Product: to.Ptr("offid1"),
	// 		Publisher: to.Ptr("pubid1"),
	// 		RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.121Z"); return t}()),
	// 		Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
	// 	},
	// }
}
