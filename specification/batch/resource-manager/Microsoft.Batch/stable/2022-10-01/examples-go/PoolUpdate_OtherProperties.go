package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolUpdate_OtherProperties.json
func ExamplePoolClient_Update_updatePoolOtherProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
				},
				{
					ID:      to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
					Version: to.Ptr("1.0"),
				}},
			Certificates: []*armbatch.CertificateReference{
				{
					ID:            to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
					StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
					StoreName:     to.Ptr("MY"),
				}},
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("key1"),
					Value: to.Ptr("value1"),
				}},
			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeSimplified),
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
