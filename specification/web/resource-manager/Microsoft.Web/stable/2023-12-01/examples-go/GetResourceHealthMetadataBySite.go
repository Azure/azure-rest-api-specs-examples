package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetResourceHealthMetadataBySite.json
func ExampleResourceHealthMetadataClient_GetBySite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceHealthMetadataClient().GetBySite(ctx, "Default-Web-NorthCentralUS", "newsiteinnewASE-NCUS", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceHealthMetadata = armappservice.ResourceHealthMetadata{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Web/sites/resourceHealthMetadata"),
	// 	ID: to.Ptr("/subscriptions/4adb32ad-8327-4cbb-b775-b84b4465bb38/resourceGroups/Default-Web-NorthCentralUS/providers/Microsoft.Web/sites/newsiteinnewASE-NCUS/resourceHealthMetadata/default"),
	// 	Properties: &armappservice.ResourceHealthMetadataProperties{
	// 		Category: to.Ptr("Shared"),
	// 		SignalAvailability: to.Ptr(true),
	// 	},
	// }
}
