const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available service aliases for this subscription in this region.
 *
 * @summary Gets all available service aliases for this subscription in this region.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AvailableServiceAliasesList.json
 */
async function getAvailableServiceAliases() {
  const subscriptionId = "subId";
  const location = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableServiceAliases.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAvailableServiceAliases().catch(console.error);
