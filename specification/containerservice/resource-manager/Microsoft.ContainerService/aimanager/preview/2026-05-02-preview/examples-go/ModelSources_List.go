package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/ModelSources_List.json
func ExampleModelSourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModelSourcesClient().NewListPager("rgaimanagers", "aimanager1", nil)
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
		// page = armcontainerserviceaimanager.ModelSourcesClientListResponse{
		// 	ModelSourceListResult: armcontainerserviceaimanager.ModelSourceListResult{
		// 		Value: []*armcontainerserviceaimanager.ModelSource{
		// 			{
		// 				Properties: &armcontainerserviceaimanager.ModelSourceProperties{
		// 					ProvisioningState: to.Ptr(armcontainerserviceaimanager.ResourceProvisioningStateSucceeded),
		// 					SourceType: to.Ptr(armcontainerserviceaimanager.ModelSourceTypeHuggingFace),
		// 					Description: to.Ptr("Hugging Face model source"),
		// 				},
		// 				ETag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/huggingface"),
		// 				Name: to.Ptr("huggingface"),
		// 				Type: to.Ptr("Microsoft.ContainerService/aiManagers/modelSources"),
		// 				SystemData: &armcontainerserviceaimanager.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources?api-version=2026-05-02-preview&$skiptoken=token"),
		// 	},
		// }
	}
}
