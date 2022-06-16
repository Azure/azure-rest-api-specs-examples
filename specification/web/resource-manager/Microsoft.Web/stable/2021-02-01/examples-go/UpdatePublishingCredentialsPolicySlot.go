package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/UpdatePublishingCredentialsPolicySlot.json
func ExampleWebAppsClient_UpdateScmAllowedSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewWebAppsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateScmAllowedSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<slot>",
		armappservice.CsmPublishingCredentialsPoliciesEntity{
			Properties: &armappservice.CsmPublishingCredentialsPoliciesEntityProperties{
				Allow: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CsmPublishingCredentialsPoliciesEntity.ID: %s\n", *res.ID)
}
