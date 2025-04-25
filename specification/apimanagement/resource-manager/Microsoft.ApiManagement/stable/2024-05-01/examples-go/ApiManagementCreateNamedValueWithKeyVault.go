package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateNamedValueWithKeyVault.json
func ExampleNamedValueClient_BeginCreateOrUpdate_apiManagementCreateNamedValueWithKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamedValueClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "testprop6", armapimanagement.NamedValueCreateContract{
		Properties: &armapimanagement.NamedValueCreateContractProperties{
			Secret: to.Ptr(true),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar")},
			DisplayName: to.Ptr("prop6namekv"),
			KeyVault: &armapimanagement.KeyVaultContractCreateProperties{
				IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
				SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
			},
		},
	}, &armapimanagement.NamedValueClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop6"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop6"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(true),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
	// 				SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.802Z"); return t}()),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}
