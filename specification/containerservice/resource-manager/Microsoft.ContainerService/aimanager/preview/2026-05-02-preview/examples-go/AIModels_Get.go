package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/AIModels_Get.json
func ExampleAIModelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAIModelsClient().Get(ctx, "eastus", "9806f0c862fdd920", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerserviceaimanager.AIModelsClientGetResponse{
	// 	AIModel: armcontainerserviceaimanager.AIModel{
	// 		Properties: &armcontainerserviceaimanager.AIModelProperties{
	// 			ModelID: to.Ptr("microsoft/Phi-4-mini-instruct"),
	// 			Description: to.Ptr("Phi-4-mini-instruct is a lightweight open model from Microsoft."),
	// 			Spec: &armcontainerserviceaimanager.ModelSpec{
	// 				License: to.Ptr("MIT"),
	// 				IsRestricted: to.Ptr(false),
	// 				MaxContextLength: to.Ptr[int32](131072),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ContainerService/locations/eastus/aiModels/9806f0c862fdd920"),
	// 		Name: to.Ptr("9806f0c862fdd920"),
	// 		Type: to.Ptr("Microsoft.ContainerService/aiModels"),
	// 		SystemData: &armcontainerserviceaimanager.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
