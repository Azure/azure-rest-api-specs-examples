package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/FeaturestoreEntityContainer/list.json
func ExampleFeaturestoreEntityContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFeaturestoreEntityContainersClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearning.FeaturestoreEntityContainersClientListOptions{Skip: nil,
		Tags:         to.Ptr("string"),
		ListViewType: to.Ptr(armmachinelearning.ListViewTypeAll),
		PageSize:     nil,
		Name:         nil,
		Description:  nil,
		CreatedBy:    nil,
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
		// page.FeaturestoreEntityContainerResourceArmPaginatedResult = armmachinelearning.FeaturestoreEntityContainerResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.FeaturestoreEntityContainer{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:51:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:51:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armmachinelearning.FeaturestoreEntityContainerProperties{
		// 				Description: to.Ptr("string"),
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				IsArchived: to.Ptr(false),
		// 				LatestVersion: to.Ptr("string"),
		// 				NextVersion: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armmachinelearning.AssetProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
