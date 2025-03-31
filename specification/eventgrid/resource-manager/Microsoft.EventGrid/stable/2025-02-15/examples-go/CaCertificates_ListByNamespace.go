package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/CaCertificates_ListByNamespace.json
func ExampleCaCertificatesClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCaCertificatesClient().NewListByNamespacePager("examplerg", "namespace123", &armeventgrid.CaCertificatesClientListByNamespaceOptions{Filter: nil,
		Top: nil,
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
		// page.CaCertificatesListResult = armeventgrid.CaCertificatesListResult{
		// 	Value: []*armeventgrid.CaCertificate{
		// 		{
		// 			Name: to.Ptr("exampleCACertificateName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/caCertificates"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/caCertificates/exampleCACertificateName1"),
		// 			Properties: &armeventgrid.CaCertificateProperties{
		// 				Description: to.Ptr("This is a test Root certificate"),
		// 				EncodedCertificate: to.Ptr("base64EncodePemFormattedCertificateString"),
		// 				ExpiryTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-12T23:06:43.000Z"); return t}()),
		// 				IssueTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T23:06:43.000Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armeventgrid.CaCertificateProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
