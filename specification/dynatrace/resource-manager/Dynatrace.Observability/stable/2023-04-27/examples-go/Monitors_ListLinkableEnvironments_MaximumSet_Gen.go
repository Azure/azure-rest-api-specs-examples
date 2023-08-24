package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListLinkableEnvironments_MaximumSet_Gen.json
func ExampleMonitorsClient_NewListLinkableEnvironmentsPager_monitorsListLinkableEnvironmentsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListLinkableEnvironmentsPager("myResourceGroup", "myMonitor", armdynatrace.LinkableEnvironmentRequest{
		Region:        to.Ptr("East US"),
		TenantID:      to.Ptr("00000000-0000-0000-0000-000000000000"),
		UserPrincipal: to.Ptr("alice@microsoft.com"),
	}, nil)
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
		// page.LinkableEnvironmentListResponse = armdynatrace.LinkableEnvironmentListResponse{
		// 	Value: []*armdynatrace.LinkableEnvironmentResponse{
		// 		{
		// 			EnvironmentID: to.Ptr("abc.123"),
		// 			EnvironmentName: to.Ptr("myEnv"),
		// 			PlanData: &armdynatrace.PlanData{
		// 				BillingCycle: to.Ptr("Monthly"),
		// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t}()),
		// 				PlanDetails: to.Ptr("dynatraceapitestplan"),
		// 				UsageType: to.Ptr("Committed"),
		// 			},
		// 	}},
		// }
	}
}
