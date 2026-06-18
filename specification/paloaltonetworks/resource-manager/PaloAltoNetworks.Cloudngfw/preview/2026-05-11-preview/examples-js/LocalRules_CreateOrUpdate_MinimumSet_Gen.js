const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a LocalRulesResource
 *
 * @summary create a LocalRulesResource
 * x-ms-original-file: 2026-05-11-preview/LocalRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function localRulesCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.localRules.createOrUpdate("firewall-rg", "lrs1", "1", {
    ruleName: "localRule1",
  });
  console.log(result);
}
