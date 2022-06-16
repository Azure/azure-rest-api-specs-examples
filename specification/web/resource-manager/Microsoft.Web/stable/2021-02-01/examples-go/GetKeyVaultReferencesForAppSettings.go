package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetKeyVaultReferencesForAppSettings.json
func ExampleWebAppsClient_GetAppSettingsKeyVaultReferences() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewWebAppsClient("<subscription-id>", cred, nil)
	pager := client.GetAppSettingsKeyVaultReferences("<resource-group-name>",
		"<name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("APIKVReference.ID: %s\n", *v.ID)
		}
	}
}
