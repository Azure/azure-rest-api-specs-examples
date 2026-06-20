package armnapsteromniagentapi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/napsteromniagentapi/armnapsteromniagentapi"
)

// Generated from example definition: 2025-12-24-preview/Organizations_LatestLinkedSaaS_MaximumSet_Gen.json
func ExampleOrganizationsClient_LatestLinkedSaaS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnapsteromniagentapi.NewClientFactory("0F0FBCF9-8374-47FC-B189-B79B84033EA3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().LatestLinkedSaaS(ctx, "rgopenapi", "contosoOrg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnapsteromniagentapi.OrganizationsClientLatestLinkedSaaSResponse{
	// 	LatestLinkedSaaSResponse: armnapsteromniagentapi.LatestLinkedSaaSResponse{
	// 		SaaSResourceID: to.Ptr("/subscriptions/0F0FBCF9-8374-47FC-B189-B79B84033EA3/resourceGroups/rgopenapi/providers/Microsoft.SaaS/resources/contosoSaaS"),
	// 		IsHiddenSaaS: to.Ptr(true),
	// 	},
	// }
}
