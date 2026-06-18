const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get counters
 *
 * @summary get counters
 * x-ms-original-file: 2026-05-11-preview/PostRules_getCounters_MinimumSet_Gen.json
 */
async function postRulesGetCountersMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.postRules.getCounters("lrs1", "1");
  console.log(result);
}
