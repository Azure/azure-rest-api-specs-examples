const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified Bastion Host.
 *
 * @summary gets the specified Bastion Host.
 * x-ms-original-file: 2025-07-01/BastionHostGet.json
 */
async function getBastionHost() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.get("rg1", "bastionhosttenant'");
  console.log(result);
}
