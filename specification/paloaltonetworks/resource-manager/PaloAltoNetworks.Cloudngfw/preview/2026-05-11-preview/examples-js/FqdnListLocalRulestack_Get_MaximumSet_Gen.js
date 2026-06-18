const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a FqdnListLocalRulestackResource
 *
 * @summary get a FqdnListLocalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/FqdnListLocalRulestack_Get_MaximumSet_Gen.json
 */
async function fqdnListLocalRulestackGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.fqdnListLocalRulestack.get("rgopenapi", "lrs1", "armid1");
  console.log(result);
}
