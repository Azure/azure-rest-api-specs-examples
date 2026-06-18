const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a FirewallResource
 *
 * @summary delete a FirewallResource
 * x-ms-original-file: 2026-05-11-preview/Firewalls_Delete_MaximumSet_Gen.json
 */
async function firewallsDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  await client.firewalls.delete("firewall-rg", "firewall1");
}
