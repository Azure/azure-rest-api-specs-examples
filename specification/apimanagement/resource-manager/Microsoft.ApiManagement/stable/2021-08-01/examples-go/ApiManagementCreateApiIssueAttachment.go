package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiIssueAttachment.json
func ExampleAPIIssueAttachmentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIIssueAttachmentClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"57d1f7558aa04f15146d9d8a",
		"57d2ef278aa04f0ad01d6cdc",
		"57d2ef278aa04f0888cba3f3",
		armapimanagement.IssueAttachmentContract{
			Properties: &armapimanagement.IssueAttachmentContractProperties{
				Content:       to.Ptr("IEJhc2U2NA=="),
				ContentFormat: to.Ptr("image/jpeg"),
				Title:         to.Ptr("Issue attachment."),
			},
		},
		&armapimanagement.APIIssueAttachmentClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
