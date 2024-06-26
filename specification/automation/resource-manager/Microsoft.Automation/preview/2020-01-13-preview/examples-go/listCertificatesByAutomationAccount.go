package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCertificatesByAutomationAccount.json
func ExampleCertificateClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", nil)
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
		// page.CertificateListResult = armautomation.CertificateListResult{
		// 	Value: []*armautomation.Certificate{
		// 		{
		// 			Name: to.Ptr("testCert"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/certificates/testCert"),
		// 			Properties: &armautomation.CertificateProperties{
		// 				Description: to.Ptr("Sample Cert"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T17:26:43.337Z"); return t}()),
		// 				ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T17:25:45.000Z"); return t}()),
		// 				IsExportable: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T17:28:55.010Z"); return t}()),
		// 				Thumbprint: to.Ptr("thumbprint of cert"),
		// 			},
		// 	}},
		// }
	}
}
