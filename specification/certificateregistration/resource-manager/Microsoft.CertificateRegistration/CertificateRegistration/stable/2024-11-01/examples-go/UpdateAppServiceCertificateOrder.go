package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/stable/2024-11-01/examples/UpdateAppServiceCertificateOrder.json
func ExampleCertificateOrdersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateOrdersClient().Update(ctx, "testrg123", "SampleCertificateOrderName", armappservice.CertificateOrderPatchResource{
		Properties: &armappservice.CertificateOrderPatchResourceProperties{
			AutoRenew: to.Ptr(true),
			Certificates: map[string]*armappservice.Certificate{
				"SampleCertName1": {
					KeyVaultID:         to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: to.Ptr("SampleSecretName1"),
				},
				"SampleCertName2": {
					KeyVaultID:         to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: to.Ptr("SampleSecretName2"),
				},
			},
			DistinguishedName: to.Ptr("CN=SampleCustomDomain.com"),
			KeySize:           to.Ptr[int32](2048),
			ProductType:       to.Ptr(armappservice.CertificateProductTypeStandardDomainValidatedSSL),
			ValidityInYears:   to.Ptr[int32](2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateOrder = armappservice.CertificateOrder{
	// 	Name: to.Ptr("SampleCertificateOrderName"),
	// 	Type: to.Ptr("Microsoft.CertificateRegistration/certificateOrders"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName"),
	// 	Location: to.Ptr("Global"),
	// 	Properties: &armappservice.CertificateOrderProperties{
	// 		AppServiceCertificateNotRenewableReasons: []*armappservice.ResourceNotRenewableReason{
	// 			to.Ptr(armappservice.ResourceNotRenewableReasonExpirationNotInRenewalTimeRange)},
	// 			AutoRenew: to.Ptr(true),
	// 			Certificates: map[string]*armappservice.Certificate{
	// 				"SampleCertName1": &armappservice.Certificate{
	// 					KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
	// 					KeyVaultSecretName: to.Ptr("SampleSecretName1"),
	// 					ProvisioningState: to.Ptr(armappservice.KeyVaultSecretStatusSucceeded),
	// 				},
	// 				"SampleCertName2": &armappservice.Certificate{
	// 					KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
	// 					KeyVaultSecretName: to.Ptr("SampleSecretName2"),
	// 					ProvisioningState: to.Ptr(armappservice.KeyVaultSecretStatusSucceeded),
	// 				},
	// 			},
	// 			Contact: &armappservice.CertificateOrderContact{
	// 				Email: to.Ptr("contact@email.com"),
	// 				NameFirst: to.Ptr("SampleFirstName"),
	// 				NameLast: to.Ptr("SampleLastName"),
	// 				Phone: to.Ptr("1-224-550-332"),
	// 			},
	// 			Csr: to.Ptr("xdf!x|"),
	// 			DistinguishedName: to.Ptr("CN=SampleCustomDomain.com"),
	// 			DomainVerificationToken: to.Ptr("awe4163Br6d32r"),
	// 			ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-12T23:40:25.000Z"); return t}()),
	// 			Intermediate: &armappservice.CertificateDetails{
	// 				Issuer: to.Ptr("CACert"),
	// 				NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-12T23:40:25.000Z"); return t}()),
	// 				NotBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-12T23:40:25.000Z"); return t}()),
	// 				SerialNumber: to.Ptr("k45kad0024"),
	// 				SignatureAlgorithm: to.Ptr("SHA"),
	// 				Subject: to.Ptr("CN=IntermediateCertSubjectName"),
	// 				Thumbprint: to.Ptr("RF3D7411A44163B6D32B3343559B03E175886EBFE"),
	// 				Version: to.Ptr[int32](5589220),
	// 			},
	// 			IsPrivateKeyExternal: to.Ptr(false),
	// 			KeySize: to.Ptr[int32](2048),
	// 			LastCertificateIssuanceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-12T23:40:25.000Z"); return t}()),
	// 			NextAutoRenewalTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-12T23:40:25.000Z"); return t}()),
	// 			ProductType: to.Ptr(armappservice.CertificateProductTypeStandardDomainValidatedSSL),
	// 			ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 			Root: &armappservice.CertificateDetails{
	// 				Issuer: to.Ptr("CACert"),
	// 				NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-12T23:40:25.000Z"); return t}()),
	// 				NotBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-12T23:40:25.000Z"); return t}()),
	// 				SerialNumber: to.Ptr("jkskd890ajdd00124"),
	// 				SignatureAlgorithm: to.Ptr("SHA"),
	// 				Subject: to.Ptr("CN=RootCertSubjectName"),
	// 				Thumbprint: to.Ptr("ZD3D7411A44163B6D32B3343559B03E175886EBFE"),
	// 				Version: to.Ptr[int32](3430921),
	// 			},
	// 			SerialNumber: to.Ptr("12324325325"),
	// 			SignedCertificate: &armappservice.CertificateDetails{
	// 				Issuer: to.Ptr("CACert"),
	// 				NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-12T23:40:25.000Z"); return t}()),
	// 				NotBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-12T23:40:25.000Z"); return t}()),
	// 				SerialNumber: to.Ptr("sds12424k"),
	// 				SignatureAlgorithm: to.Ptr("SHA"),
	// 				Subject: to.Ptr("CN=CertSubjectName"),
	// 				Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
	// 				Version: to.Ptr[int32](1232490),
	// 			},
	// 			Status: to.Ptr(armappservice.CertificateOrderStatusIssued),
	// 			ValidityInYears: to.Ptr[int32](2),
	// 		},
	// 	}
}
