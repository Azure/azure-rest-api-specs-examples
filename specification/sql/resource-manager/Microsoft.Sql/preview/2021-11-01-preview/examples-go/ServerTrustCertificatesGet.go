package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerTrustCertificatesGet.json
func ExampleServerTrustCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerTrustCertificatesClient().Get(ctx, "testrg", "testcl", "customerCertificateName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerTrustCertificate = armsql.ServerTrustCertificate{
	// 	Name: to.Ptr("customerCertificateName"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/serverTrustCertificates"),
	// 	ID: to.Ptr("/subscriptions/38e0dc56-907f-45ba-a97c-74233baad471/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/serverTrustCertificates/customerCertificateName"),
	// 	Properties: &armsql.ServerTrustCertificateProperties{
	// 		CertificateName: to.Ptr("customerCertificateName"),
	// 		Thumbprint: to.Ptr("57CFA9CF16F2FB2775AF059A95C6D5B897DA2C05"),
	// 	},
	// }
}
