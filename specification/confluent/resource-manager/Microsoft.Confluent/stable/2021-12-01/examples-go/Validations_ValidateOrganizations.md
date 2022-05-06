Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconfluent%2Farmconfluent%2Fv0.4.0/sdk/resourcemanager/confluent/armconfluent/README.md) on how to add the SDK to your project and authenticate.

```go
package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Validations_ValidateOrganizations.json
func ExampleValidationsClient_ValidateOrganization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armconfluent.NewValidationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ValidateOrganization(ctx,
		"<resource-group-name>",
		"<organization-name>",
		armconfluent.OrganizationResource{
			Location: to.Ptr("<location>"),
			Properties: &armconfluent.OrganizationResourceProperties{
				OfferDetail: &armconfluent.OfferDetail{
					ID:          to.Ptr("<id>"),
					PlanID:      to.Ptr("<plan-id>"),
					PlanName:    to.Ptr("<plan-name>"),
					PublisherID: to.Ptr("<publisher-id>"),
					TermUnit:    to.Ptr("<term-unit>"),
				},
				UserDetail: &armconfluent.UserDetail{
					EmailAddress: to.Ptr("<email-address>"),
					FirstName:    to.Ptr("<first-name>"),
					LastName:     to.Ptr("<last-name>"),
				},
			},
			Tags: map[string]*string{
				"Environment": to.Ptr("Dev"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
