const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists pools for a project
 *
 * @summary Lists pools for a project
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Pools_List.json
 */
async function poolsListByProject() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "DevProject";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.pools.listByProject(resourceGroupName, projectName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

poolsListByProject().catch(console.error);
