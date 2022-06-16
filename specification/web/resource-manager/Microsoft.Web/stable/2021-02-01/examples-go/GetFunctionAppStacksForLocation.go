package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetFunctionAppStacksForLocation.json
func ExampleProviderClient_GetFunctionAppStacksForLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.GetFunctionAppStacksForLocation("<location>",
		&armappservice.ProviderGetFunctionAppStacksForLocationOptions{StackOsType: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("FunctionAppStack.ID: %s\n", *v.ID)
		}
	}
}
