package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cecf11996f2bfc7958f6d0de814badab23f9720/specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/RegionConfigurationByResourceGroup.json
func ExampleServiceClient_RegionConfigurationByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().RegionConfigurationByResourceGroup(ctx, "YourResourceGroupName", "westus", armdatabox.RegionConfigurationRequest{
		DeviceCapabilityRequest: &armdatabox.DeviceCapabilityRequest{
			Model:   to.Ptr(armdatabox.ModelNameDataBoxDisk),
			SKUName: to.Ptr(armdatabox.SKUNameDataBoxDisk),
		},
		ScheduleAvailabilityRequest: &armdatabox.ScheduleAvailabilityRequest{
			Model:           to.Ptr(armdatabox.ModelNameDataBox),
			SKUName:         to.Ptr(armdatabox.SKUNameDataBox),
			StorageLocation: to.Ptr("westus"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegionConfigurationResponse = armdatabox.RegionConfigurationResponse{
	// 	DeviceCapabilityResponse: &armdatabox.DeviceCapabilityResponse{
	// 		DeviceCapabilityDetails: []*armdatabox.DeviceCapabilityDetails{
	// 			{
	// 				HardwareEncryption: to.Ptr(armdatabox.HardwareEncryptionEnabled),
	// 		}},
	// 	},
	// 	ScheduleAvailabilityResponse: &armdatabox.ScheduleAvailabilityResponse{
	// 		AvailableDates: []*time.Time{
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-11T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-12T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-13T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-14T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-15T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-16T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-17T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-18T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-19T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-20T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-22T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-23T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-24T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-25T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-26T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-27T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-28T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-29T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-30T00:00:00.000Z"); return t}()),
	// 			to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T00:00:00.000Z"); return t}())},
	// 		},
	// 	}
}
