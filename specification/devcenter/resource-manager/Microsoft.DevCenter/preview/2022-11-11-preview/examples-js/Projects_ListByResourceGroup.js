const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all projects in the resource group.
 *
 * @summary Lists all projects in the resource group.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Projects_ListByResourceGroup.json
 */
async function projectsListByResourceGroup() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projects.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

projectsListByResourceGroup().catch(console.error);
