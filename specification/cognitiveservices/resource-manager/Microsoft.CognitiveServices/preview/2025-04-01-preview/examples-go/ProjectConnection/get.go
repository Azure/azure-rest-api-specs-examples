package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5802c95f18bfba1003be50e545d07f8bb679c857/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/ProjectConnection/get.json
func ExampleProjectConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectConnectionClient().Get(ctx, "resourceGroup-1", "account-1", "project-1", "connection-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionPropertiesV2BasicResource = armcognitiveservices.ConnectionPropertiesV2BasicResource{
	// 	Name: to.Ptr("connection-1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/connections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/projects/project-1/connections/connection-1"),
	// 	Properties: &armcognitiveservices.NoneAuthTypeConnectionProperties{
	// 		AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypeNone),
	// 		Category: to.Ptr(armcognitiveservices.ConnectionCategoryContainerRegistry),
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T14:30:00.000Z"); return t}()),
	// 		Target: to.Ptr("[tartget url]"),
	// 	},
	// }
}
