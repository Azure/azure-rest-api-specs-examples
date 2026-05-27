package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListApiIssues.json
func ExampleAPIIssueClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIIssueClient().NewListByServicePager("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", nil)
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
		// page = armapimanagement.APIIssueClientListByServiceResponse{
		// 	IssueCollection: armapimanagement.IssueCollection{
		// 		Count: to.Ptr[int64](1),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.IssueContract{
		// 			{
		// 				Name: to.Ptr("57d2ef278aa04f0ad01d6cdc"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/apis/issues"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d1f7558aa04f15146d9d8a/issues/57d2ef278aa04f0ad01d6cdc"),
		// 				Properties: &armapimanagement.IssueContractProperties{
		// 					Description: to.Ptr("New API issue description"),
		// 					APIID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d1f7558aa04f15146d9d8a"),
		// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T22:21:20.467Z"); return t}()),
		// 					State: to.Ptr(armapimanagement.StateOpen),
		// 					Title: to.Ptr("New API issue"),
		// 					UserID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
