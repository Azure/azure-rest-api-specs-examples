package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/Workflow_CreateOrUpdate.json
func ExampleWorkflowClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevhub.NewWorkflowClient("subscriptionId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "resourceGroup1", "workflow1", armdevhub.Workflow{
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"appname": to.Ptr("testApp"),
		},
		Properties: &armdevhub.WorkflowProperties{
			GithubWorkflowProfile: &armdevhub.GitHubWorkflowProfile{
				Acr: &armdevhub.ACR{
					AcrRegistryName:   to.Ptr("registry1"),
					AcrRepositoryName: to.Ptr("repo1"),
					AcrResourceGroup:  to.Ptr("resourceGroup1"),
					AcrSubscriptionID: to.Ptr("subscriptionId1"),
				},
				AksResourceID: to.Ptr("/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
				BranchName:    to.Ptr("branch1"),
				DeploymentProperties: &armdevhub.DeploymentProperties{
					KubeManifestLocations: []*string{
						to.Ptr("/src/manifests/")},
					ManifestType: to.Ptr(armdevhub.ManifestTypeKube),
					Overrides: map[string]*string{
						"key1": to.Ptr("value1"),
					},
				},
				DockerBuildContext: to.Ptr("repo1/src/"),
				Dockerfile:         to.Ptr("repo1/images/Dockerfile"),
				OidcCredentials: &armdevhub.GitHubWorkflowProfileOidcCredentials{
					AzureClientID: to.Ptr("12345678-3456-7890-5678-012345678901"),
					AzureTenantID: to.Ptr("66666666-3456-7890-5678-012345678901"),
				},
				RepositoryName:  to.Ptr("repo1"),
				RepositoryOwner: to.Ptr("owner1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
