package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspaceNamedValues.json
func ExampleWorkspaceNamedValueClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceNamedValueClient().NewListByServicePager("rg1", "apimService1", "wks1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armapimanagement.WorkspaceNamedValueClientListByServiceResponse{
		// 	NamedValueCollection: armapimanagement.NamedValueCollection{
		// 		Count: to.Ptr[int64](2),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.NamedValueContract{
		// 			{
		// 				Name: to.Ptr("592f1174cc83890dc4f32686"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/592f1174cc83890dc4f32686"),
		// 				Properties: &armapimanagement.NamedValueContractProperties{
		// 					DisplayName: to.Ptr("Logger-Credentials-592f1174cc83890dc4f32687"),
		// 					Secret: to.Ptr(false),
		// 					Value: to.Ptr("propValue"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testprop6"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop6"),
		// 				Properties: &armapimanagement.NamedValueContractProperties{
		// 					DisplayName: to.Ptr("prop6namekv"),
		// 					KeyVault: &armapimanagement.KeyVaultContractProperties{
		// 						IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
		// 						LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
		// 							Code: to.Ptr("Success"),
		// 							TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.8024882Z"); return t}()),
		// 						},
		// 						SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
		// 					},
		// 					Secret: to.Ptr(true),
		// 					Tags: []*string{
		// 						to.Ptr("foo"),
		// 						to.Ptr("bar"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
