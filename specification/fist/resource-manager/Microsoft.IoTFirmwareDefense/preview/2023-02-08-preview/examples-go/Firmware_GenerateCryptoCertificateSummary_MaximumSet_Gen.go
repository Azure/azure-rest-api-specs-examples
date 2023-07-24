package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoCertificateSummary_MaximumSet_Gen.json
func ExampleFirmwareClient_GenerateCryptoCertificateSummary_firmwareGenerateCryptoCertificateSummaryMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwareClient().GenerateCryptoCertificateSummary(ctx, "FirmwareAnalysisRG", "default", "DECAFBAD-0000-0000-0000-BADBADBADBAD", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CryptoCertificateSummary = armiotfirmwaredefense.CryptoCertificateSummary{
	// 	Expired: to.Ptr[int64](3),
	// 	ExpiringSoon: to.Ptr[int64](1),
	// 	PairedKeys: to.Ptr[int64](2),
	// 	SelfSigned: to.Ptr[int64](3),
	// 	ShortKeySize: to.Ptr[int64](1),
	// 	TotalCertificates: to.Ptr[int64](5),
	// 	WeakSignature: to.Ptr[int64](1),
	// }
}
