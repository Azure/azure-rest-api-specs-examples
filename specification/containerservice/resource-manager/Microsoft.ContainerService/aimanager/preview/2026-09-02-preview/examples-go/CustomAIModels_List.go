package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-09-02-preview/CustomAIModels_List.json
func ExampleCustomAIModelsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomAIModelsClient().NewListPager("rg1", "aimanager1", nil)
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
		// page = armcontainerserviceaimanager.CustomAIModelsClientListResponse{
		// 	CustomAIModelListResult: armcontainerserviceaimanager.CustomAIModelListResult{
		// 		Value: []*armcontainerserviceaimanager.CustomAIModel{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1/customAIModels/custom-model1"),
		// 				Name: to.Ptr("custom-model1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/aiManagers/customAIModels"),
		// 				SystemData: &armcontainerserviceaimanager.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 				ETag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
		// 				Properties: &armcontainerserviceaimanager.CustomAIModelProperties{
		// 					ProvisioningState: to.Ptr(armcontainerserviceaimanager.CustomAIModelProvisioningStateSucceeded),
		// 					ModelID: to.Ptr("meta-llama/Llama-2-7b-chat"),
		// 					BaseModel: &armcontainerserviceaimanager.BaseModelReference{
		// 						ID: to.Ptr("meta-llama/Llama-2-7b-chat"),
		// 					},
		// 					ModelSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/foundry-private-source"),
		// 					Description: to.Ptr("Custom Llama 2 7B model for our organization"),
		// 					Spec: &armcontainerserviceaimanager.CustomAIModelSpec{
		// 						License: to.Ptr("llama2"),
		// 						IsRestricted: to.Ptr(false),
		// 						MaxContextLength: to.Ptr[int32](4096),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
