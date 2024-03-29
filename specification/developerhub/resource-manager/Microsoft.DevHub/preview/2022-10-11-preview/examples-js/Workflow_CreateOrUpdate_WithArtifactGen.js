const { DeveloperHubServiceClient } = require("@azure/arm-devhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workflow
 *
 * @summary Creates or updates a workflow
 * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_CreateOrUpdate_WithArtifactGen.json
 */
async function createWorkflowWithArtifactGeneration() {
  const subscriptionId = process.env["DEVHUB_SUBSCRIPTION_ID"] || "subscriptionId1";
  const resourceGroupName = process.env["DEVHUB_RESOURCE_GROUP"] || "resourceGroup1";
  const workflowName = "workflow1";
  const parameters = {
    acr: {
      acrRegistryName: "registry1",
      acrRepositoryName: "repo1",
      acrResourceGroup: "resourceGroup1",
      acrSubscriptionId: "subscriptionId1",
    },
    aksResourceId:
      "/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1",
    appName: "my-app",
    branchName: "branch1",
    deploymentProperties: {
      kubeManifestLocations: ["/src/manifests/"],
      manifestType: "kube",
      overrides: { key1: "value1" },
    },
    dockerBuildContext: "repo1/src/",
    dockerfile: "repo1/images/Dockerfile",
    dockerfileGenerationMode: "enabled",
    dockerfileOutputDirectory: "./",
    generationLanguage: "javascript",
    imageName: "myimage",
    imageTag: "latest",
    languageVersion: "14",
    location: "location1",
    manifestGenerationMode: "enabled",
    manifestOutputDirectory: "./",
    manifestType: "kube",
    namespacePropertiesArtifactGenerationPropertiesNamespace: "my-namespace",
    oidcCredentials: {
      azureClientId: "12345678-3456-7890-5678-012345678901",
      azureTenantId: "66666666-3456-7890-5678-012345678901",
    },
    port: "80",
    repositoryName: "repo1",
    repositoryOwner: "owner1",
    tags: { appname: "testApp" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeveloperHubServiceClient(credential, subscriptionId);
  const result = await client.workflowOperations.createOrUpdate(
    resourceGroupName,
    workflowName,
    parameters
  );
  console.log(result);
}
