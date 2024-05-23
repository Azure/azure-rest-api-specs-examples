package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlert.json
func ExampleIotSecuritySolutionsAnalyticsAggregatedAlertClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotSecuritySolutionsAnalyticsAggregatedAlertClient().Get(ctx, "MyGroup", "default", "IoT_Bruteforce_Fail/2019-02-02", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IoTSecurityAggregatedAlert = armsecurity.IoTSecurityAggregatedAlert{
	// 	Name: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Fail/2019-02-02"),
	// 	Type: to.Ptr("Microsoft.Security/iotSecuritySolutions/analyticsModels/aggregatedAlerts"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/IoTSecuritySolutions/Locations/eastus/default/IoT_Bruteforce_Fail/2019-02-02"),
	// 	Properties: &armsecurity.IoTSecurityAggregatedAlertProperties{
	// 		Description: to.Ptr("Multiple unsuccsseful login attempts identified. A Bruteforce attack on the device failed."),
	// 		ActionTaken: to.Ptr("Detected"),
	// 		AggregatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-02-02"); return t}()),
	// 		AlertDisplayName: to.Ptr("Failed Bruteforce"),
	// 		AlertType: to.Ptr("IoT_Bruteforce_Fail"),
	// 		Count: to.Ptr[int64](50),
	// 		EffectedResourceType: to.Ptr("IoT Device"),
	// 		LogAnalyticsQuery: to.Ptr("SecurityAlert | where tolower(ResourceId) == tolower('/subscriptions/b77ec8a9-04ed-48d2-a87a-e5887b978ba6/resourceGroups/IoT-Solution-DemoEnv/providers/Microsoft.Devices/IotHubs/rtogm-hub') and tolower(AlertName) == tolower('Custom Alert - number of device to cloud messages in MQTT protocol is not in the allowed range') | extend DeviceId=parse_json(ExtendedProperties)['DeviceId'] | project DeviceId, TimeGenerated, DisplayName, AlertSeverity, Description, RemediationSteps, ExtendedProperties"),
	// 		RemediationSteps: to.Ptr(""),
	// 		ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 		SystemSource: to.Ptr("Devices"),
	// 		TopDevicesList: []*armsecurity.IoTSecurityAggregatedAlertPropertiesTopDevicesListItem{
	// 			{
	// 				AlertsCount: to.Ptr[int64](100),
	// 				DeviceID: to.Ptr("testDevice1"),
	// 				LastOccurrence: to.Ptr("10:42"),
	// 			},
	// 			{
	// 				AlertsCount: to.Ptr[int64](80),
	// 				DeviceID: to.Ptr("testDevice2"),
	// 				LastOccurrence: to.Ptr("15:42"),
	// 		}},
	// 		VendorName: to.Ptr("Microsoft"),
	// 	},
	// }
}
