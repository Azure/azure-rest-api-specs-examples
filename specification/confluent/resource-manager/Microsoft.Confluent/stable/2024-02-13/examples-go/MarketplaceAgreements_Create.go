package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/MarketplaceAgreements_Create.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceAgreementsClient().Create(ctx, &armconfluent.MarketplaceAgreementsClientCreateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgreementResource = armconfluent.AgreementResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Confluent/agreements"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Confluent/agreements/default"),
	// 	Properties: &armconfluent.AgreementProperties{
	// 		Accepted: to.Ptr(true),
	// 		LicenseTextLink: to.Ptr("test.licenseLink1"),
	// 		Plan: to.Ptr("planid1"),
	// 		PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink1"),
	// 		Product: to.Ptr("offid1"),
	// 		Publisher: to.Ptr("pubid1"),
	// 		RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-05T17:33:07.121Z"); return t}()),
	// 		Signature: to.Ptr("YKWOQOKH2BCKZ46O7SCKHANWEENRFRU5WB4LXDFUYWCBWTS4AG4SGQXCOZYIR5ZJCZTXRMZKYZMO2BJSL5YKPLAR4LBFRUNS6CRYE7A"),
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// }
}
