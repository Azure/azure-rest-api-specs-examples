package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPortalRevisions.json
func ExamplePortalRevisionClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPortalRevisionClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.PortalRevisionClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.PortalRevisionCollection = armapimanagement.PortalRevisionCollection{
		// 	Value: []*armapimanagement.PortalRevisionContract{
		// 		{
		// 			Name: to.Ptr("20201112000000"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalRevisions/20201112000000"),
		// 			Properties: &armapimanagement.PortalRevisionContractProperties{
		// 				Description: to.Ptr("portal revision"),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:10:09.673Z"); return t}()),
		// 				IsCurrent: to.Ptr(false),
		// 				Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
		// 				UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:12:41.46Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20201112101010"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalRevisions/20201112101010"),
		// 			Properties: &armapimanagement.PortalRevisionContractProperties{
		// 				Description: to.Ptr("portal revision 1"),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:51:36.47Z"); return t}()),
		// 				IsCurrent: to.Ptr(true),
		// 				Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
		// 				UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:52:00.097Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
