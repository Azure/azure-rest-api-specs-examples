package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/GetEffectiveInformationProtectionPolicy_example.json
func ExampleInformationProtectionPoliciesClient_Get_getTheEffectiveInformationProtectionPolicyForAManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewInformationProtectionPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e", armsecurity.InformationProtectionPolicyNameEffective, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
