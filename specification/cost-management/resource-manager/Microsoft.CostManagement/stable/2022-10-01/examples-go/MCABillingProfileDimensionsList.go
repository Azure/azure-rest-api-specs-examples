package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/MCABillingProfileDimensionsList.json
func ExampleDimensionsClient_NewListPager_billingProfileDimensionsListMca() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDimensionsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579", &armcostmanagement.DimensionsClientListOptions{Filter: nil,
		Expand:    nil,
		Skiptoken: nil,
		Top:       nil,
	})
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
		// page.DimensionsListResult = armcostmanagement.DimensionsListResult{
		// 	Value: []*armcostmanagement.Dimension{
		// 		{
		// 			Name: to.Ptr("dimensions_ResourceGroup_2019-10-01_2019-10-31"),
		// 			Type: to.Ptr("microsoft.CostManagement/dimensions"),
		// 			ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/providers/microsoft.CostManagement/dimensions_ResourceGroup_2019-10-01_2019-10-31"),
		// 			Properties: &armcostmanagement.DimensionProperties{
		// 				Description: to.Ptr("Resource group"),
		// 				Category: to.Ptr("ResourceGroup"),
		// 				Data: []*string{
		// 				},
		// 				FilterEnabled: to.Ptr(true),
		// 				GroupingEnabled: to.Ptr(true),
		// 				Total: to.Ptr[int32](377),
		// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-31T07:00:00.000Z"); return t}()),
		// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-01T07:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dimensions_ResourceType_2019-10-01_2019-10-31"),
		// 			Type: to.Ptr("microsoft.CostManagement/dimensions"),
		// 			ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/providers/microsoft.CostManagement/dimensions_ResourceType_2019-10-01_2019-10-31"),
		// 			Properties: &armcostmanagement.DimensionProperties{
		// 				Description: to.Ptr("Resource type"),
		// 				Category: to.Ptr("ResourceType"),
		// 				Data: []*string{
		// 				},
		// 				FilterEnabled: to.Ptr(true),
		// 				GroupingEnabled: to.Ptr(true),
		// 				Total: to.Ptr[int32](37),
		// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-31T07:00:00.000Z"); return t}()),
		// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-01T07:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
