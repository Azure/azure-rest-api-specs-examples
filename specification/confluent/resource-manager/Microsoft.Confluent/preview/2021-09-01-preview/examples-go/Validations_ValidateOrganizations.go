package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Validations_ValidateOrganizations.json
func ExampleValidationsClient_ValidateOrganization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfluent.NewValidationsClient("<subscription-id>", cred, nil)
	res, err := client.ValidateOrganization(ctx,
		"<resource-group-name>",
		"<organization-name>",
		armconfluent.OrganizationResource{
			Location: to.StringPtr("<location>"),
			Properties: &armconfluent.OrganizationResourceProperties{
				OfferDetail: &armconfluent.OfferDetail{
					ID:          to.StringPtr("<id>"),
					PlanID:      to.StringPtr("<plan-id>"),
					PlanName:    to.StringPtr("<plan-name>"),
					PublisherID: to.StringPtr("<publisher-id>"),
					TermUnit:    to.StringPtr("<term-unit>"),
				},
				UserDetail: &armconfluent.UserDetail{
					EmailAddress: to.StringPtr("<email-address>"),
					FirstName:    to.StringPtr("<first-name>"),
					LastName:     to.StringPtr("<last-name>"),
				},
			},
			Tags: map[string]*string{
				"Environment": to.StringPtr("Dev"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ValidationsClientValidateOrganizationResult)
}
