package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: 2025-05-25-preview/Alerts_List.json
func ExampleAlertsClient_NewGetAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("subscriptions/3b540246-808d-4331-99aa-917b808a9166", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsClient().NewGetAllPager(&armalertsmanagement.AlertsClientGetAllOptions{
		IncludeContext: to.Ptr(true)})
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
		// page = armalertsmanagement.AlertsClientGetAllResponse{
		// 	AlertsList: armalertsmanagement.AlertsList{
		// 		NextLink: to.Ptr("https://management.azure.com:443/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/providers/Microsoft.AlertsManagement/alerts?api-version=2018-05-05-preview&timeRange=1d&ctoken=%2bRID%3aPlwOAPHEGwB9UwEAAAAgCw%3d%3d%23RT%3a2%23TRC%3a500%23RTD%3aqtQyMDE4LTA2LTEyVDE1OjEyOjE1"),
		// 		Value: []*armalertsmanagement.Alert{
		// 			{
		// 				Name: to.Ptr("cpu alert"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
		// 				ID: to.Ptr("/subscriptions/3b540246-808d-4331-99aa-917b808a9166/resourcegroups/servicedeskresourcegroup/providers/microsoft.insights/components/servicedeskappinsight/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502100"),
		// 				Properties: &armalertsmanagement.AlertProperties{
		// 					Context: map[string]any{
		// 						"AffectedConfigurationItems": []any{
		// 							"",
		// 						},
		// 						"AlertRuleName": "Test number of results",
		// 						"AlertThresholdOperator": "Greater Than Or Equal To",
		// 						"AlertThresholdValue": 0,
		// 						"AlertType": "Number of results",
		// 						"Description": "",
		// 						"Frequency": 5,
		// 						"IncludeSearchResults": true,
		// 						"LinkToFilteredSearchResultsAPI": "https://api.applicationinsights.io/v1/apps/e72c8301-003e-4251-aac9-2374b3320ecf/query?query=traces&timespan=2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToFilteredSearchResultsUI": "https://portal.azure.com#@0ef55770-ee07-488d-8dc5-75f53fa5a901/blade/Microsoft_Azure_Monitoring_Logs/LogsBlade/source/Alerts.EmailLinks/scope/%7B%22resources%22%3A%5B%7B%22resourceId%22%3A%22%2Fsubscriptions%2F0ef55770-ee07-488d-8dc5-75f53fa5a901%2FresourceGroups%2Fexample_resource_group%2Fproviders%2Fmicrosoft.insights%2Fcomponents%2FPortal%22%7D%5D%7D/q/eJwrKUdd0ejhGAA%3D%3D/prettify/1/timespan/2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToSearchResults": "https://portal.azure.com#@0ef55770-ee07-488d-8dc5-75f53fa5a901/blade/Microsoft_Azure_Monitoring_Logs/LogsBlade/source/Alerts.EmailLinks/scope/%7B%22resources%22%3A%5B%7B%22resourceId%22%3A%22%2Fsubscriptions%2F0ef55770-ee07-488d-8dc5-75f53fa5a901%2FresourceGroups%2Fexample_name%2Fproviders%2Fmicrosoft.insights%2Fcomponents%2FPortal%22%7D%5D%7D/q/eJwreefe4tGAA%3D%3D/prettify/1/timespan/2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToSearchResultsAPI": "https://api.applicationinsights.io/v1/apps/e72c8301-003e-4251-aac9-2374b3320ecf/query?query=traces&timespan=2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"ResultCount": 3,
		// 						"SearchIntervalDurationMin": "30",
		// 						"SearchIntervalEndtimeUtc": "2023-04-19T12:32:25Z",
		// 						"SearchIntervalInSeconds": 1800,
		// 						"SearchIntervalStartTimeUtc": "2023-04-19T12:02:25Z",
		// 						"SearchQuery": "traces",
		// 						"SeverityDescription": "Informational",
		// 						"SubscriptionId": "0ef55770-ee07-488d-8dc5-75f53fa5a901",
		// 						"WorkspaceId": "e72c8301-003e-4251-aac9-2374b3320ecf",
		// 					},
		// 					CustomProperties: map[string]*string{
		// 						"key1": to.Ptr("value1"),
		// 						"key2": to.Ptr("value2"),
		// 					},
		// 					EgressConfig: map[string]any{
		// 					},
		// 					Essentials: &armalertsmanagement.Essentials{
		// 						Description: to.Ptr("description of the alert"),
		// 						ActionStatus: &armalertsmanagement.ActionStatus{
		// 							IsSuppressed: to.Ptr(false),
		// 						},
		// 						AlertRule: to.Ptr("https://servisdffsdf.portal.mms.microsoft.com/#Workspace/overview/settings/details/Edit Alert Rule/details/index?savedSearchId=&scheduleId="),
		// 						AlertState: to.Ptr(armalertsmanagement.AlertStateAcknowledged),
		// 						LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						LastModifiedUserName: to.Ptr("System"),
		// 						MonitorCondition: to.Ptr(armalertsmanagement.MonitorConditionFired),
		// 						MonitorConditionResolvedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						MonitorService: to.Ptr(armalertsmanagement.MonitorServiceApplicationInsights),
		// 						Severity: to.Ptr(armalertsmanagement.SeveritySev3),
		// 						SignalType: to.Ptr(armalertsmanagement.SignalTypeLog),
		// 						SmartGroupID: to.Ptr("23d6b2ce-8c54-468f-aff0-sd32aebb7e56"),
		// 						SmartGroupingReason: to.Ptr("Occurred frequently with other alerts"),
		// 						SourceCreatedID: to.Ptr("6cd6b2ce-8c54-468f-aff0-9d12aebb7e49"),
		// 						StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						TargetResource: to.Ptr("/subscriptions/3b540246-808d-4331-99aa-917b808a9166/resourcegroups/servicedeskresourcegroup/providers/microsoft.insights/components/servicedeskappinsight"),
		// 						TargetResourceGroup: to.Ptr("servicedeskresourcegroup"),
		// 						TargetResourceName: to.Ptr("servicedeskappinsight"),
		// 						TargetResourceType: to.Ptr("components"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("cpu alert"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
		// 				ID: to.Ptr("/subscriptions/3b540246-808d-4331-99aa-917b808a9166/resourcegroups/cind/providers/microsoft.operationalinsights/workspaces/servicedeskwcus/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502100"),
		// 				Properties: &armalertsmanagement.AlertProperties{
		// 					Context: map[string]any{
		// 						"AffectedConfigurationItems": []any{
		// 							"",
		// 						},
		// 						"AlertRuleName": "Test number of results",
		// 						"AlertThresholdOperator": "Greater Than Or Equal To",
		// 						"AlertThresholdValue": 0,
		// 						"AlertType": "Number of results",
		// 						"Description": "",
		// 						"Frequency": 5,
		// 						"IncludeSearchResults": true,
		// 						"LinkToFilteredSearchResultsAPI": "https://api.applicationinsights.io/v1/apps/e72c8301-003e-4251-aac9-2374b3320ecf/query?query=traces&timespan=2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToFilteredSearchResultsUI": "https://portal.azure.com#@0ef55770-ee07-488d-8dc5-75f53fa5a901/blade/Microsoft_Azure_Monitoring_Logs/LogsBlade/source/Alerts.EmailLinks/scope/%7B%22resources%22%3A%5B%7B%22resourceId%22%3A%22%2Fsubscriptions%2F0ef55770-ee07-488d-8dc5-75f53fa5a901%2FresourceGroups%2Fexample_resource_group%2Fproviders%2Fmicrosoft.insights%2Fcomponents%2FPortal%22%7D%5D%7D/q/eJwrKUdd0ejhGAA%3D%3D/prettify/1/timespan/2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToSearchResults": "https://portal.azure.com#@0ef55770-ee07-488d-8dc5-75f53fa5a901/blade/Microsoft_Azure_Monitoring_Logs/LogsBlade/source/Alerts.EmailLinks/scope/%7B%22resources%22%3A%5B%7B%22resourceId%22%3A%22%2Fsubscriptions%2F0ef55770-ee07-488d-8dc5-75f53fa5a901%2FresourceGroups%2Fexample_name%2Fproviders%2Fmicrosoft.insights%2Fcomponents%2FPortal%22%7D%5D%7D/q/eJwreefe4tGAA%3D%3D/prettify/1/timespan/2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"LinkToSearchResultsAPI": "https://api.applicationinsights.io/v1/apps/e72c8301-003e-4251-aac9-2374b3320ecf/query?query=traces&timespan=2023-04-19T12%3a02%3a25.0000000Z%2f2023-04-19T12%3a32%3a25.0000000Z",
		// 						"ResultCount": 3,
		// 						"SearchIntervalDurationMin": "30",
		// 						"SearchIntervalEndtimeUtc": "2023-04-19T12:32:25Z",
		// 						"SearchIntervalInSeconds": 1800,
		// 						"SearchIntervalStartTimeUtc": "2023-04-19T12:02:25Z",
		// 						"SearchQuery": "traces",
		// 						"SeverityDescription": "Informational",
		// 						"SubscriptionId": "0ef55770-ee07-488d-8dc5-75f53fa5a901",
		// 						"WorkspaceId": "e72c8301-003e-4251-aac9-2374b3320ecf",
		// 					},
		// 					CustomProperties: map[string]*string{
		// 						"category": to.Ptr("performance"),
		// 						"environment": to.Ptr("production"),
		// 						"priority": to.Ptr("high"),
		// 					},
		// 					EgressConfig: map[string]any{
		// 					},
		// 					Essentials: &armalertsmanagement.Essentials{
		// 						Description: to.Ptr("description of the alert"),
		// 						ActionStatus: &armalertsmanagement.ActionStatus{
		// 							IsSuppressed: to.Ptr(false),
		// 						},
		// 						AlertRule: to.Ptr("https://servicsdfsdf.portal.mms.microsoft.com/#Workspace/overview/settings/details/Edit Alert Rule/details/index?savedSearchId=&scheduleId="),
		// 						AlertState: to.Ptr(armalertsmanagement.AlertStateNew),
		// 						LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						LastModifiedUserName: to.Ptr("System"),
		// 						MonitorCondition: to.Ptr(armalertsmanagement.MonitorConditionFired),
		// 						MonitorConditionResolvedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						MonitorService: to.Ptr(armalertsmanagement.MonitorServiceApplicationInsights),
		// 						Severity: to.Ptr(armalertsmanagement.SeveritySev3),
		// 						SignalType: to.Ptr(armalertsmanagement.SignalTypeLog),
		// 						SmartGroupID: to.Ptr("d1c49c89-ea95-4697-a299-c0f5ebac62f1"),
		// 						SmartGroupingReason: to.Ptr("Alerts that frequently occur together have been grouped."),
		// 						SourceCreatedID: to.Ptr("6cd6b2ce-8c54-468f-aff0-9d12aebb7e49"),
		// 						StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:51:02Z"); return t}()),
		// 						TargetResource: to.Ptr("/subscriptions/3b540246-808d-4331-99aa-917b808a9166/resourcegroups/cind/providers/microsoft.operationalinsights/workspaces/servicedeskwcus"),
		// 						TargetResourceGroup: to.Ptr("servicedeskresourcegroup"),
		// 						TargetResourceName: to.Ptr("servicedeskwcus"),
		// 						TargetResourceType: to.Ptr("components"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
