const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the list of advanced security objects
 *
 * @summary get the list of advanced security objects
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_listAdvancedSecurityObjects_MinimumSet_Gen.json
 */
async function globalRulestackListAdvancedSecurityObjectsMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listAdvancedSecurityObjects(
    "praval",
    "globalRulestacks",
  );
  console.log(result);
}
