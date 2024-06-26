package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/RuleSets_ListResourceUsage.json
func ExampleRuleSetsClient_NewListResourceUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRuleSetsClient().NewListResourceUsagePager("RG", "profile1", "ruleSet1", nil)
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
		// page.UsagesListResult = armcdn.UsagesListResult{
		// 	Value: []*armcdn.Usage{
		// 		{
		// 			Name: &armcdn.UsageName{
		// 				LocalizedValue: to.Ptr("rule"),
		// 				Value: to.Ptr("rule"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/rulesets/ruleSet1/rules/rule1"),
		// 			Limit: to.Ptr[int64](25),
		// 			Unit: to.Ptr(armcdn.UsageUnitCount),
		// 	}},
		// }
	}
}
