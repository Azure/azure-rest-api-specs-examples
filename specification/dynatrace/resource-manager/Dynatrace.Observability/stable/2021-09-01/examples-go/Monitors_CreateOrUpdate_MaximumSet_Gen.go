package armdynatrace_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_CreateOrUpdate_MaximumSet_Gen.json
func ExampleMonitorsClient_BeginCreateOrUpdate_monitorsCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdynatrace.NewMonitorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", armdynatrace.MonitorResource{
		Location: to.Ptr("West US 2"),
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
		Identity: &armdynatrace.IdentityProperties{
			Type: to.Ptr(armdynatrace.ManagedIdentityTypeSystemAssigned),
		},
		Properties: &armdynatrace.MonitorProperties{
			DynatraceEnvironmentProperties: &armdynatrace.EnvironmentProperties{
				AccountInfo:            &armdynatrace.AccountInfo{},
				EnvironmentInfo:        &armdynatrace.EnvironmentInfo{},
				SingleSignOnProperties: &armdynatrace.SingleSignOnProperties{},
			},
			LiftrResourceCategory:         to.Ptr(armdynatrace.LiftrResourceCategoriesUnknown),
			MarketplaceSubscriptionStatus: to.Ptr(armdynatrace.MarketplaceSubscriptionStatusActive),
			MonitoringStatus:              to.Ptr(armdynatrace.MonitoringStatusEnabled),
			PlanData: &armdynatrace.PlanData{
				BillingCycle:  to.Ptr("Monthly"),
				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t }()),
				PlanDetails:   to.Ptr("dynatraceapitestplan"),
				UsageType:     to.Ptr("Committed"),
			},
			ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateAccepted),
			UserInfo: &armdynatrace.UserInfo{
				Country:      to.Ptr("westus2"),
				EmailAddress: to.Ptr("alice@microsoft.com"),
				FirstName:    to.Ptr("Alice"),
				LastName:     to.Ptr("Bobab"),
				PhoneNumber:  to.Ptr("123456"),
			},
		},
	}, nil)
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
