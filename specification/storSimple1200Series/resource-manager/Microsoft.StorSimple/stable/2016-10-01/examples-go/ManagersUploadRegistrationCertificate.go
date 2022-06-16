package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersUploadRegistrationCertificate.json
func ExampleManagersClient_UploadRegistrationCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewManagersClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UploadRegistrationCertificate(ctx,
		"windows",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest2",
		armstorsimple1200series.UploadCertificateRequest{
			ContractVersion: to.Ptr(armstorsimple1200series.ContractVersionsV201212),
			Properties: &armstorsimple1200series.RawCertificateData{
				AuthType:    to.Ptr(armstorsimple1200series.AuthTypeAzureActiveDirectory),
				Certificate: to.Ptr("MIIC3TCCAcWgAwIBAgIQEr0bAWD6wJtA4LIbZ9NtgzANBgkqhkiG9w0BAQUFADAeMRwwGgYDVQQDExNXaW5kb3dzIEF6dXJlIFRvb2xzMB4XDTE4MDkxMDE1MzY0MFoXDTE4MDkxMzE1NDY0MFowHjEcMBoGA1UEAxMTV2luZG93cyBBenVyZSBUb29sczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANUsKkz2Z4fECKMyNeLb9v3pr1XF4dVe+MITDtgphjl81ng190Y0IHgCVnh4YjfplUSpMk/1xii0rI5AAPwoz3ze2qRPtnwCiTaoVLkUe6knNRPzrRvVXHB81J0/14MO0lwdByHhdccRcVJZPLt5724t4aQny82v2AayJdDDkBBWNlpcqPy6n3sygP00THMPP0O3sFqy924eHqoDj3qSw79/meaZBJt9S5odPuFoskxjHuI4lM6BmK1Ql7p8Wo9/GhTOIoMz81orKPHRDleLjutwL4mb6NnhI5rfT/MxnHD6m82c4YYqiZC3XiTyJWVCkWkp7PK92OdRp6FA87rdKDMCAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADggEBAIYlezVU68TuEblkn06vM5dfzSmHKJOQgW61nDlLnyKrmSJtzKZLCAswTE2VyJHwKNdZgW15coJFINjWBLWcLr0/GjNV4u3Z+UL3NhBFQd5xuMtKsIhuoscKtyk0JHQXpBvHNmOUCobfQfOBQfTVC7kmyWdtlGztFUVxD28s6S5gMb1FEWWN68NOOJ3/ZhaTbUEM54yw8Hk8/f0L/Zn/7BYHUyWWA3KStAaYn89C/ZFF+952ark2VaKGIjBRQzgrJEIR8dI4r46I3DoEfzGPESKvQPvVLhOX84RG0PLPOtnRbHBVew1Nh3HE9kgCubkPKK+NPWE9IHZPoRmOTWBe+zU="),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
