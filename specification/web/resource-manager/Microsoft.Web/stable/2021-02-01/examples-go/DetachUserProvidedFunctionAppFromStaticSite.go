package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/DetachUserProvidedFunctionAppFromStaticSite.json
func ExampleStaticSitesClient_DetachUserProvidedFunctionAppFromStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	_, err = client.DetachUserProvidedFunctionAppFromStaticSite(ctx,
		"<resource-group-name>",
		"<name>",
		"<function-app-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
