const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list predefined URL categories for rulestack
 *
 * @summary list predefined URL categories for rulestack
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_listPredefinedUrlCategories_MinimumSet_Gen.json
 */
async function globalRulestackListPredefinedUrlCategoriesMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listPredefinedUrlCategories("praval");
  console.log(result);
}
