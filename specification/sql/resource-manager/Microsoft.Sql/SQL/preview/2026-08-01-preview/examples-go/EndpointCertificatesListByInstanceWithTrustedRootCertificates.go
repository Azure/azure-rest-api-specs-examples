package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/EndpointCertificatesListByInstanceWithTrustedRootCertificates.json
func ExampleEndpointCertificatesClient_NewListByInstancePager_getAListOfEndpointCertificatesIncludingTrustedRootCertificates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("38e0dc56-907f-45ba-a97c-74233baad471", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEndpointCertificatesClient().NewListByInstancePager("testrg", "testcl", nil)
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
		// page = armsql.EndpointCertificatesClientListByInstanceResponse{
		// 	EndpointCertificateListResult: armsql.EndpointCertificateListResult{
		// 		Value: []*armsql.EndpointCertificate{
		// 			{
		// 				Properties: &armsql.EndpointCertificateProperties{
		// 					PublicBlob: to.Ptr("0x308203B23082021AA003020102021034C597BA"),
		// 					TrustedRootCertificates: []*armsql.EndpointTrustedRootCertificateInfo{
		// 						{
		// 							PublicBlob: to.Ptr("0x3082038E30820276A0030201020210033AF1E6"),
		// 							Subject: to.Ptr("CN=DigiCert Global Root G2"),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/endpointCertificates/SERVICE_BROKER"),
		// 				Name: to.Ptr("SERVICE_BROKER"),
		// 				Type: to.Ptr("Microsoft.Sql/managedInstances/endpointCertificates"),
		// 			},
		// 			{
		// 				Properties: &armsql.EndpointCertificateProperties{
		// 					PublicBlob: to.Ptr("0x308203B23082021AA003020102021034C597BA"),
		// 					TrustedRootCertificates: []*armsql.EndpointTrustedRootCertificateInfo{
		// 						{
		// 							PublicBlob: to.Ptr("0x3082038E30820276A0030201020210033AF1E6"),
		// 							Subject: to.Ptr("CN=DigiCert Global Root G2"),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/endpointCertificates/DATABASE_MIRRORING"),
		// 				Name: to.Ptr("DATABASE_MIRRORING"),
		// 				Type: to.Ptr("Microsoft.Sql/managedInstances/endpointCertificates"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
