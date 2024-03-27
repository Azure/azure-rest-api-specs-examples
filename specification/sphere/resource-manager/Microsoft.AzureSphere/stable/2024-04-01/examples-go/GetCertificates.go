package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/09c37754dac91874ff689ed1e60effb4268c8669/specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCertificates.json
func ExampleCertificatesClient_NewListByCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificatesClient().NewListByCatalogPager("MyResourceGroup1", "MyCatalog1", &armsphere.CertificatesClientListByCatalogOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
	})
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
		// page.CertificateListResult = armsphere.CertificateListResult{
		// 	Value: []*armsphere.Certificate{
		// 	},
		// }
	}
}
