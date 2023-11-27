const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reset counters
 *
 * @summary Reset counters
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_resetCounters_MaximumSet_Gen.json
 */
async function preRulesResetCountersMaximumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const firewallName = "firewall1";
  const options = { firewallName };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.resetCounters(globalRulestackName, priority, options);
  console.log(result);
}
