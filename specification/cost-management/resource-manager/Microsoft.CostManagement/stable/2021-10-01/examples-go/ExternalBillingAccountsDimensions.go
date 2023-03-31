package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountsDimensions.json
func ExampleDimensionsClient_NewByExternalCloudProviderTypePager_externalBillingAccountDimensionList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDimensionsClient().NewByExternalCloudProviderTypePager(armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts, "100", &armcostmanagement.DimensionsClientByExternalCloudProviderTypeOptions{Filter: nil,
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
		// 			Name: to.Ptr("dimensions_ResourceType_2019-12-01_2019-12-31"),
		// 			Type: to.Ptr("microsoft.consumption/dimensions"),
		// 			ID: to.Ptr("providers/Microsoft.CostManagement/externalBillingAccounts/100/dimensions_ResourceType_2019-12-01_2019-12-31"),
		// 			Properties: &armcostmanagement.DimensionProperties{
		// 				Description: to.Ptr("Resource type"),
		// 				Category: to.Ptr("ResourceType"),
		// 				Data: []*string{
		// 				},
		// 				FilterEnabled: to.Ptr(true),
		// 				GroupingEnabled: to.Ptr(true),
		// 				Total: to.Ptr[int32](0),
		// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-31T00:00:00Z"); return t}()),
		// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T00:00:00Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dimensions_ResourceId_2019-12-01_2019-12-31"),
		// 			Type: to.Ptr("microsoft.consumption/dimensions"),
		// 			ID: to.Ptr("providers/Microsoft.CostManagement/externalBillingAccounts/100/dimensions_ResourceId_2019-12-01_2019-12-31"),
		// 			Properties: &armcostmanagement.DimensionProperties{
		// 				Description: to.Ptr("Resource ID"),
		// 				Category: to.Ptr("ResourceId"),
		// 				Data: []*string{
		// 				},
		// 				FilterEnabled: to.Ptr(true),
		// 				GroupingEnabled: to.Ptr(true),
		// 				Total: to.Ptr[int32](0),
		// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-31T00:00:00Z"); return t}()),
		// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T00:00:00Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
