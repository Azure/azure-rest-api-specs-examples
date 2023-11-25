package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersUploadRegistrationCertificate.json
func ExampleManagersClient_UploadRegistrationCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().UploadRegistrationCertificate(ctx, "windows", "ResourceGroupForSDKTest", "ManagerForSDKTest2", armstorsimple1200series.UploadCertificateRequest{
		ContractVersion: to.Ptr(armstorsimple1200series.ContractVersionsV201212),
		Properties: &armstorsimple1200series.RawCertificateData{
			AuthType:    to.Ptr(armstorsimple1200series.AuthTypeAzureActiveDirectory),
			Certificate: to.Ptr("MIIC3TCCAcWgAwIBAgIQEr0bAWD6wJtA4LIbZ9NtgzANBgkqhkiG9w0BAQUFADAeMRwwGgYDVQQDExNXaW5kb3dzIEF6dXJlIFRvb2xzMB4XDTE4MDkxMDE1MzY0MFoXDTE4MDkxMzE1NDY0MFowHjEcMBoGA1UEAxMTV2luZG93cyBBenVyZSBUb29sczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANUsKkz2Z4fECKMyNeLb9v3pr1XF4dVe+MITDtgphjl81ng190Y0IHgCVnh4YjfplUSpMk/1xii0rI5AAPwoz3ze2qRPtnwCiTaoVLkUe6knNRPzrRvVXHB81J0/14MO0lwdByHhdccRcVJZPLt5724t4aQny82v2AayJdDDkBBWNlpcqPy6n3sygP00THMPP0O3sFqy924eHqoDj3qSw79/meaZBJt9S5odPuFoskxjHuI4lM6BmK1Ql7p8Wo9/GhTOIoMz81orKPHRDleLjutwL4mb6NnhI5rfT/MxnHD6m82c4YYqiZC3XiTyJWVCkWkp7PK92OdRp6FA87rdKDMCAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADggEBAIYlezVU68TuEblkn06vM5dfzSmHKJOQgW61nDlLnyKrmSJtzKZLCAswTE2VyJHwKNdZgW15coJFINjWBLWcLr0/GjNV4u3Z+UL3NhBFQd5xuMtKsIhuoscKtyk0JHQXpBvHNmOUCobfQfOBQfTVC7kmyWdtlGztFUVxD28s6S5gMb1FEWWN68NOOJ3/ZhaTbUEM54yw8Hk8/f0L/Zn/7BYHUyWWA3KStAaYn89C/ZFF+952ark2VaKGIjBRQzgrJEIR8dI4r46I3DoEfzGPESKvQPvVLhOX84RG0PLPOtnRbHBVew1Nh3HE9kgCubkPKK+NPWE9IHZPoRmOTWBe+zU="),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UploadCertificateResponse = armstorsimple1200series.UploadCertificateResponse{
	// 	Name: to.Ptr("windows"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/certificates"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourcegroups/cleanupservice/providers/Microsoft.StorSimple/managers/res-jem-hel/certificates/windows"),
	// 	Properties: &armstorsimple1200series.ResourceCertificateAndAADDetails{
	// 		AADAuthority: to.Ptr("https://login.windows.net"),
	// 		AADTenantID: to.Ptr("625ad78b-e6ce-49ed-b9cd-29f4aa12d77e"),
	// 		AuthType: to.Ptr(armstorsimple1200series.AuthTypeAzureActiveDirectory),
	// 		AzureManagementEndpointAudience: to.Ptr("https://pod01-id1.wus.backup.windowsazure.com/restapi/"),
	// 		Certificate: to.Ptr("MIIC3TCCAcWgAwIBAgIQEr0bAWD6wJtA4LIbZ9NtgzANBgkqhkiG9w0BAQUFADAeMRwwGgYDVQQDExNXaW5kb3dzIEF6dXJlIFRvb2xzMB4XDTE4MDkxMDE1MzY0MFoXDTE4MDkxMzE1NDY0MFowHjEcMBoGA1UEAxMTV2luZG93cyBBenVyZSBUb29sczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANUsKkz2Z4fECKMyNeLb9v3pr1XF4dVe+MITDtgphjl81ng190Y0IHgCVnh4YjfplUSpMk/1xii0rI5AAPwoz3ze2qRPtnwCiTaoVLkUe6knNRPzrRvVXHB81J0/14MO0lwdByHhdccRcVJZPLt5724t4aQny82v2AayJdDDkBBWNlpcqPy6n3sygP00THMPP0O3sFqy924eHqoDj3qSw79/meaZBJt9S5odPuFoskxjHuI4lM6BmK1Ql7p8Wo9/GhTOIoMz81orKPHRDleLjutwL4mb6NnhI5rfT/MxnHD6m82c4YYqiZC3XiTyJWVCkWkp7PK92OdRp6FA87rdKDMCAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADggEBAIYlezVU68TuEblkn06vM5dfzSmHKJOQgW61nDlLnyKrmSJtzKZLCAswTE2VyJHwKNdZgW15coJFINjWBLWcLr0/GjNV4u3Z+UL3NhBFQd5xuMtKsIhuoscKtyk0JHQXpBvHNmOUCobfQfOBQfTVC7kmyWdtlGztFUVxD28s6S5gMb1FEWWN68NOOJ3/ZhaTbUEM54yw8Hk8/f0L/Zn/7BYHUyWWA3KStAaYn89C/ZFF+952ark2VaKGIjBRQzgrJEIR8dI4r46I3DoEfzGPESKvQPvVLhOX84RG0PLPOtnRbHBVew1Nh3HE9kgCubkPKK+NPWE9IHZPoRmOTWBe+zU="),
	// 		FriendlyName: to.Ptr(""),
	// 		Issuer: to.Ptr("CN=Windows Azure Tools"),
	// 		ResourceID: to.Ptr[int64](5047111583862266000),
	// 		ServicePrincipalClientID: to.Ptr("a65f2051-51a3-4fa0-965e-76195bc73f4d"),
	// 		ServicePrincipalObjectID: to.Ptr("86d19f0e-82b7-49bf-a16e-1fbb4e5288d5"),
	// 		Subject: to.Ptr("CN=Windows Azure Tools"),
	// 		Thumbprint: to.Ptr("0526BD0123A52EABDB586AF0C080ABEEF3BB4240"),
	// 		ValidFrom: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-10T15:36:40.000Z"); return t}()),
	// 		ValidTo: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T15:46:40.000Z"); return t}()),
	// 	},
	// }
}
