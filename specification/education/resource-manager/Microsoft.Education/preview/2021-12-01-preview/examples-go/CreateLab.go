package armeducation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/CreateLab.json
func ExampleLabsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeducation.NewLabsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		"{invoiceSectionName}",
		armeducation.LabDetails{
			Properties: &armeducation.LabProperties{
				Description: to.Ptr("example lab description"),
				BudgetPerStudent: &armeducation.Amount{
					Currency: to.Ptr("USD"),
					Value:    to.Ptr[float32](100),
				},
				DisplayName:    to.Ptr("example lab"),
				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-09T22:11:29.422Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
