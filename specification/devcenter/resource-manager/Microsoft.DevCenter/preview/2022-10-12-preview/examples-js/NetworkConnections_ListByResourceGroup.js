const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists network connections in a resource group
 *
 * @summary Lists network connections in a resource group
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/NetworkConnections_ListByResourceGroup.json
 */
async function networkConnectionsListByResourceGroup() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkConnections.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

networkConnectionsListByResourceGroup().catch(console.error);
