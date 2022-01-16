Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconfluent%2Farmconfluent%2Fv0.2.0/sdk/resourcemanager/confluent/armconfluent/README.md) on how to add the SDK to your project and authenticate.

```go
package armconfluent_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Create.json
func ExampleOrganizationClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfluent.NewOrganizationClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<organization-name>",
		&armconfluent.OrganizationClientBeginCreateOptions{Body: &armconfluent.OrganizationResource{
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
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OrganizationClientCreateResult)
}
```
