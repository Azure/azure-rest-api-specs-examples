```go
package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_Create.json
func ExampleOrganizationClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconfluent.NewOrganizationClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"myOrganization",
		&armconfluent.OrganizationClientBeginCreateOptions{Body: &armconfluent.OrganizationResource{
			Location: to.Ptr("West US"),
			Properties: &armconfluent.OrganizationResourceProperties{
				OfferDetail: &armconfluent.OfferDetail{
					ID:          to.Ptr("string"),
					PlanID:      to.Ptr("string"),
					PlanName:    to.Ptr("string"),
					PublisherID: to.Ptr("string"),
					TermUnit:    to.Ptr("string"),
				},
				UserDetail: &armconfluent.UserDetail{
					EmailAddress: to.Ptr("contoso@microsoft.com"),
					FirstName:    to.Ptr("string"),
					LastName:     to.Ptr("string"),
				},
			},
			Tags: map[string]*string{
				"Environment": to.Ptr("Dev"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconfluent%2Farmconfluent%2Fv1.0.0/sdk/resourcemanager/confluent/armconfluent/README.md) on how to add the SDK to your project and authenticate.
