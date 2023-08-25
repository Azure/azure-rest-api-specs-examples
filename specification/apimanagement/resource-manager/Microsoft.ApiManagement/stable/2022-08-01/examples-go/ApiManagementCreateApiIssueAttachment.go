package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiIssueAttachment.json
func ExampleAPIIssueAttachmentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIIssueAttachmentClient().CreateOrUpdate(ctx, "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "57d2ef278aa04f0ad01d6cdc", "57d2ef278aa04f0888cba3f3", armapimanagement.IssueAttachmentContract{
		Properties: &armapimanagement.IssueAttachmentContractProperties{
			Content:       to.Ptr("IEJhc2U2NA=="),
			ContentFormat: to.Ptr("image/jpeg"),
			Title:         to.Ptr("Issue attachment."),
		},
	}, &armapimanagement.APIIssueAttachmentClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IssueAttachmentContract = armapimanagement.IssueAttachmentContract{
	// 	Name: to.Ptr("57d2ef278aa04f0888cba3f3"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/issues/attachments"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d1f7558aa04f15146d9d8a/issues/57d2ef278aa04f0ad01d6cdc/attachments/57d2ef278aa04f0888cba3f3"),
	// 	Properties: &armapimanagement.IssueAttachmentContractProperties{
	// 		Content: to.Ptr("https://.../image.jpg"),
	// 		ContentFormat: to.Ptr("link"),
	// 		Title: to.Ptr("Issue attachment."),
	// 	},
	// }
}
