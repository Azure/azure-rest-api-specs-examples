package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/UpdateAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsOrgsClient().BeginUpdate(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", armsecurity.AzureDevOpsOrg{
		Properties: &armsecurity.AzureDevOpsOrgProperties{
			ActionableRemediation: &armsecurity.ActionableRemediation{
				State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
			},
			OnboardingState: to.Ptr(armsecurity.OnboardingStateNotApplicable),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsOrg = armsecurity.AzureDevOpsOrg{
	// 	Name: to.Ptr("myAzDevOpsOrg"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
	// 	Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 	},
	// }
}
