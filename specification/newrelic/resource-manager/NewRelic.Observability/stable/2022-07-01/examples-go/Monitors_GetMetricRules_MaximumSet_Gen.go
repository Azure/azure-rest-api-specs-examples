package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_GetMetricRules_MaximumSet_Gen.json
func ExampleMonitorsClient_GetMetricRules_monitorsGetMetricRulesMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().GetMetricRules(ctx, "rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz", armnewrelicobservability.MetricsRequest{
		UserEmail: to.Ptr("ruxvg@xqkmdhrnoo.hlmbpm"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetricRules = armnewrelicobservability.MetricRules{
	// 	FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 		{
	// 			Name: to.Ptr("qnvixg"),
	// 			Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 			Value: to.Ptr("ihyabcsjvhkfzckfjvgvtlhdyvmwge"),
	// 	}},
	// 	SendMetrics: to.Ptr(armnewrelicobservability.SendMetricsStatusEnabled),
	// 	UserEmail: to.Ptr("test@testing.com"),
	// }
}
