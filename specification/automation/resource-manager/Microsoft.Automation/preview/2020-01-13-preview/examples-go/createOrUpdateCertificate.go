package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateCertificate.json
func ExampleCertificateClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateClient().CreateOrUpdate(ctx, "rg", "myAutomationAccount18", "testCert", armautomation.CertificateCreateOrUpdateParameters{
		Name: to.Ptr("testCert"),
		Properties: &armautomation.CertificateCreateOrUpdateProperties{
			Description:  to.Ptr("Sample Cert"),
			Base64Value:  to.Ptr("base 64 value of cert"),
			IsExportable: to.Ptr(false),
			Thumbprint:   to.Ptr("thumbprint of cert"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Certificate = armautomation.Certificate{
	// 	Name: to.Ptr("testCert"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/certificates/testCert"),
	// 	Properties: &armautomation.CertificateProperties{
	// 		Description: to.Ptr("sample certificate. Description updated"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T17:26:43.337+00:00"); return t}()),
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T17:25:45+00:00"); return t}()),
	// 		IsExportable: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T17:28:55.01+00:00"); return t}()),
	// 		Thumbprint: to.Ptr("thumbprint of cert"),
	// 	},
	// }
}
