const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes the specified Bastion Host.
 *
 * @summary deletes the specified Bastion Host.
 * x-ms-original-file: 2025-07-01/BastionHostDeveloperDelete.json
 */
async function deleteDeveloperBastionHost() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  await client.bastionHosts.delete("rg2", "bastionhostdeveloper");
}
