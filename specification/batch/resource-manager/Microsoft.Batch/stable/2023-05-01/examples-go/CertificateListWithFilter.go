package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/CertificateListWithFilter.json
func ExampleCertificateClient_NewListByBatchAccountPager_listCertificatesFilterAndSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.CertificateClientListByBatchAccountOptions{Maxresults: nil,
		Select: to.Ptr("properties/format,properties/provisioningState"),
		Filter: to.Ptr("properties/provisioningStateTransitionTime gt '2017-05-01' or properties/provisioningState eq 'Failed'"),
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
		// page.ListCertificatesResult = armbatch.ListCertificatesResult{
		// 	Value: []*armbatch.Certificate{
		// 		{
		// 			Name: to.Ptr("sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/certificates"),
		// 			Etag: to.Ptr("W/\"0x8D4EDD5118668F7\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/samplecct/certificates/sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
		// 			Properties: &armbatch.CertificateProperties{
		// 				Format: to.Ptr(armbatch.CertificateFormatPfx),
		// 				ProvisioningState: to.Ptr(armbatch.CertificateProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sha1-aeb228ffb0bf67a793d61dce263ebd16949f15a1"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/certificates"),
		// 			Etag: to.Ptr("W/\"0x8D4EDD5118572E0\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/samplecct/certificates/sha1-aeb228ffb0bf67a793d61dce263ebd16949f15a1"),
		// 			Properties: &armbatch.CertificateProperties{
		// 				Format: to.Ptr(armbatch.CertificateFormatCer),
		// 				ProvisioningState: to.Ptr(armbatch.CertificateProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
