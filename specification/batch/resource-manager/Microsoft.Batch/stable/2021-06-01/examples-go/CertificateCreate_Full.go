package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateCreate_Full.json
func ExampleCertificateClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		armbatch.CertificateCreateOrUpdateParameters{
			Properties: &armbatch.CertificateCreateOrUpdateProperties{
				Format:              armbatch.CertificateFormatPfx.ToPtr(),
				Thumbprint:          to.StringPtr("<thumbprint>"),
				ThumbprintAlgorithm: to.StringPtr("<thumbprint-algorithm>"),
				Data:                to.StringPtr("<data>"),
				Password:            to.StringPtr("<password>"),
			},
		},
		&armbatch.CertificateClientCreateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientCreateResult)
}
