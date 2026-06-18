const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a GlobalRulestackResource
 *
 * @summary update a GlobalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_Update_MinimumSet_Gen.json
 */
async function globalRulestackUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.update("praval", {});
  console.log(result);
}
