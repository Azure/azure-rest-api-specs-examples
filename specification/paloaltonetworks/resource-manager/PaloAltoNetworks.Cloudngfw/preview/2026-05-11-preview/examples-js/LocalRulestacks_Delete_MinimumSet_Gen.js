const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a LocalRulestackResource
 *
 * @summary delete a LocalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/LocalRulestacks_Delete_MinimumSet_Gen.json
 */
async function localRulestacksDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  await client.localRulestacks.delete("rgopenapi", "lrs1");
}
