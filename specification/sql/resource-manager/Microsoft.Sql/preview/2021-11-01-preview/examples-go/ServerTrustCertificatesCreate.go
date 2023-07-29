package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerTrustCertificatesCreate.json
func ExampleServerTrustCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerTrustCertificatesClient().BeginCreateOrUpdate(ctx, "testrg", "testcl", "customerCertificateName", armsql.ServerTrustCertificate{
		Properties: &armsql.ServerTrustCertificateProperties{
			PublicBlob: to.Ptr("308203AE30820296A0030201020210"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerTrustCertificate = armsql.ServerTrustCertificate{
	// 	Name: to.Ptr("customerCertificateName"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/serverTrustCertificates"),
	// 	ID: to.Ptr("/subscriptions/0574222d-5c7f-489c-a172-b3013eafab53/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/serverTrustCertificates/customerCertificateName"),
	// 	Properties: &armsql.ServerTrustCertificateProperties{
	// 		CertificateName: to.Ptr("customerCertificateName"),
	// 		Thumbprint: to.Ptr("33702D20EC86119985283"),
	// 	},
	// }
}
