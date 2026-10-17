package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/ListRaiRegos.json
func ExampleRaiRegosClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiRegosClient().NewListPager("resource-group", "safety-account", &armcognitiveservices.RaiRegosClientListOptions{
		Top: to.Ptr[int32](10)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcognitiveservices.RaiRegosClientListResponse{
		// 	RaiRegoListResult: armcognitiveservices.RaiRegoListResult{
		// 		Value: []*armcognitiveservices.RaiRego{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiRegos/input-guard"),
		// 				Name: to.Ptr("input-guard"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiRegos"),
		// 				Etag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
		// 				Properties: &armcognitiveservices.RaiRegoProperties{
		// 					Encoding: to.Ptr(armcognitiveservices.RaiRegoEncodingNone),
		// 					Rego: to.Ptr("package input_guard\n\nimport rego.v1\n\ndefault verdict := {\"decision\": \"allow\"}\n\nverdict := {\"decision\": \"deny\", \"reason\": \"prompt_injection_detected\"} if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiRegos?api-version=2026-09-15-preview&%24top=10&%24skiptoken=aW5wdXQtZ3VhcmQ="),
		// 	},
		// }
	}
}
