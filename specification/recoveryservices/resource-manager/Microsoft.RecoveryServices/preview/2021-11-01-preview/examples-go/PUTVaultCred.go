package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/PUTVaultCred.json
func ExampleVaultCertificatesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewVaultCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<certificate-name>",
		armrecoveryservices.CertificateRequest{
			Properties: &armrecoveryservices.RawCertificateData{
				AuthType:    armrecoveryservices.AuthType("AAD").ToPtr(),
				Certificate: []byte("MTTC3TCCAcWgAwIBAgIQEj9h+ZLlXK9KrqZX9UkAnzANBgkqhkiG9w0BAQUFADAeMRwwGgYDVQQDExNXaW5kb3dzIEF6dXJlIFRvb2xzMB4XDTE3MTIxODA5MTc1M1oXDTE3MTIyMzA5Mjc1M1owHjEcMBoGA1UEAxMTV2luZG93cyBBenVyZSBUb29sczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK773/eZZ69RbZZAT05r9MjUxu9y1L1Pn1EgPk62IPJyHlO3OZA922eSBahhP4bgmFljN4LVReqQ5eT/wqO0Zhc+yFkUy4U4RdbQLeUZt2W7yy9XLXgVvqeYDgsjg/QhHetgHArQBW+tlQq5+zPdU7zchI4rbShSJrWhLrZFWiOyFPsuAE4joUQHNlRifdCTsBGKk8HRCY3j1S3c4bfEn3zxlrvrXXssRuW5mJM95rMk0tskoRxXSCi6i9bnlki2Cs9mpVMmBFeofs41KwzlWU0TgpdD8s1QEdvfGB5NbByfetPX7MfJaTBeHZEGbv/Iq8l72u8sPBoOhcaH7qDE/mECAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADggEBAILfgHluye1Q+WelhgWhpBBdIq2C0btfV8eFsZaTlBUrM0fwpxQSlAWc2oYHVMQI4A5iUjbDOY35O4yc+TnWKDBKf+laqDP+yos4aiUPuadGUZfvDk7kuw7xeECs64JpHAIEKdRHFW9rD3gwG+nIWaDnEL/7rTyhL3kXrRW2MSUAL8g3GX8Z45c+MQY0jmASIqWdhGn1vpAGyA9mKkzsqg7FXjg8GZb24tGl5Ky85+ip4dkBfXinDD8WwaGyjhGGK97ErvNmN36qly/H0H1Qngiovg1FbHDmkcFO5QclnEJsFFmcO2CcHp5Fqh2wXn5O1cQaxCIRTpQ/uXRpDjl2wKs="),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VaultCertificatesClientCreateResult)
}
