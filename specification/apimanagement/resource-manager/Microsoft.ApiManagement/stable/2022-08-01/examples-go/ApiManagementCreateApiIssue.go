package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiIssue.json
func ExampleAPIIssueClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIIssueClient().CreateOrUpdate(ctx, "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "57d2ef278aa04f0ad01d6cdc", armapimanagement.IssueContract{
		Properties: &armapimanagement.IssueContractProperties{
			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T22:21:20.467Z"); return t }()),
			State:       to.Ptr(armapimanagement.StateOpen),
			Description: to.Ptr("New API issue description"),
			Title:       to.Ptr("New API issue"),
			UserID:      to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		},
	}, &armapimanagement.APIIssueClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IssueContract = armapimanagement.IssueContract{
	// 	Name: to.Ptr("57d2ef278aa04f0ad01d6cdc"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/issues"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d1f7558aa04f15146d9d8a/issues/57d2ef278aa04f0ad01d6cdc"),
	// 	Properties: &armapimanagement.IssueContractProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T22:21:20.467Z"); return t}()),
	// 		State: to.Ptr(armapimanagement.StateOpen),
	// 		Description: to.Ptr("New API issue description"),
	// 		Title: to.Ptr("New API issue"),
	// 		UserID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
	// 	},
	// }
}
