package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3db6867b8e524ea6d1bc7a3bbb989fe50dd2f184/specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/SearchCheckNameAvailability.json
func ExampleServicesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckNameAvailability(ctx, armsearch.CheckNameAvailabilityInput{
		Name: to.Ptr("mysearchservice"),
		Type: to.Ptr("searchServices"),
	}, &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityOutput = armsearch.CheckNameAvailabilityOutput{
	// 	Message: to.Ptr(""),
	// 	IsNameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armsearch.UnavailableNameReasonAlreadyExists),
	// }
}
