package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic/v3"
)

// Generated from example definition: 2025-06-01/DetachAndDeleteTrafficFilter_Delete.json
func ExampleDetachAndDeleteTrafficFilterClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDetachAndDeleteTrafficFilterClient().Delete(ctx, "myResourceGroup", "myMonitor", &armelastic.DetachAndDeleteTrafficFilterClientDeleteOptions{
		RulesetID: to.Ptr("31d91b5afb6f4c2eaaf104c97b1991dd")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armelastic.DetachAndDeleteTrafficFilterClientDeleteResponse{
	// }
}
