const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Deployment stack with a given name at Management Group scope.
 *
 * @summary Gets a Deployment stack with a given name at Management Group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackManagementGroupGet.json
 */
async function deploymentStacksManagementGroupGet() {
  const managementGroupId = "myMg";
  const deploymentStackName = "simpleDeploymentStack";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const result = await client.deploymentStacks.getAtManagementGroup(
    managementGroupId,
    deploymentStackName,
  );
  console.log(result);
}
