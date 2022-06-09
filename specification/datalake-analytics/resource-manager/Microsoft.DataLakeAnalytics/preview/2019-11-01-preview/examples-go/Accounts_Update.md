```go
package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Update.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"contosorg",
		"contosoadla",
		&armdatalakeanalytics.AccountsClientBeginUpdateOptions{Parameters: &armdatalakeanalytics.UpdateDataLakeAnalyticsAccountParameters{
			Properties: &armdatalakeanalytics.UpdateDataLakeAnalyticsAccountProperties{
				ComputePolicies: []*armdatalakeanalytics.UpdateComputePolicyWithAccountParameters{
					{
						Name: to.Ptr("test_policy"),
						Properties: &armdatalakeanalytics.UpdateComputePolicyProperties{
							MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
							MinPriorityPerJob:            to.Ptr[int32](1),
							ObjectID:                     to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
							ObjectType:                   to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
						},
					}},
				FirewallAllowAzureIPs: to.Ptr(armdatalakeanalytics.FirewallAllowAzureIPsStateEnabled),
				FirewallRules: []*armdatalakeanalytics.UpdateFirewallRuleWithAccountParameters{
					{
						Name: to.Ptr("test_rule"),
						Properties: &armdatalakeanalytics.UpdateFirewallRuleProperties{
							EndIPAddress:   to.Ptr("2.2.2.2"),
							StartIPAddress: to.Ptr("1.1.1.1"),
						},
					}},
				FirewallState:                to.Ptr(armdatalakeanalytics.FirewallStateEnabled),
				MaxDegreeOfParallelism:       to.Ptr[int32](1),
				MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
				MaxJobCount:                  to.Ptr[int32](1),
				MinPriorityPerJob:            to.Ptr[int32](1),
				NewTier:                      to.Ptr(armdatalakeanalytics.TierTypeConsumption),
				QueryStoreRetention:          to.Ptr[int32](1),
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatalake-analytics%2Farmdatalakeanalytics%2Fv0.6.0/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics/README.md) on how to add the SDK to your project and authenticate.
