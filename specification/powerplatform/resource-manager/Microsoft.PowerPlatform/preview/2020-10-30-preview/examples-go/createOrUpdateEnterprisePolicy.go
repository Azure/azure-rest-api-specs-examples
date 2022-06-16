package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/createOrUpdateEnterprisePolicy.json
func ExampleEnterprisePoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerplatform.NewEnterprisePoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"enterprisePolicy",
		"resourceGroup",
		armpowerplatform.EnterprisePolicy{
			Location: to.Ptr("East US"),
			Tags: map[string]*string{
				"Organization": to.Ptr("Administration"),
			},
			Identity: &armpowerplatform.EnterprisePolicyIdentity{
				Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
			},
			Kind: to.Ptr(armpowerplatform.EnterprisePolicyKindLockbox),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
