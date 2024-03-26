package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/09c37754dac91874ff689ed1e60effb4268c8669/specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetDevice.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1", "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armsphere.Device{
	// 	Name: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/myProduct1/deviceGroups/myDeviceGroup1/devices/00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	// 	Properties: &armsphere.DeviceProperties{
	// 		ChipSKU: to.Ptr("MyChipSku1"),
	// 		LastAvailableOsVersion: to.Ptr("AvailableOsVersion1"),
	// 		LastInstalledOsVersion: to.Ptr("InstalledOsVersion1"),
	// 		LastOsUpdateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-30T23:54:21.960Z"); return t}()),
	// 		LastUpdateRequestUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-01T23:54:21.960Z"); return t}()),
	// 	},
	// }
}
