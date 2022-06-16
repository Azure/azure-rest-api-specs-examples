package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiIssueAttachment.json
func ExampleAPIIssueAttachmentClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIIssueAttachmentClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.GetEntityTag(ctx,
		"rg1",
		"apimService1",
		"57d2ef278aa04f0888cba3f3",
		"57d2ef278aa04f0ad01d6cdc",
		"57d2ef278aa04f0888cba3f3",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
