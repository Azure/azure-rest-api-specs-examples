package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_CreateOrUpdate_MaximumSet_Gen.json
func ExampleTagRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdynatrace.NewTagRulesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myMonitor",
		"default",
		armdynatrace.TagRule{
			Properties: &armdynatrace.MonitoringTagRulesProperties{
				LogRules: &armdynatrace.LogRules{
					FilteringTags: []*armdynatrace.FilteringTag{
						{
							Name:   to.Ptr("Environment"),
							Action: to.Ptr(armdynatrace.TagActionInclude),
							Value:  to.Ptr("Prod"),
						},
						{
							Name:   to.Ptr("Environment"),
							Action: to.Ptr(armdynatrace.TagActionExclude),
							Value:  to.Ptr("Dev"),
						}},
					SendAADLogs:          to.Ptr(armdynatrace.SendAADLogsStatusEnabled),
					SendActivityLogs:     to.Ptr(armdynatrace.SendActivityLogsStatusEnabled),
					SendSubscriptionLogs: to.Ptr(armdynatrace.SendSubscriptionLogsStatusEnabled),
				},
				MetricRules: &armdynatrace.MetricRules{
					FilteringTags: []*armdynatrace.FilteringTag{
						{
							Name:   to.Ptr("Environment"),
							Action: to.Ptr(armdynatrace.TagActionInclude),
							Value:  to.Ptr("Prod"),
						}},
				},
				ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateAccepted),
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
