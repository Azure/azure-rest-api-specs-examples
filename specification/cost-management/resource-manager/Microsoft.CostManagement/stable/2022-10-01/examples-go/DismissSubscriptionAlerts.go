package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DismissSubscriptionAlerts.json
func ExampleAlertsClient_Dismiss_patchSubscriptionAlerts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().Dismiss(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "22222222-2222-2222-2222-222222222222", armcostmanagement.DismissAlertPayload{
		Properties: &armcostmanagement.AlertProperties{
			Status: to.Ptr(armcostmanagement.AlertStatusDismissed),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Alert = armcostmanagement.Alert{
	// 	Name: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 	Type: to.Ptr("Microsoft.CostManagement/alerts"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/alerts/22222222-2222-2222-2222-222222222222"),
	// 	Properties: &armcostmanagement.AlertProperties{
	// 		Description: to.Ptr(""),
	// 		CloseTime: to.Ptr("0001-01-01T00:00:00"),
	// 		CostEntityID: to.Ptr("budget1"),
	// 		CreationTime: to.Ptr("2020-04-27T11:07:52.7143901Z"),
	// 		Definition: &armcostmanagement.AlertPropertiesDefinition{
	// 			Type: to.Ptr(armcostmanagement.AlertTypeBudget),
	// 			Category: to.Ptr(armcostmanagement.AlertCategoryCost),
	// 			Criteria: to.Ptr(armcostmanagement.AlertCriteriaCostThresholdExceeded),
	// 		},
	// 		ModificationTime: to.Ptr("2020-04-28T11:06:02.8999373Z"),
	// 		Source: to.Ptr(armcostmanagement.AlertSourcePreset),
	// 		Status: to.Ptr(armcostmanagement.AlertStatusDismissed),
	// 		StatusModificationTime: to.Ptr("0001-01-01T00:00:00"),
	// 		Details: &armcostmanagement.AlertPropertiesDetails{
	// 			Amount: to.Ptr[float64](200000),
	// 			ContactEmails: []*string{
	// 				to.Ptr("1234@contoso.com")},
	// 				ContactGroups: []*string{
	// 				},
	// 				ContactRoles: []*string{
	// 				},
	// 				CurrentSpend: to.Ptr[float64](161000.12),
	// 				MeterFilter: []any{
	// 				},
	// 				Operator: to.Ptr(armcostmanagement.AlertOperatorGreaterThan),
	// 				PeriodStartDate: to.Ptr("2020-03-01T00:00:00Z"),
	// 				ResourceFilter: []any{
	// 				},
	// 				ResourceGroupFilter: []any{
	// 				},
	// 				TagFilter: map[string]any{
	// 				},
	// 				Threshold: to.Ptr[float64](0.8),
	// 				TimeGrainType: to.Ptr(armcostmanagement.AlertTimeGrainTypeQuarterly),
	// 				TriggeredBy: to.Ptr("22222222-2222-2222-2222-222222222222_1_01"),
	// 				Unit: to.Ptr("USD"),
	// 			},
	// 		},
	// 	}
}
