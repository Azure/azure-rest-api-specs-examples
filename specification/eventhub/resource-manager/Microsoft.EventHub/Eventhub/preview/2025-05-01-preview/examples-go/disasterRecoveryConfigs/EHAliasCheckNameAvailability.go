package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2025-05-01-preview/disasterRecoveryConfigs/EHAliasCheckNameAvailability.json
func ExampleDisasterRecoveryConfigsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().CheckNameAvailability(ctx, "exampleResourceGroup", "sdk-Namespace-9080", armeventhub.CheckNameAvailabilityParameter{
		Name: to.Ptr("sdk-DisasterRecovery-9474"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResult: armeventhub.CheckNameAvailabilityResult{
	// 		Message: to.Ptr(""),
	// 		NameAvailable: to.Ptr(true),
	// 		Reason: to.Ptr(armeventhub.UnavailableReasonNone),
	// 	},
	// }
}
