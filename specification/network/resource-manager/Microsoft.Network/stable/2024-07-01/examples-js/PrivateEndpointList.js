const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets all private endpoints in a resource group.
 *
 * @summary Gets all private endpoints in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PrivateEndpointList.json
 */
async function listPrivateEndpointsInResourceGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.privateEndpoints.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
