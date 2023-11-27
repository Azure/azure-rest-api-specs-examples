package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactManifestListCredential.json
func ExampleArtifactManifestsClient_ListCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactManifestsClient().ListCredential(ctx, "rg", "TestPublisher", "TestArtifactStore", "TestArtifactManifestName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridnetwork.ArtifactManifestsClientListCredentialResponse{
	// 	                            ArtifactAccessCredentialClassification: &armhybridnetwork.AzureContainerRegistryScopedTokenCredential{
	// 		CredentialType: to.Ptr(armhybridnetwork.CredentialTypeAzureContainerRegistryScopedToken),
	// 		AcrServerURL: to.Ptr("test.azurecr.io"),
	// 		AcrToken: to.Ptr("dummytoken"),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-07T18:01:41.334Z"); return t}()),
	// 		Repositories: []*string{
	// 			to.Ptr("")},
	// 			Username: to.Ptr("dummyuser"),
	// 		},
	// 		                        }
}
