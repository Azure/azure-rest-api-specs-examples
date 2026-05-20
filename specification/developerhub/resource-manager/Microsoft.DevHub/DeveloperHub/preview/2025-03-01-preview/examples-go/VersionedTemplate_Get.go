package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/VersionedTemplate_Get.json
func ExampleVersionedTemplateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("a0a37f63-7183-4e86-9ac7-ce8036a3ed31", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVersionedTemplateClient().Get(ctx, "test-template", "0.0.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevhub.VersionedTemplateClientGetResponse{
	// 	VersionedTemplate: armdevhub.VersionedTemplate{
	// 		Name: to.Ptr("0.0.1"),
	// 		Type: to.Ptr("Micfosoft.DevHub/templates/versions"),
	// 		ID: to.Ptr("/subscriptions/a0a37f63-7183-4e86-9ac7-ce8036a3ed31/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
	// 		Properties: &armdevhub.VersionedTemplateProperties{
	// 			Parameters: []*armdevhub.Parameter{
	// 				{
	// 					Name: to.Ptr("test-parameter"),
	// 					Default: &armdevhub.ParameterDefault{
	// 						Value: to.Ptr("test"),
	// 					},
	// 					ParameterKind: to.Ptr(armdevhub.ParameterKindAzureResourceGroup),
	// 					ParameterType: to.Ptr(armdevhub.ParameterTypeString),
	// 					Required: to.Ptr(true),
	// 				},
	// 			},
	// 			TemplateType: to.Ptr(armdevhub.TemplateTypeManifest),
	// 			Version: to.Ptr("0.0.1"),
	// 		},
	// 		SystemData: &armdevhub.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T09:56:23.933651400Z"); return t}()),
	// 			CreatedBy: to.Ptr(""),
	// 			CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T14:35:15.206059900Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("aks-dev-hub-devhub"),
	// 			LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
