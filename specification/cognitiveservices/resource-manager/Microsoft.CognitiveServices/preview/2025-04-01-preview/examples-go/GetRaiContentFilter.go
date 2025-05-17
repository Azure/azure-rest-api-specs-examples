package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5802c95f18bfba1003be50e545d07f8bb679c857/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/GetRaiContentFilter.json
func ExampleRaiContentFiltersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiContentFiltersClient().Get(ctx, "WestUS", "IndirectAttack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RaiContentFilter = armcognitiveservices.RaiContentFilter{
	// 	Name: to.Ptr("IndirectAttack"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/locations/raiContentFilters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CognitiveServices/locations/WestUS/raiContentFilters/IndirectAttack"),
	// 	Properties: &armcognitiveservices.RaiContentFilterProperties{
	// 		Name: to.Ptr("Indirect Attack"),
	// 		IsMultiLevelFilter: to.Ptr(false),
	// 		Source: to.Ptr(armcognitiveservices.RaiPolicyContentSourcePrompt),
	// 	},
	// }
}
