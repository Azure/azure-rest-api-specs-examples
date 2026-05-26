package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementRefreshNamedValue.json
func ExampleNamedValueClient_BeginRefreshSecret() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamedValueClient().BeginRefreshSecret(ctx, "rg1", "apimService1", "testprop2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.NamedValueClientRefreshSecretResponse{
	// 	NamedValueContract: armapimanagement.NamedValueContract{
	// 		Name: to.Ptr("testprop6"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/namedValues"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop6"),
	// 		Properties: &armapimanagement.NamedValueContractProperties{
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.8024882Z"); return t}()),
	// 				},
	// 				SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
	// 			},
	// 			Secret: to.Ptr(true),
	// 			Tags: []*string{
	// 				to.Ptr("foo"),
	// 				to.Ptr("bar"),
	// 			},
	// 		},
	// 	},
	// }
}
