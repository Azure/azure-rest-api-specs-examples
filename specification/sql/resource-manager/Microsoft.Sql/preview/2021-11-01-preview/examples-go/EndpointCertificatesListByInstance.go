package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/EndpointCertificatesListByInstance.json
func ExampleEndpointCertificatesClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.EndpointCertificateListResult = armsql.EndpointCertificateListResult{
		// 	Value: []*armsql.EndpointCertificate{
		// 		{
		// 			Name: to.Ptr("SERVICE_BROKER"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/endpointCertificates"),
		// 			ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/endpointCertificates/SERVICE_BROKER"),
		// 			Properties: &armsql.EndpointCertificateProperties{
		// 				PublicBlob: to.Ptr("0x308203B23082021AA003020102021034C597BA"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DATABASE_MIRRORING"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/endpointCertificates"),
		// 			ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/endpointCertificates/DATABASE_MIRRORING"),
		// 			Properties: &armsql.EndpointCertificateProperties{
		// 				PublicBlob: to.Ptr("0x308203B23082021AA003020102021034C597BA"),
		// 			},
		// 	}},
		// }
	}
}
