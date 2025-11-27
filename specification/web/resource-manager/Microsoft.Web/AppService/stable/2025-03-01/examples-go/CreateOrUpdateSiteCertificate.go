package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/CreateOrUpdateSiteCertificate.json
func ExampleSiteCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSiteCertificatesClient().CreateOrUpdate(ctx, "testrg123", "testSiteName", "testc6282", armappservice.AppCertificate{
		Location: to.Ptr("East US"),
		Properties: &armappservice.AppCertificateProperties{
			HostNames: []*string{
				to.Ptr("ServerCert")},
			Password: to.Ptr("<password>"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppCertificate = armappservice.AppCertificate{
	// 	Name: to.Ptr("testc6282"),
	// 	Type: to.Ptr("Microsoft.Web/sites/certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/testSiteName/certificates/testc6282"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.AppCertificateProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		FriendlyName: to.Ptr(""),
	// 		HostNames: []*string{
	// 			to.Ptr("ServerCert")},
	// 			IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
	// 			Issuer: to.Ptr("CACert"),
	// 			SubjectName: to.Ptr("ServerCert"),
	// 			Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
	// 		},
	// 	}
}
