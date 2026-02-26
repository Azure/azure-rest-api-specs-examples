package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2024-11-30-preview/IspCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet_Gen.json
func ExampleIspCacheNodesOperationsClient_GetCacheNodeAutoUpdateHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIspCacheNodesOperationsClient().GetCacheNodeAutoUpdateHistory(ctx, "rgConnectedCache", "MccRPTest1", "MCCCachenode1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.IspCacheNodesOperationsClientGetCacheNodeAutoUpdateHistoryResponse{
	// 	MccCacheNodeAutoUpdateHistory: &armconnectedcache.MccCacheNodeAutoUpdateHistory{
	// 		Properties: &armconnectedcache.MccCacheNodeAutoUpdateHistoryProperties{
	// 			CustomerID: to.Ptr("fqyulypmqacwoifqnddnkcexbgm"),
	// 			CacheNodeID: to.Ptr("elwtomk"),
	// 			AutoUpdateHistory: []*armconnectedcache.MccCacheNodeAutoUpdateInfo{
	// 				{
	// 					ImageURIBeforeUpdate: to.Ptr("ne"),
	// 					ImageURITargeted: to.Ptr("zqgjxlqoucwyjf"),
	// 					ImageURITerminal: to.Ptr("akng"),
	// 					AutoUpdateRingType: to.Ptr[int32](16),
	// 					MovedToTerminalStateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-08T07:01:45.657Z"); return t}()),
	// 					RuleRequestedWeek: to.Ptr[int32](1),
	// 					RuleRequestedDay: to.Ptr[int32](9),
	// 					CreatedDateTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-08T07:01:45.657Z"); return t}()),
	// 					UpdatedRegistryDateTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-08T07:01:45.657Z"); return t}()),
	// 					PlanChangeLogText: to.Ptr("uhfsokrwx"),
	// 					AutoUpdateLastAppliedStatus: to.Ptr[int32](24),
	// 					AutoUpdateLastAppliedStatusText: to.Ptr("cwoy"),
	// 					AutoUpdateLastAppliedStatusDetailedText: to.Ptr("oicfonhqlnrc"),
	// 					PlanID: to.Ptr[int64](17),
	// 					TimeToGoLiveDateTime: to.Ptr("wqhjexgtkqzu"),
	// 					RuleRequestedMinute: to.Ptr("ewoqhdmofybbpf"),
	// 					RuleRequestedHour: to.Ptr("degiarxknlfqfgwz"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key5753": to.Ptr("ubrjiectme"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rgConnectedCache/providers/Microsoft.ConnectedCache/ispCustomers/MccRPTest1/ispCacheNodes/MCCCachenode1"),
	// 		Name: to.Ptr("cgy"),
	// 		Type: to.Ptr("kqbnkmxgllkdxfnsogmmdogjtotj"),
	// 		SystemData: &armconnectedcache.SystemData{
	// 			CreatedBy: to.Ptr("gambtqj"),
	// 			CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
	// 			LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 		},
	// 	},
	// }
}
