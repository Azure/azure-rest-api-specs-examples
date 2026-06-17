const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get changelog
 *
 * @summary get changelog
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_getChangeLog_MaximumSet_Gen.json
 */
async function globalRulestackGetChangeLogMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.getChangeLog("praval");
  console.log(result);
}
