const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Deployment stack by name at Management Group scope. When operation completes, status code 200 returned without content.
 *
 * @summary Deletes a Deployment stack by name at Management Group scope. When operation completes, status code 200 returned without content.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackManagementGroupDelete.json
 */
async function deploymentStacksManagementGroupDelete() {
  const managementGroupId = "myMg";
  const deploymentStackName = "simpleDeploymentStack";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const result = await client.deploymentStacks.beginDeleteAtManagementGroupAndWait(
    managementGroupId,
    deploymentStackName,
  );
  console.log(result);
}
