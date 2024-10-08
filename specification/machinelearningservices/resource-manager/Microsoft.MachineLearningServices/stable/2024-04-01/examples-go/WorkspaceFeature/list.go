package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/WorkspaceFeature/list.json
func ExampleWorkspaceFeaturesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceFeaturesClient().NewListPager("myResourceGroup", "testworkspace", nil)
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
		// page.ListAmlUserFeatureResult = armmachinelearning.ListAmlUserFeatureResult{
		// 	Value: []*armmachinelearning.AmlUserFeature{
		// 		{
		// 			Description: to.Ptr("Create, edit or delete AutoML experiments in the SDK"),
		// 			DisplayName: to.Ptr("Create edit experiments UI"),
		// 			ID: to.Ptr("automatedml_createeditexperimentsui"),
		// 		},
		// 		{
		// 			Description: to.Ptr("Upgrade workspace from Basic to enterprise from the UI"),
		// 			DisplayName: to.Ptr("Upgrade workspace UI"),
		// 			ID: to.Ptr("workspace_upgradeworkspaceui"),
		// 	}},
		// }
	}
}
