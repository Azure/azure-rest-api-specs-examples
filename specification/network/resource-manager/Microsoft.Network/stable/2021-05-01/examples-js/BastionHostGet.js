const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Bastion Host.
 *
 * @summary Gets the specified Bastion Host.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/BastionHostGet.json
 */
async function getBastionHost() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const bastionHostName = "bastionhosttenant'";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.get(resourceGroupName, bastionHostName);
  console.log(result);
}

getBastionHost().catch(console.error);
