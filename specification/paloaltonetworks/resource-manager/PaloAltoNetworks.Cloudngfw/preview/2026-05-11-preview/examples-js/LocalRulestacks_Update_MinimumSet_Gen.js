const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a LocalRulestackResource
 *
 * @summary update a LocalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/LocalRulestacks_Update_MinimumSet_Gen.json
 */
async function localRulestacksUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.localRulestacks.update("rgopenapi", "lrs1", {});
  console.log(result);
}
