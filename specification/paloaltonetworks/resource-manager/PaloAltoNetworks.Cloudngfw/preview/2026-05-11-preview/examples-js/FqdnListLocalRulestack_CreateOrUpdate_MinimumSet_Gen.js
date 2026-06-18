const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FqdnListLocalRulestackResource
 *
 * @summary create a FqdnListLocalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/FqdnListLocalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function fqdnListLocalRulestackCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.fqdnListLocalRulestack.createOrUpdate("rgopenapi", "lrs1", "armid1", {
    fqdnList: ["string1", "string2"],
  });
  console.log(result);
}
