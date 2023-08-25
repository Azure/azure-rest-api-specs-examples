const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Deployment Stack by name. When operation completes, status code 200 returned without content.
 *
 * @summary Deletes a Deployment Stack by name. When operation completes, status code 200 returned without content.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/preview/2022-08-01-preview/examples/DeploymentStackManagementGroupDelete.json
 */
async function deploymentStacksDelete() {
  const managementGroupId = "myMg";
  const deploymentStackName = "simpleDeploymentStack";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const result = await client.deploymentStacks.beginDeleteAtManagementGroupAndWait(
    managementGroupId,
    deploymentStackName
  );
  console.log(result);
}
