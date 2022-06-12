```go
package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesCreateOrUpdateSecuritySettings.json
func ExampleDevicesClient_BeginCreateOrUpdateSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewDevicesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateSecuritySettings(ctx,
		"HSDK-T4ZA3EAJFR",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.SecuritySettings{
			Properties: &armstorsimple1200series.SecuritySettingsProperties{
				DeviceAdminPassword: &armstorsimple1200series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:             to.Ptr(armstorsimple1200series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
					Value:                           to.Ptr("<value>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple1200series%2Farmstorsimple1200series%2Fv1.0.0/sdk/resourcemanager/storsimple1200series/armstorsimple1200series/README.md) on how to add the SDK to your project and authenticate.
