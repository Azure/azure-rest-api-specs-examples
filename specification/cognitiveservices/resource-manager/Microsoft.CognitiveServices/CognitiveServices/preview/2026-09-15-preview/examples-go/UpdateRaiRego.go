package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/UpdateRaiRego.json
func ExampleRaiRegosClient_CreateOrUpdate_replaceAReusableRegoArtifactConditionally() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiRegosClient().CreateOrUpdate(ctx, "resource-group", "safety-account", "input-guard", armcognitiveservices.RaiRego{
		Properties: &armcognitiveservices.RaiRegoProperties{
			Encoding: to.Ptr(armcognitiveservices.RaiRegoEncodingNone),
			Rego:     to.Ptr("package input_guard\n\nimport rego.v1\n\ndefault verdict := {\"decision\": \"allow\"}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.Hate.severity >= 4\n}\n\nverdict := {\"decision\": \"deny\", \"reason\": \"unsafe_input\"} if {\n  blocking_signal\n}"),
		},
	}, &armcognitiveservices.RaiRegosClientCreateOrUpdateOptions{
		IfMatch: to.Ptr("\"00000000-0000-0000-0000-000000000001\"")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiRegosClientCreateOrUpdateResponse{
	// 	EtagHeader: to.Ptr("\"00000000-0000-0000-0000-000000000004\""),
	// 	RaiRego: armcognitiveservices.RaiRego{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiRegos/input-guard"),
	// 		Name: to.Ptr("input-guard"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiRegos"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000004\""),
	// 		Properties: &armcognitiveservices.RaiRegoProperties{
	// 			Encoding: to.Ptr(armcognitiveservices.RaiRegoEncodingNone),
	// 			Rego: to.Ptr("package input_guard\n\nimport rego.v1\n\ndefault verdict := {\"decision\": \"allow\"}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.Hate.severity >= 4\n}\n\nverdict := {\"decision\": \"deny\", \"reason\": \"unsafe_input\"} if {\n  blocking_signal\n}"),
	// 		},
	// 	},
	// }
}
