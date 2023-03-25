package armnewrelicobservability_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Monitors_SwitchBilling_MaximumSet_Gen.json
func ExampleMonitorsClient_SwitchBilling_monitorsSwitchBillingMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMonitorsClient().SwitchBilling(ctx, "rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz", armnewrelicobservability.SwitchBillingRequest{
		AzureResourceID: to.Ptr("enfghpfw"),
		OrganizationID:  to.Ptr("k"),
		PlanData: &armnewrelicobservability.PlanData{
			BillingCycle:  to.Ptr(armnewrelicobservability.BillingCycleYEARLY),
			EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.786Z"); return t }()),
			PlanDetails:   to.Ptr("tbbiaga"),
			UsageType:     to.Ptr(armnewrelicobservability.UsageTypePAYG),
		},
		UserEmail: to.Ptr("ruxvg@xqkmdhrnoo.hlmbpm"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
