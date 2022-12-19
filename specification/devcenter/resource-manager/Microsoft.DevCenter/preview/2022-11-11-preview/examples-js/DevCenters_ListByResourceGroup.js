const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all devcenters in a resource group.
 *
 * @summary Lists all devcenters in a resource group.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/DevCenters_ListByResourceGroup.json
 */
async function devCentersListByResourceGroup() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devCenters.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

devCentersListByResourceGroup().catch(console.error);
