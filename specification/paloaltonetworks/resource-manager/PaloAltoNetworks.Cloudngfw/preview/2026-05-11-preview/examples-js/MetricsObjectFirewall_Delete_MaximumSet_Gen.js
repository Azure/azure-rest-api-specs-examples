const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a MetricsObjectFirewallResource
 *
 * @summary delete a MetricsObjectFirewallResource
 * x-ms-original-file: 2026-05-11-preview/MetricsObjectFirewall_Delete_MaximumSet_Gen.json
 */
async function metricsObjectFirewallDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaa";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  await client.metricsObjectFirewall.delete("rgopenapi", "aaaaaaaaaaaaaaaaaaaaaaaa");
}
