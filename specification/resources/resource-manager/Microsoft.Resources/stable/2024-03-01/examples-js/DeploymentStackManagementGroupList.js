const { DeploymentStacksClient } = require("@azure/arm-resourcesdeploymentstacks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Deployment stacks within the specified Management Group.
 *
 * @summary Lists all the Deployment stacks within the specified Management Group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackManagementGroupList.json
 */
async function deploymentStacksManagementGroupList() {
  const managementGroupId = "myMg";
  const credential = new DefaultAzureCredential();
  const client = new DeploymentStacksClient(credential);
  const resArray = new Array();
  for await (let item of client.deploymentStacks.listAtManagementGroup(managementGroupId)) {
    resArray.push(item);
  }
  console.log(resArray);
}
