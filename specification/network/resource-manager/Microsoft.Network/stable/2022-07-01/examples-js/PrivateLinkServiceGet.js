const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified private link service by resource group.
 *
 * @summary Gets the specified private link service by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PrivateLinkServiceGet.json
 */
async function getPrivateLinkService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.get(resourceGroupName, serviceName);
  console.log(result);
}

getPrivateLinkService().catch(console.error);
