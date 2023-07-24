package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCveSummary_MaximumSet_Gen.json
func ExampleFirmwareClient_GenerateCveSummary_firmwareGenerateCveSummaryMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwareClient().GenerateCveSummary(ctx, "rgworkspaces-firmwares", "A7", "umrkdttp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CveSummary = armiotfirmwaredefense.CveSummary{
	// 	Critical: to.Ptr[int64](29),
	// 	High: to.Ptr[int64](10),
	// 	Low: to.Ptr[int64](27),
	// 	Medium: to.Ptr[int64](18),
	// 	Undefined: to.Ptr[int64](22),
	// 	Unknown: to.Ptr[int64](2),
	// }
}
