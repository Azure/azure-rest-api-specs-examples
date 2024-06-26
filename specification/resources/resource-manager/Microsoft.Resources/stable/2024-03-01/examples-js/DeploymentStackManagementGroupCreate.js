const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Deployment stack at Management Group scope.
 *
 * @summary Creates or updates a Deployment stack at Management Group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackManagementGroupCreate.json
 */
async function deploymentStacksManagementGroupCreateOrUpdate() {
  const managementGroupId = "myMg";
  const deploymentStackName = "simpleDeploymentStack";
  const deploymentStack = {
    location: "eastus",
    properties: {
      actionOnUnmanage: {
        managementGroups: "detach",
        resourceGroups: "delete",
        resources: "delete",
      },
      denySettings: {
        applyToChildScopes: false,
        excludedActions: ["action"],
        excludedPrincipals: ["principal"],
        mode: "denyDelete",
      },
      parameters: { parameter1: { value: "a string" } },
    },
    tags: { tagkey: "tagVal" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const result = await client.deploymentStacks.beginCreateOrUpdateAtManagementGroupAndWait(
    managementGroupId,
    deploymentStackName,
    deploymentStack,
  );
  console.log(result);
}
