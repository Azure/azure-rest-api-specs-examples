const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reset counters
 *
 * @summary Reset counters
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PreRules_resetCounters_MinimumSet_Gen.json
 */
async function preRulesResetCountersMinimumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "lrs1";
  const priority = "1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.preRules.resetCounters(globalRulestackName, priority);
  console.log(result);
}
