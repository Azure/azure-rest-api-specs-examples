const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get counters
 *
 * @summary Get counters
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PreRules_getCounters_MinimumSet_Gen.json
 */
async function preRulesGetCountersMinimumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.getCounters(globalRulestackName, priority);
  console.log(result);
}
