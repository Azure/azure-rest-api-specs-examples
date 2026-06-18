const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to revert rulestack configuration
 *
 * @summary revert rulestack configuration
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_revert_MaximumSet_Gen.json
 */
async function globalRulestackRevertMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  await client.globalRulestack.revert("praval");
}
