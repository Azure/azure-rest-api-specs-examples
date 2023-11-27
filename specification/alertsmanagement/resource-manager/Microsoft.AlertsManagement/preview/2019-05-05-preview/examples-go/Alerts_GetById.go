package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_GetById.json
func ExampleAlertsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().GetByID(ctx, "66114d64-d9d9-478b-95c9-b789d6502100", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Alert = armalertsmanagement.Alert{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
	// 	ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502100"),
	// 	Properties: &armalertsmanagement.AlertProperties{
	// 		Context: map[string]any{
	// 		},
	// 		EgressConfig: map[string]any{
	// 		},
	// 		Essentials: &armalertsmanagement.Essentials{
	// 			Description: to.Ptr("description of the alert"),
	// 			ActionStatus: &armalertsmanagement.ActionStatus{
	// 				IsSuppressed: to.Ptr(false),
	// 			},
	// 			AlertRule: to.Ptr("https://sersdfsdfinsight.portal.mms.microsoft.com/#Workspace/overview/settings/details/Edit Alert Rule/details/index?savedSearchId=&scheduleId="),
	// 			AlertState: to.Ptr(armalertsmanagement.AlertStateAcknowledged),
	// 			LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02.000Z"); return t}()),
	// 			LastModifiedUserName: to.Ptr("System"),
	// 			MonitorCondition: to.Ptr(armalertsmanagement.MonitorConditionFired),
	// 			MonitorConditionResolvedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02.000Z"); return t}()),
	// 			MonitorService: to.Ptr(armalertsmanagement.MonitorServiceApplicationInsights),
	// 			Severity: to.Ptr(armalertsmanagement.SeveritySev3),
	// 			SignalType: to.Ptr(armalertsmanagement.SignalTypeLog),
	// 			SmartGroupID: to.Ptr("23d6b2ce-8c54-468f-aff0-sd32aebb7e56"),
	// 			SmartGroupingReason: to.Ptr("Occurred frequently with other alerts"),
	// 			SourceCreatedID: to.Ptr("6cd6b2ce-8c54-468f-aff0-9d12aebb7e49"),
	// 			StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02.000Z"); return t}()),
	// 			TargetResource: to.Ptr("/subscriptions/3b540246-808d-4331-99aa-917b808a9166/resourcegroups/servicedeskresourcegroup/providers/microsoft.insights/components/servicedeskappinsight"),
	// 			TargetResourceGroup: to.Ptr("servicedeskresourcegroup"),
	// 			TargetResourceName: to.Ptr("servicedeskappinsight"),
	// 			TargetResourceType: to.Ptr("components"),
	// 		},
	// 	},
	// }
}
