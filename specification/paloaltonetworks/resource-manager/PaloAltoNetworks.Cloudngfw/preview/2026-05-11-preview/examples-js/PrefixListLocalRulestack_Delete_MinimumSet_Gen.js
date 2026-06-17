const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a PrefixListResource
 *
 * @summary delete a PrefixListResource
 * x-ms-original-file: 2026-05-11-preview/PrefixListLocalRulestack_Delete_MinimumSet_Gen.json
 */
async function prefixListLocalRulestackDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  await client.prefixListLocalRulestack.delete("rgopenapi", "lrs1", "armid1");
}
