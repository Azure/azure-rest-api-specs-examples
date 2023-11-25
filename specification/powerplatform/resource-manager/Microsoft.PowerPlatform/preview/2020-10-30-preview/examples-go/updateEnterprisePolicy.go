package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/updateEnterprisePolicy.json
func ExampleEnterprisePoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnterprisePoliciesClient().Update(ctx, "enterprisePolicy", "resourceGroup", armpowerplatform.PatchEnterprisePolicy{
		Location: to.Ptr("East US"),
		Tags: map[string]*string{
			"Organization": to.Ptr("Administration"),
		},
		Identity: &armpowerplatform.EnterprisePolicyIdentity{
			Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnterprisePolicy = armpowerplatform.EnterprisePolicy{
	// 	Name: to.Ptr("enterprisePolicy"),
	// 	Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.PowerPlatform/enterprisePolicies/enterprisePolicy"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Organization": to.Ptr("Administration"),
	// 	},
	// 	Identity: &armpowerplatform.EnterprisePolicyIdentity{
	// 		Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
	// 		SystemAssignedIdentityPrincipalID: to.Ptr("principalId"),
	// 		TenantID: to.Ptr("tenantId"),
	// 	},
	// 	Kind: to.Ptr(armpowerplatform.EnterprisePolicyKindLockbox),
	// 	SystemData: &armpowerplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 	},
	// }
}
