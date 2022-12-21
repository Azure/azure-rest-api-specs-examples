package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
func ExampleSavingsPlanClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbillingbenefits.NewSavingsPlanClient(nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(&armbillingbenefits.SavingsPlanClientListAllOptions{Filter: to.Ptr("(properties%2farchived+eq+false)"),
		Orderby:        to.Ptr("properties/displayName asc"),
		RefreshSummary: nil,
		Skiptoken:      to.Ptr[float32](50),
		SelectedState:  nil,
		Take:           to.Ptr[float32](1),
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
