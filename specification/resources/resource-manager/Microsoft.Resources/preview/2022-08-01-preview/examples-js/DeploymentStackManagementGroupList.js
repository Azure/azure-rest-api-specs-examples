const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Deployment Stacks within the specified management group.
 *
 * @summary Lists all the Deployment Stacks within the specified management group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/preview/2022-08-01-preview/examples/DeploymentStackManagementGroupList.json
 */
async function deploymentStacksList() {
  const managementGroupId = "myMg";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const resArray = new Array();
  for await (let item of client.deploymentStacks.listAtManagementGroup(managementGroupId)) {
    resArray.push(item);
  }
  console.log(resArray);
}
