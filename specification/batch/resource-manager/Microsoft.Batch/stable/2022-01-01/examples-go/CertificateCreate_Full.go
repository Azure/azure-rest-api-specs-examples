package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateCreate_Full.json
func ExampleCertificateClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewCertificateClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"default-azurebatch-japaneast",
		"sampleacct",
		"sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e",
		armbatch.CertificateCreateOrUpdateParameters{
			Properties: &armbatch.CertificateCreateOrUpdateProperties{
				Format:              to.Ptr(armbatch.CertificateFormatPfx),
				Thumbprint:          to.Ptr("0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
				ThumbprintAlgorithm: to.Ptr("sha1"),
				Data:                to.Ptr("MIIJsgIBAzCCCW4GCSqGSIb3DQE..."),
				Password:            to.Ptr("<ExamplePassword>"),
			},
		},
		&armbatch.CertificateClientCreateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
