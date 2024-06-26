package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/Devices/Get.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "eastus", "myGroup", "fa30e727-16e1-4e81-84f1-d26b9153d1b2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceModel = armiotsecurity.DeviceModel{
	// 	Name: to.Ptr("fa30e727-16e1-4e81-84f1-d26b9153d1b2"),
	// 	Type: to.Ptr("Microsoft.IoTSecurity/locations/deviceGroups/devices"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.IoTSecurity/locations/eastus/deviceGroups/myGroup/devices/fa30e727-16e1-4e81-84f1-d26b9153d1b2"),
	// 	Properties: &armiotsecurity.DeviceProperties{
	// 		AuthorizedState: to.Ptr(armiotsecurity.AuthorizedStateAuthorized),
	// 		Criticality: to.Ptr(armiotsecurity.CriticalityNormal),
	// 		DeviceCategoryDisplayName: to.Ptr("OT"),
	// 		DeviceCategoryID: to.Ptr[int32](6),
	// 		DeviceDataSource: to.Ptr(armiotsecurity.DeviceDataSourceOtSensor),
	// 		DeviceName: to.Ptr("10.168.140.1"),
	// 		DeviceStatus: to.Ptr(armiotsecurity.DeviceStatusActive),
	// 		DeviceSubTypeDisplayName: to.Ptr("Historian"),
	// 		DeviceSubTypeID: to.Ptr[int32](2),
	// 		DeviceTypeDisplayName: to.Ptr("Industrial"),
	// 		DeviceTypeID: to.Ptr[int32](17),
	// 		FirstSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-13T06:32:25.000Z"); return t}()),
	// 		LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-13T06:32:25.000Z"); return t}()),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-13T06:32:25.000Z"); return t}()),
	// 		Nics: []*armiotsecurity.Nic{
	// 			{
	// 				IPv4Address: to.Ptr("10.168.140.1"),
	// 				LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-13T06:32:25.000Z"); return t}()),
	// 				MacAddress: to.Ptr("34-E1-2D-77-80-D0"),
	// 				MacCertainty: to.Ptr(armiotsecurity.MacCertaintyCertain),
	// 				Vlans: []*string{
	// 					to.Ptr("name(1)->2"),
	// 					to.Ptr("3"),
	// 					to.Ptr("another_name(4)")},
	// 			}},
	// 			OnboardingStatus: to.Ptr(armiotsecurity.OnboardingStatusInsufficientInfo),
	// 			OperatingSystem: &armiotsecurity.OperatingSystem{
	// 				Platform: to.Ptr("Windows"),
	// 				Version: to.Ptr("10\\1604"),
	// 			},
	// 			ProgrammingState: to.Ptr(armiotsecurity.ProgrammingStateNotProgrammingDevice),
	// 			PurdueLevel: to.Ptr(armiotsecurity.PurdueLevelProcessControl),
	// 			RiskScore: to.Ptr[int32](90),
	// 			SchemaVersion: to.Ptr("1"),
	// 			Sensor: &armiotsecurity.Sensor{
	// 				Name: to.Ptr("mySensor"),
	// 				Site: to.Ptr("mySite"),
	// 				Zone: to.Ptr("myZone"),
	// 			},
	// 		},
	// 		SystemData: &armiotsecurity.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armiotsecurity.CreatedByTypeUser),
	// 		},
	// 	}
}
