package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/PrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "rg1", "ddb1", "sql", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armpowerplatform.PrivateLinkResource{
	// 	Name: to.Ptr("sql"),
	// 	Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/privateLinkResources"),
	// 	ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.PowerPlatform/enterprisePolicies/ddb1/privateLinkResources/sql"),
	// 	Properties: &armpowerplatform.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("sql"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("ddb1"),
	// 			to.Ptr("ddb1-westus")},
	// 		},
	// 	}
}
