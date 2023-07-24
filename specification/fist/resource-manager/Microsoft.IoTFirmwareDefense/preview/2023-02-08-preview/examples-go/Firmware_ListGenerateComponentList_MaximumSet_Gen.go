package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListGenerateComponentList_MaximumSet_Gen.json
func ExampleFirmwareClient_NewListGenerateComponentListPager_firmwareListGenerateComponentListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirmwareClient().NewListGenerateComponentListPager("rgworkspaces-firmwares", "A7", "umrkdttp", nil)
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
		// page.ComponentList = armiotfirmwaredefense.ComponentList{
		// 	Value: []*armiotfirmwaredefense.Component{
		// 		{
		// 			ComponentID: to.Ptr("dwtwgmbjtftea"),
		// 			ComponentName: to.Ptr("wbjvfrmhmumeahdhkla"),
		// 			IsUpdateAvailable: to.Ptr(armiotfirmwaredefense.IsUpdateAvailableTrue),
		// 			License: to.Ptr("tirj"),
		// 			Paths: []*string{
		// 				to.Ptr("bjxczktgyhpatxkgwomhgfsw")},
		// 				ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T06:40:30.304Z"); return t}()),
		// 				Version: to.Ptr("dmokuviyznjbdkfxbmvhlnglfnf"),
		// 		}},
		// 	}
	}
}
