package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/KubeEnvironments_Update.json
func ExampleKubeEnvironmentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewKubeEnvironmentsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.KubeEnvironmentPatchResource{
			Properties: &armappservice.KubeEnvironmentPatchResourceProperties{
				StaticIP: to.StringPtr("<static-ip>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("KubeEnvironment.ID: %s\n", *res.ID)
}
