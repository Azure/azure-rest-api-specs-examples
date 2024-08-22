package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/EnvironmentContainer/list.json
func ExampleRegistryEnvironmentContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistryEnvironmentContainersClient().NewListPager("testrg123", "testregistry", &armmachinelearning.RegistryEnvironmentContainersClientListOptions{Skip: nil,
		ListViewType: nil,
	})
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
		// page.EnvironmentContainerResourceArmPaginatedResult = armmachinelearning.EnvironmentContainerResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.EnvironmentContainer{
		// 		{
		// 			Name: to.Ptr("testEnvironment"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/registries/environments"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/registries/testregistry/environments/testEnvironment"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("John Smith"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("John Smith"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmachinelearning.EnvironmentContainerProperties{
		// 				Description: to.Ptr("string"),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
