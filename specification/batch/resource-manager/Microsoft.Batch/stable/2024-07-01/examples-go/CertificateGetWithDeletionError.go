package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/CertificateGetWithDeletionError.json
func ExampleCertificateClient_Get_getCertificateWithDeletionError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e", nil)
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
	// 		DeleteCertificateError: &armbatch.DeleteCertificateError{
	// 			Code: to.Ptr("NodesReferencingCertificate"),
	// 			Message: to.Ptr("The specified certificate is being used by the below mentioned node(s)\nRequestId:2dc78afc-b15b-42d2-8c85-39cb61a0799e\nTime:2017-08-28T10:22:52.8633406Z"),
	// 			Target: to.Ptr("BatchAccount"),
	// 			Details: []*armbatch.DeleteCertificateError{
	// 				{
	// 					Code: to.Ptr("Nodes"),
	// 					Message: to.Ptr("node1, node3"),
	// 			}},
	// 		},
	// 		PreviousProvisioningState: to.Ptr(armbatch.CertificateProvisioningStateDeleting),
	// 		PreviousProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-21T00:15:25.562Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbatch.CertificateProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-21T01:47:38.442Z"); return t}()),
	// 		PublicData: to.Ptr("MIICrjCCAZagAwI..."),
	// 	},
	// }
}
