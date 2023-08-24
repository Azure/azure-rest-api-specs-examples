package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetNamedValueWithKeyVault.json
func ExampleNamedValueClient_Get_apiManagementGetNamedValueWithKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamedValueClient().Get(ctx, "rg1", "apimService1", "testprop6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop6"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/namedValues"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop6"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(true),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
	// 				SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.8024882Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 	}
}
