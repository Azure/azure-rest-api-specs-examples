package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListCertificates.json
func ExampleCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificatesClient().NewListPager(&armappservice.CertificatesClientListOptions{Filter: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AppCertificateCollection = armappservice.AppCertificateCollection{
		// 	Value: []*armappservice.AppCertificate{
		// 		{
		// 			Name: to.Ptr("testc6282"),
		// 			Type: to.Ptr("Microsoft.Web/certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.AppCertificateProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
		// 				FriendlyName: to.Ptr(""),
		// 				HostNames: []*string{
		// 					to.Ptr("ServerCert")},
		// 					IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
		// 					Issuer: to.Ptr("CACert"),
		// 					SubjectName: to.Ptr("ServerCert"),
		// 					Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testc4912"),
		// 				Type: to.Ptr("Microsoft.Web/certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc4912"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armappservice.AppCertificateProperties{
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2040-12-31T23:59:59.000Z"); return t}()),
		// 					FriendlyName: to.Ptr(""),
		// 					HostNames: []*string{
		// 						to.Ptr("ServerCert2")},
		// 						IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-12-12T23:40:25.000Z"); return t}()),
		// 						Issuer: to.Ptr("CACert"),
		// 						SubjectName: to.Ptr("ServerCert2"),
		// 						Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B0490D5886EBFE"),
		// 					},
		// 			}},
		// 		}
	}
}
