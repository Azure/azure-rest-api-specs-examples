package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/CaCertificates_Get.json
func ExampleCaCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCaCertificatesClient().Get(ctx, "examplerg", "exampleNamespaceName1", "exampleCACertificateName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CaCertificate = armeventgrid.CaCertificate{
	// 	Name: to.Ptr("exampleCACertificateName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/caCertificates"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/caCertificates/exampleCACertificateName1"),
	// 	Properties: &armeventgrid.CaCertificateProperties{
	// 		Description: to.Ptr("This is a test Root certificate"),
	// 		EncodedCertificate: to.Ptr("base64EncodePemFormattedCertificateString"),
	// 		ExpiryTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-12T23:06:43+00:00"); return t}()),
	// 		IssueTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T23:06:43+00:00"); return t}()),
	// 		ProvisioningState: to.Ptr(armeventgrid.CaCertificateProvisioningStateSucceeded),
	// 	},
	// }
}
