package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/Sensors/Get.json
func ExampleSensorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSensorsClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/myHub", "mySensor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SensorModel = armiotsecurity.SensorModel{
	// 	Name: to.Ptr("mySensor"),
	// 	Type: to.Ptr("Microsoft.IoTSecurity/sensors"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/myHub/providers/Microsoft.IoTSecurity/sensors/mySensor"),
	// 	Properties: &armiotsecurity.SensorProperties{
	// 		ConnectivityTime: to.Ptr("2020-11-17T12:31:25Z"),
	// 		DynamicLearning: to.Ptr(true),
	// 		LearningMode: to.Ptr(true),
	// 		SensorStatus: to.Ptr(armiotsecurity.SensorStatusOk),
	// 		SensorType: to.Ptr(armiotsecurity.SensorTypeOt),
	// 		SensorVersion: to.Ptr("2020.11.01.1643"),
	// 		TiAutomaticUpdates: to.Ptr(true),
	// 		TiStatus: to.Ptr(armiotsecurity.TiStatusOk),
	// 		TiVersion: to.Ptr("2020-11-17T12:31:25Z"),
	// 		Zone: to.Ptr("Zone Name"),
	// 	},
	// 	SystemData: &armiotsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
	// 	},
	// }
}
