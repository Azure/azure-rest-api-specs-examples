package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contosorg",
		"contosoadla",
		armdatalakeanalytics.CreateDataLakeAnalyticsAccountParameters{
			Location: to.Ptr("eastus2"),
			Properties: &armdatalakeanalytics.CreateDataLakeAnalyticsAccountProperties{
				ComputePolicies: []*armdatalakeanalytics.CreateComputePolicyWithAccountParameters{
					{
						Name: to.Ptr("test_policy"),
						Properties: &armdatalakeanalytics.CreateOrUpdateComputePolicyProperties{
							MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
							MinPriorityPerJob:            to.Ptr[int32](1),
							ObjectID:                     to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
							ObjectType:                   to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
						},
					}},
				DataLakeStoreAccounts: []*armdatalakeanalytics.AddDataLakeStoreWithAccountParameters{
					{
						Name: to.Ptr("test_adls"),
						Properties: &armdatalakeanalytics.AddDataLakeStoreProperties{
							Suffix: to.Ptr("test_suffix"),
						},
					}},
				DefaultDataLakeStoreAccount: to.Ptr("test_adls"),
				FirewallAllowAzureIPs:       to.Ptr(armdatalakeanalytics.FirewallAllowAzureIPsStateEnabled),
				FirewallRules: []*armdatalakeanalytics.CreateFirewallRuleWithAccountParameters{
					{
						Name: to.Ptr("test_rule"),
						Properties: &armdatalakeanalytics.CreateOrUpdateFirewallRuleProperties{
							EndIPAddress:   to.Ptr("2.2.2.2"),
							StartIPAddress: to.Ptr("1.1.1.1"),
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
						Name: to.Ptr("test_storage"),
						Properties: &armdatalakeanalytics.AddStorageAccountProperties{
							AccessKey: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
							Suffix:    to.Ptr("test_suffix"),
						},
					}},
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
			},
		},
		nil)
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
