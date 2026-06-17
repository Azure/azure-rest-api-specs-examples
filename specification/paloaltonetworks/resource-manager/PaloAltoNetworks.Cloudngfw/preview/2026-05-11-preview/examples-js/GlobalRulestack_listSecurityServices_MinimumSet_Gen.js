const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list the security services for rulestack
 *
 * @summary list the security services for rulestack
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_listSecurityServices_MinimumSet_Gen.json
 */
async function globalRulestackListSecurityServicesMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listSecurityServices("praval", "globalRulestacks");
  console.log(result);
}
