package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/PatchAppServicePlan.json
func ExampleAppServicePlansClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.AppServicePlanPatchResource{
			ProxyOnlyResource: armappservice.ProxyOnlyResource{
				Kind: to.StringPtr("<kind>"),
			},
			Properties: &armappservice.AppServicePlanPatchResourceProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AppServicePlan.ID: %s\n", *res.ID)
}
