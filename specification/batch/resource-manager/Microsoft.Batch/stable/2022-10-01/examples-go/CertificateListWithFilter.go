package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/CertificateListWithFilter.json
func ExampleCertificateClient_NewListByBatchAccountPager_listCertificatesFilterAndSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewCertificateClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.CertificateClientListByBatchAccountOptions{Maxresults: nil,
		Select: to.Ptr("properties/format,properties/provisioningState"),
		Filter: to.Ptr("properties/provisioningStateTransitionTime gt '2017-05-01' or properties/provisioningState eq 'Failed'"),
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
