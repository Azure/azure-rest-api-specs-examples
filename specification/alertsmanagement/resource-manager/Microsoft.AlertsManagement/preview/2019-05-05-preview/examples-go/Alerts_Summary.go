package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_Summary.json
func ExampleAlertsClient_GetSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().GetSummary(ctx, armalertsmanagement.AlertsSummaryGroupByFields("severity,alertState"), &armalertsmanagement.AlertsClientGetSummaryOptions{IncludeSmartGroupsCount: nil,
		TargetResource:      nil,
		TargetResourceType:  nil,
		TargetResourceGroup: nil,
		MonitorService:      nil,
		MonitorCondition:    nil,
		Severity:            nil,
		AlertState:          nil,
		AlertRule:           nil,
		TimeRange:           nil,
		CustomTimeRange:     nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertsSummary = armalertsmanagement.AlertsSummary{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/alertsSummary"),
	// 	ID: to.Ptr("/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/providers/Microsoft.AlertsManagement/alertsSummary/current"),
	// 	Properties: &armalertsmanagement.AlertsSummaryGroup{
	// 		Groupedby: to.Ptr("severity"),
	// 		SmartGroupsCount: to.Ptr[int64](100),
	// 		Total: to.Ptr[int64](14189),
	// 		Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 			{
	// 				Name: to.Ptr("Sev0"),
	// 				Count: to.Ptr[int64](6517),
	// 				Groupedby: to.Ptr("alertState"),
	// 				Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 					{
	// 						Name: to.Ptr("New"),
	// 						Count: to.Ptr[int64](6517),
	// 					},
	// 					{
	// 						Name: to.Ptr("Acknowledged"),
	// 						Count: to.Ptr[int64](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("Closed"),
	// 						Count: to.Ptr[int64](0),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("Sev1"),
	// 				Count: to.Ptr[int64](3175),
	// 				Groupedby: to.Ptr("alertState"),
	// 				Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 					{
	// 						Name: to.Ptr("New"),
	// 						Count: to.Ptr[int64](3175),
	// 					},
	// 					{
	// 						Name: to.Ptr("Acknowledged"),
	// 						Count: to.Ptr[int64](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("Closed"),
	// 						Count: to.Ptr[int64](0),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("Sev2"),
	// 				Count: to.Ptr[int64](1120),
	// 				Groupedby: to.Ptr("alertState"),
	// 				Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 					{
	// 						Name: to.Ptr("New"),
	// 						Count: to.Ptr[int64](1120),
	// 					},
	// 					{
	// 						Name: to.Ptr("Acknowledged"),
	// 						Count: to.Ptr[int64](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("Closed"),
	// 						Count: to.Ptr[int64](0),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("Sev3"),
	// 				Count: to.Ptr[int64](1902),
	// 				Groupedby: to.Ptr("alertState"),
	// 				Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 					{
	// 						Name: to.Ptr("New"),
	// 						Count: to.Ptr[int64](1902),
	// 					},
	// 					{
	// 						Name: to.Ptr("Acknowledged"),
	// 						Count: to.Ptr[int64](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("Closed"),
	// 						Count: to.Ptr[int64](0),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("Sev4"),
	// 				Count: to.Ptr[int64](1475),
	// 				Groupedby: to.Ptr("alertState"),
	// 				Values: []*armalertsmanagement.AlertsSummaryGroupItem{
	// 					{
	// 						Name: to.Ptr("New"),
	// 						Count: to.Ptr[int64](1475),
	// 					},
	// 					{
	// 						Name: to.Ptr("Acknowledged"),
	// 						Count: to.Ptr[int64](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("Closed"),
	// 						Count: to.Ptr[int64](0),
	// 				}},
	// 		}},
	// 	},
	// }
}
