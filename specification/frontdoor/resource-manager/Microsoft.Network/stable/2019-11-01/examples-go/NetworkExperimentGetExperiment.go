package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetExperiment.json
func ExampleExperimentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().Get(ctx, "MyResourceGroup", "MyProfile", "MyExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Experiment = armfrontdoor.Experiment{
	// 	Name: to.Ptr("MyExperiment"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/MyResourceGroup/providers/Microsoft.Network/NetworkExperimentProfiles/MyProfile/Experiments/MyExperiment"),
	// 	Location: to.Ptr("WestUs"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armfrontdoor.ExperimentProperties{
	// 		Description: to.Ptr("this is my first experiment!"),
	// 		EnabledState: to.Ptr(armfrontdoor.StateEnabled),
	// 		EndpointA: &armfrontdoor.Endpoint{
	// 			Name: to.Ptr("endpoint A"),
	// 			Endpoint: to.Ptr("endpointA.net"),
	// 		},
	// 		EndpointB: &armfrontdoor.Endpoint{
	// 			Name: to.Ptr("endpoint B"),
	// 			Endpoint: to.Ptr("endpointB.net"),
	// 		},
	// 		ResourceState: to.Ptr(armfrontdoor.NetworkExperimentResourceStateCreating),
	// 		ScriptFileURI: to.Ptr("www.myScript.com"),
	// 		Status: to.Ptr("ongoing"),
	// 	},
	// }
}
