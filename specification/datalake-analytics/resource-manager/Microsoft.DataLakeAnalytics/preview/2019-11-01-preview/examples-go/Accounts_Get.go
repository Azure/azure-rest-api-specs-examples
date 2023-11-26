package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "contosorg", "contosoadla", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armdatalakeanalytics.Account{
	// 	Name: to.Ptr("test_account"),
	// 	Type: to.Ptr("Microsoft.DataLakeAnalytics/accounts"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rgaba12041/providers/Microsoft.DataLakeAnalytics/accounts/testaba15818"),
	// 	Properties: &armdatalakeanalytics.AccountProperties{
	// 		AccountID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
	// 		Endpoint: to.Ptr("test_endpoint"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatalakeanalytics.DataLakeAnalyticsAccountStatusSucceeded),
	// 		State: to.Ptr(armdatalakeanalytics.DataLakeAnalyticsAccountStateActive),
	// 		ComputePolicies: []*armdatalakeanalytics.ComputePolicy{
	// 			{
	// 				Name: to.Ptr("test_policy"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("test_policy_id"),
	// 				Properties: &armdatalakeanalytics.ComputePolicyProperties{
	// 					MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
	// 					MinPriorityPerJob: to.Ptr[int32](1),
	// 					ObjectID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 					ObjectType: to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
	// 				},
	// 		}},
	// 		CurrentTier: to.Ptr(armdatalakeanalytics.TierTypeConsumption),
	// 		DataLakeStoreAccounts: []*armdatalakeanalytics.DataLakeStoreAccountInformation{
	// 			{
	// 				Name: to.Ptr("test_adls"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("test_adls_id"),
	// 				Properties: &armdatalakeanalytics.DataLakeStoreAccountInformationProperties{
	// 					Suffix: to.Ptr("test_suffix"),
	// 				},
	// 		}},
	// 		DefaultDataLakeStoreAccount: to.Ptr("test_adls"),
	// 		FirewallAllowAzureIPs: to.Ptr(armdatalakeanalytics.FirewallAllowAzureIPsStateEnabled),
	// 		FirewallRules: []*armdatalakeanalytics.FirewallRule{
	// 			{
	// 				Name: to.Ptr("test_rule"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("test_firewall_id"),
	// 				Properties: &armdatalakeanalytics.FirewallRuleProperties{
	// 					EndIPAddress: to.Ptr("2.2.2.2"),
	// 					StartIPAddress: to.Ptr("1.1.1.1"),
	// 				},
	// 		}},
	// 		FirewallState: to.Ptr(armdatalakeanalytics.FirewallStateEnabled),
	// 		MaxDegreeOfParallelism: to.Ptr[int32](30),
	// 		MaxDegreeOfParallelismPerJob: to.Ptr[int32](1),
	// 		MaxJobCount: to.Ptr[int32](3),
	// 		MinPriorityPerJob: to.Ptr[int32](1),
	// 		NewTier: to.Ptr(armdatalakeanalytics.TierTypeConsumption),
	// 		PublicDataLakeStoreAccounts: []*armdatalakeanalytics.DataLakeStoreAccountInformation{
	// 			{
	// 				Name: to.Ptr("test_adls"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("test_adls_id"),
	// 				Properties: &armdatalakeanalytics.DataLakeStoreAccountInformationProperties{
	// 					Suffix: to.Ptr("test_suffix"),
	// 				},
	// 		}},
	// 		QueryStoreRetention: to.Ptr[int32](30),
	// 		StorageAccounts: []*armdatalakeanalytics.StorageAccountInformation{
	// 			{
	// 				Name: to.Ptr("test_storage"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("test_storage_id"),
	// 				Properties: &armdatalakeanalytics.StorageAccountInformationProperties{
	// 					Suffix: to.Ptr("test_suffix"),
	// 				},
	// 		}},
	// 		SystemMaxDegreeOfParallelism: to.Ptr[int32](1),
	// 		SystemMaxJobCount: to.Ptr[int32](1),
	// 	},
	// }
}
