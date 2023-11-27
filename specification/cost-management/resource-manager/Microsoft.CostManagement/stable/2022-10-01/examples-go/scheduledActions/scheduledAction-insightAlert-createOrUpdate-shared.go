package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-insightAlert-createOrUpdate-shared.json
func ExampleScheduledActionsClient_CreateOrUpdateByScope_createOrUpdateInsightAlertScheduledActionByScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().CreateOrUpdateByScope(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "dailyAnomalyByResource", armcostmanagement.ScheduledAction{
		Kind: to.Ptr(armcostmanagement.ScheduledActionKindInsightAlert),
		Properties: &armcostmanagement.ScheduledActionProperties{
			DisplayName: to.Ptr("Daily anomaly by resource"),
			Notification: &armcostmanagement.NotificationProperties{
				Subject: to.Ptr("Cost anomaly detected in the resource"),
				To: []*string{
					to.Ptr("user@gmail.com"),
					to.Ptr("team@gmail.com")},
			},
			Schedule: &armcostmanagement.ScheduleProperties{
				EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.128Z"); return t }()),
				Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyDaily),
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.128Z"); return t }()),
			},
			Status: to.Ptr(armcostmanagement.ScheduledActionStatusEnabled),
			ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
		},
	}, &armcostmanagement.ScheduledActionsClientCreateOrUpdateByScopeOptions{IfMatch: to.Ptr("")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScheduledAction = armcostmanagement.ScheduledAction{
	// 	Name: to.Ptr("dailyAnomalyByResource"),
	// 	Type: to.Ptr("Microsoft.CostManagement/ScheduledActions"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/scheduledActions/dailyAnomalyByResource"),
	// 	ETag: to.Ptr("\"1d4ff9fe66f1d10\""),
	// 	Kind: to.Ptr(armcostmanagement.ScheduledActionKindInsightAlert),
	// 	SystemData: &armcostmanagement.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.128Z"); return t}()),
	// 		CreatedBy: to.Ptr("testuser"),
	// 		CreatedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.128Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("testuser"),
	// 		LastModifiedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
	// 	},
	// 	Properties: &armcostmanagement.ScheduledActionProperties{
	// 		DisplayName: to.Ptr("Daily anomaly by resource"),
	// 		FileDestination: &armcostmanagement.FileDestination{
	// 			FileFormats: []*armcostmanagement.FileFormat{
	// 			},
	// 		},
	// 		Notification: &armcostmanagement.NotificationProperties{
	// 			Subject: to.Ptr("Cost anomaly detected in the resource"),
	// 			To: []*string{
	// 				to.Ptr("user@gmail.com"),
	// 				to.Ptr("team@gmail.com")},
	// 			},
	// 			Schedule: &armcostmanagement.ScheduleProperties{
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.128Z"); return t}()),
	// 				Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyDaily),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.128Z"); return t}()),
	// 			},
	// 			Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
	// 			Status: to.Ptr(armcostmanagement.ScheduledActionStatusEnabled),
	// 			ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
	// 		},
	// 	}
}
