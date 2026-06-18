const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a PrefixListGlobalRulestackResource
 *
 * @summary create a PrefixListGlobalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/PrefixListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function prefixListGlobalRulestackCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.prefixListGlobalRulestack.createOrUpdate("praval", "armid1", {
    prefixList: ["1.0.0.0/24"],
  });
  console.log(result);
}
