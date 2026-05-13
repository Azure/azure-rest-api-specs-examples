package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/Monitors_ListLinkableEnvironments_MaximumSet_Gen.json
func ExampleMonitorsClient_NewListLinkableEnvironmentsPager_monitorsListLinkableEnvironmentsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armdynatrace.MonitorsClientListLinkableEnvironmentsResponse{
		// 	LinkableEnvironmentListResponse: armdynatrace.LinkableEnvironmentListResponse{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Dynatrace.Observability/monitors?api-version=2024-04-24&$skiptoken=abc123"),
		// 		Value: []*armdynatrace.LinkableEnvironmentResponse{
		// 			{
		// 				EnvironmentID: to.Ptr("abc.123"),
		// 				EnvironmentName: to.Ptr("myEnv"),
		// 				PlanData: &armdynatrace.PlanData{
		// 					BillingCycle: to.Ptr("Monthly"),
		// 					EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t}()),
		// 					PlanDetails: to.Ptr("dynatraceapitestplan"),
		// 					UsageType: to.Ptr("Committed"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
