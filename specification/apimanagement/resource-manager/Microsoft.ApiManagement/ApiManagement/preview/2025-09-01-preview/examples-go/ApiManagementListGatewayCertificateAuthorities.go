package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListGatewayCertificateAuthorities.json
func ExampleGatewayCertificateAuthorityClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewayCertificateAuthorityClient().NewListByServicePager("rg1", "apimService1", "gw1", nil)
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
		// page = armapimanagement.GatewayCertificateAuthorityClientListByServiceResponse{
		// 	GatewayCertificateAuthorityCollection: armapimanagement.GatewayCertificateAuthorityCollection{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.GatewayCertificateAuthorityContract{
		// 			{
		// 				Name: to.Ptr("cert1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/gateways/certificateAuthorities"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1/certificateAuthorities/cert1"),
		// 				Properties: &armapimanagement.GatewayCertificateAuthorityContractProperties{
		// 					IsTrusted: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("cert2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/gateways/certificateAuthorities"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1/certificateAuthorities/cert2"),
		// 				Properties: &armapimanagement.GatewayCertificateAuthorityContractProperties{
		// 					IsTrusted: to.Ptr(true),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
