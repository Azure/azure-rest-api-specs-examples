package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentListExperiments.json
func ExampleExperimentsClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListByProfilePager("MyResourceGroup", "MyProfile", nil)
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
		// page.ExperimentList = armfrontdoor.ExperimentList{
		// 	Value: []*armfrontdoor.Experiment{
		// 		{
		// 			Name: to.Ptr("MyExperiment"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/MyResourceGroup/providers/Microsoft.Network/NetworkExperimentProfiles/MyProfile/Experiments"),
		// 			Properties: &armfrontdoor.ExperimentProperties{
		// 				Description: to.Ptr("this is my first experiment!"),
		// 				EnabledState: to.Ptr(armfrontdoor.StateEnabled),
		// 				EndpointA: &armfrontdoor.Endpoint{
		// 					Name: to.Ptr("endpoint A"),
		// 					Endpoint: to.Ptr("endpointA.net"),
		// 				},
		// 				EndpointB: &armfrontdoor.Endpoint{
		// 					Name: to.Ptr("endpoint B"),
		// 					Endpoint: to.Ptr("endpointB.net"),
		// 				},
		// 				ResourceState: to.Ptr(armfrontdoor.NetworkExperimentResourceStateCreating),
		// 				ScriptFileURI: to.Ptr("www.myScript.com"),
		// 				Status: to.Ptr("ongoing"),
		// 			},
		// 	}},
		// }
	}
}
