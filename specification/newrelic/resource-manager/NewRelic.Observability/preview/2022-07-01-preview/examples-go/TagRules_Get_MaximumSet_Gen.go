package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/TagRules_Get_MaximumSet_Gen.json
func ExampleTagRulesClient_Get_tagRulesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTagRulesClient().Get(ctx, "rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", "bxcantgzggsepbhqmedjqyrqeezmfb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagRule = armnewrelicobservability.TagRule{
	// 	Name: to.Ptr("ddoieozflu"),
	// 	Type: to.Ptr("roafonrkfwwuv"),
	// 	ID: to.Ptr("ycdsgeiitvxcd"),
	// 	SystemData: &armnewrelicobservability.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-02T00:39:52.964Z"); return t}()),
	// 		CreatedBy: to.Ptr("wqrkemruqrvclsoevdftfeof"),
	// 		CreatedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-02T00:39:52.964Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("usdyoriebtakpdotcfp"),
	// 		LastModifiedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnewrelicobservability.MonitoringTagRulesProperties{
	// 		LogRules: &armnewrelicobservability.LogRules{
	// 			FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 				{
	// 					Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 					Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 					Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 			}},
	// 			SendAADLogs: to.Ptr(armnewrelicobservability.SendAADLogsStatusEnabled),
	// 			SendActivityLogs: to.Ptr(armnewrelicobservability.SendActivityLogsStatusEnabled),
	// 			SendSubscriptionLogs: to.Ptr(armnewrelicobservability.SendSubscriptionLogsStatusEnabled),
	// 		},
	// 		MetricRules: &armnewrelicobservability.MetricRules{
	// 			FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 				{
	// 					Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 					Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 					Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 			}},
	// 			UserEmail: to.Ptr("test@testing.com"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateSucceeded),
	// 	},
	// }
}
