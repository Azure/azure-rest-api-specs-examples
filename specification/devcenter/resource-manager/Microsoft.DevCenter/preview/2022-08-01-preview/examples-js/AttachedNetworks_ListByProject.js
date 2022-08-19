const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the attached NetworkConnections for a Project.
 *
 * @summary Lists the attached NetworkConnections for a Project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_ListByProject.json
 */
async function attachedNetworksListByProject() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.attachedNetworks.listByProject(resourceGroupName, projectName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

attachedNetworksListByProject().catch(console.error);
