package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PostRules_getCounters_MaximumSet_Gen.json
func ExamplePostRulesClient_GetCounters_postRulesGetCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPostRulesClient().GetCounters(ctx, "lrs1", "1", &armpanngfw.PostRulesClientGetCountersOptions{
		FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.PostRulesClientGetCountersResponse{
	// 	RuleCounter: &armpanngfw.RuleCounter{
	// 		AppSeen: &armpanngfw.AppSeenData{
	// 			AppSeenList: []*armpanngfw.AppSeenInfo{
	// 				{
	// 					Category: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 					Risk: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					StandardPorts: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 					SubCategory: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 					Tag: to.Ptr("aaaaaaaaaa"),
	// 					Technology: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 					Title: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 			Count: to.Ptr[int32](13),
	// 		},
	// 		FirewallName: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		HitCount: to.Ptr[int32](20),
	// 		LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		Priority: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		RequestTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		RuleListName: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 		RuleName: to.Ptr("aaaa"),
	// 		RuleStackName: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 		Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 	},
	// }
}
