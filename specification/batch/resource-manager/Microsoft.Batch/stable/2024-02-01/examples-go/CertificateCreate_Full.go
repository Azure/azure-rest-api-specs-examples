package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/CertificateCreate_Full.json
func ExampleCertificateClient_Create_createCertificateFull() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateClient().Create(ctx, "default-azurebatch-japaneast", "sampleacct", "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e", armbatch.CertificateCreateOrUpdateParameters{
		Properties: &armbatch.CertificateCreateOrUpdateProperties{
			Format:              to.Ptr(armbatch.CertificateFormatPfx),
			Thumbprint:          to.Ptr("0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
			ThumbprintAlgorithm: to.Ptr("sha1"),
			Data:                to.Ptr("MIIJsgIBAzCCCW4GCSqGSIb3DQE..."),
			Password:            to.Ptr("<ExamplePassword>"),
		},
	}, &armbatch.CertificateClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Certificate = armbatch.Certificate{
	// 	Name: to.Ptr("sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/certificates"),
	// 	Etag: to.Ptr("W/\"0x8D4EDD5118668F7\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/samplecct/certificates/sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
	// 	Properties: &armbatch.CertificateProperties{
	// 		Format: to.Ptr(armbatch.CertificateFormatPfx),
	// 		Thumbprint: to.Ptr("0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
	// 		ThumbprintAlgorithm: to.Ptr("sha1"),
	// 		ProvisioningState: to.Ptr(armbatch.CertificateProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-21T01:47:38.442Z"); return t}()),
	// 		PublicData: to.Ptr("MIICrjCCAZagAwI..."),
	// 	},
	// }
}
