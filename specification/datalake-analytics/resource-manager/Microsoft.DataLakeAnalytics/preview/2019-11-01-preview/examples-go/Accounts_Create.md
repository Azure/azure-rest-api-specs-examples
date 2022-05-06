Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatalake-analytics%2Farmdatalakeanalytics%2Fv0.5.0/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatalakeanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armdatalakeanalytics.CreateDataLakeAnalyticsAccountParameters{
			Location: to.Ptr("<location>"),
			Properties: &armdatalakeanalytics.CreateDataLakeAnalyticsAccountProperties{
				ComputePolicies: []*armdatalakeanalytics.CreateComputePolicyWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakeanalytics.CreateOrUpdateComputePolicyProperties{
							MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
							MinPriorityPerJob:            to.Ptr[int32](1),
							ObjectID:                     to.Ptr("<object-id>"),
							ObjectType:                   to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
						},
					}},
				DataLakeStoreAccounts: []*armdatalakeanalytics.AddDataLakeStoreWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakeanalytics.AddDataLakeStoreProperties{
							Suffix: to.Ptr("<suffix>"),
						},
					}},
				DefaultDataLakeStoreAccount: to.Ptr("<default-data-lake-store-account>"),
				FirewallAllowAzureIPs:       to.Ptr(armdatalakeanalytics.FirewallAllowAzureIPsStateEnabled),
				FirewallRules: []*armdatalakeanalytics.CreateFirewallRuleWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakeanalytics.CreateOrUpdateFirewallRuleProperties{
							EndIPAddress:   to.Ptr("<end-ipaddress>"),
							StartIPAddress: to.Ptr("<start-ipaddress>"),
						},
					}},
				FirewallState:                to.Ptr(armdatalakeanalytics.FirewallStateEnabled),
				MaxDegreeOfParallelism:       to.Ptr[int32](30),
				MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
				MaxJobCount:                  to.Ptr[int32](3),
				MinPriorityPerJob:            to.Ptr[int32](1),
				NewTier:                      to.Ptr(armdatalakeanalytics.TierTypeConsumption),
				QueryStoreRetention:          to.Ptr[int32](30),
				StorageAccounts: []*armdatalakeanalytics.AddStorageAccountWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakeanalytics.AddStorageAccountProperties{
							AccessKey: to.Ptr("<access-key>"),
							Suffix:    to.Ptr("<suffix>"),
						},
					}},
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
			},
		},
		&armdatalakeanalytics.AccountsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
