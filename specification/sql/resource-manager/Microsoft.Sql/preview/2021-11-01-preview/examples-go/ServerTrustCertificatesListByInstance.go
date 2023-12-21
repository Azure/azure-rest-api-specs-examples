package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerTrustCertificatesListByInstance.json
func ExampleServerTrustCertificatesClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerTrustCertificatesClient().NewListByInstancePager("testrg", "testcl", nil)
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
		// page.ServerTrustCertificatesListResult = armsql.ServerTrustCertificatesListResult{
		// 	Value: []*armsql.ServerTrustCertificate{
		// 		{
		// 			Name: to.Ptr("customerCertificate1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/serverTrustCertificates"),
		// 			ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/serverTrustCertificates/customerCertificate1"),
		// 			Properties: &armsql.ServerTrustCertificateProperties{
		// 				CertificateName: to.Ptr("customerCertificate1"),
		// 				Thumbprint: to.Ptr("33702D20EC861199852837AE6BD1A71544B681E2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("customerCertificate2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/serverTrustCertificates"),
		// 			ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/serverTrustCertificates/customerCertificate2"),
		// 			Properties: &armsql.ServerTrustCertificateProperties{
		// 				CertificateName: to.Ptr("customerCertificate2"),
		// 				Thumbprint: to.Ptr("57CFA9CF16F2FB2775AF059A95C6D5B897DA2C05"),
		// 			},
		// 	}},
		// }
	}
}
