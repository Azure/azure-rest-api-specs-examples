const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list of AppIds for GlobalRulestack ApiVersion
 *
 * @summary list of AppIds for GlobalRulestack ApiVersion
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_listAppIds_MinimumSet_Gen.json
 */
async function globalRulestackListAppIdsMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listAppIds("praval");
  console.log(result);
}
