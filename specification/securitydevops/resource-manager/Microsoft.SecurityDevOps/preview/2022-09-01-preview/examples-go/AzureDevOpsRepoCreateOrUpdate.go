package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoCreateOrUpdate.json
func ExampleAzureDevOpsRepoClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsRepoClient().BeginCreateOrUpdate(ctx, "westusrg", "testconnector", "myOrg", "myProject", "myRepo", armsecuritydevops.AzureDevOpsRepo{
		Properties: &armsecuritydevops.AzureDevOpsRepoProperties{
			ActionableRemediation: &armsecuritydevops.ActionableRemediation{
				BranchConfiguration: &armsecuritydevops.TargetBranchConfiguration{
					Names: []*string{
						to.Ptr("main")},
				},
				Categories: []*armsecuritydevops.RuleCategory{
					to.Ptr(armsecuritydevops.RuleCategorySecrets)},
				SeverityLevels: []*string{
					to.Ptr("High")},
				State: to.Ptr(armsecuritydevops.ActionableRemediationStateEnabled),
			},
			RepoID:  to.Ptr("00000000-0000-0000-0000-000000000000"),
			RepoURL: to.Ptr("https://dev.azure.com/myOrg/myProject/_git/myRepo"),
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
	// res.AzureDevOpsRepo = armsecuritydevops.AzureDevOpsRepo{
	// 	Name: to.Ptr("myRepo"),
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/azureDevOpsConnectos/orgs/projects/repos"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg/projects/myProject/repos/myRepo"),
	// 	Properties: &armsecuritydevops.AzureDevOpsRepoProperties{
	// 		ActionableRemediation: &armsecuritydevops.ActionableRemediation{
	// 			BranchConfiguration: &armsecuritydevops.TargetBranchConfiguration{
	// 				Names: []*string{
	// 					to.Ptr("main")},
	// 				},
	// 				Categories: []*armsecuritydevops.RuleCategory{
	// 					to.Ptr(armsecuritydevops.RuleCategorySecrets)},
	// 					SeverityLevels: []*string{
	// 						to.Ptr("High")},
	// 						State: to.Ptr(armsecuritydevops.ActionableRemediationStateEnabled),
	// 					},
	// 					ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 					RepoID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					RepoURL: to.Ptr("https://dev.azure.com/myOrg/myProject/_git/myRepo"),
	// 				},
	// 			}
}
