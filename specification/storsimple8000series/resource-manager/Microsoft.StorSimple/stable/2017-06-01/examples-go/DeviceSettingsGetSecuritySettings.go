package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetSecuritySettings.json
func ExampleDeviceSettingsClient_GetSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetSecuritySettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecuritySettings = armstorsimple8000series.SecuritySettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/securitySettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.SecuritySettingsProperties{
	// 		ChapSettings: &armstorsimple8000series.ChapSettings{
	// 			InitiatorUser: to.Ptr("test-initiator-user"),
	// 			TargetUser: to.Ptr("test-target-user"),
	// 		},
	// 		RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettings{
	// 			RemoteManagementCertificate: to.Ptr("-----BEGIN CERTIFICATE-----\r\nMIIDfjCCAmagAwIBAgIQR9UA510Lj4FPZGVXI0oyQDANBgkqhkiG9w0BAQUFADBWMSAwHgYDVQQL\r\nDBdIY3NSZW1vdGVNYW5hZ2VtZW50Q2VydDEeMBwGA1UECgwVTWljcm9zb2Z0IENvcnBvcmF0aW9u\r\nMRIwEAYDVQQDDAkxMjM0NTY3ODkwIBcNMTcwNjA3MTExMDUzWhgPMjExNjA2MDcxMTEwNTNaMFYx\r\nIDAeBgNVBAsMF0hjc1JlbW90ZU1hbmFnZW1lbnRDZXJ0MR4wHAYDVQQKDBVNaWNyb3NvZnQgQ29y\r\ncG9yYXRpb24xEjAQBgNVBAMMCTEyMzQ1Njc4OTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\r\nggEBAMp25cGMmEJdwur4y2Qgs9AKimTYGHdL5R9o2rcf+rGpgLC2wMY4qWsz3Ub4EWDwgT26rCiX\r\nv/24FK3xGB0OLxzTC3EbBJ7g7uCkapAKp/cFcmr+uBZiBhXe3hNnkeP8H2IfXNRSKa54XUMSqoFi\r\naavSBstS8FMuDMKopIVe6o4j+oDZBJGbb4a2anxuSyBM3cNhFTaOCLnfKynREC4sbbBL6enaFs0i\r\nRdIobT1FDSb2sLSUz4ONa3ZN7Os25UFvj1OcWeDjT1vJKcjx+u5owILvifRGoGbPETtCjwJ306+7\r\nihl+0GQ7teD2+O1h/Q7qdBm3gQAX9ZTsQAoiRYyJ14UCAwEAAaNGMEQwEwYDVR0lBAwwCgYIKwYB\r\nBQUHAwEwHQYDVR0OBBYEFJNK6cTqsGCcZkjP1q0LF53B2+G1MA4GA1UdDwEB/wQEAwIFIDANBgkq\r\nhkiG9w0BAQUFAAOCAQEAl+2ApGz9hpt6+sND+Q3TGczBqh2BSprgK//KwPdWthSLnxPf7f06MITN\r\ng1sTxFlbrFb46Y0UPN5YKTMb7wxRd+Nqrz6l7X4kp6IR+2sMX1ydkWyJ6HZrosc96fkFvFUt9B+J\r\n25gb9YVRotNUeborkznDFW6VeeS0kxJVps5r3fcuhz2vOPW1q/U0UWNV+LnrzRT+7MauUTYyLe0q\r\nuobmnD8NC2HhZcJ04xYhCKOJQ/NNJyaURwGJ2TZSYwS0HOnQViZp/SipjhyOOwwB6h0W62ElOovN\r\nAxVgTBEQjbG0jsMOm079hLUKuIBYmH4HJDmF3QKml0er1NcXeieATPLOWQ==\r\n-----END CERTIFICATE-----\r\n"),
	// 			RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
	// 		},
	// 	},
	// }
}
