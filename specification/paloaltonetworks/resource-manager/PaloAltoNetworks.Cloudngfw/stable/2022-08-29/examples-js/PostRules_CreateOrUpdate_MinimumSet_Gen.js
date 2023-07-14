const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a PostRulesResource
 *
 * @summary Create a PostRulesResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PostRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function postRulesCreateOrUpdateMinimumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const resource = { ruleName: "postRule1" };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.postRules.beginCreateOrUpdateAndWait(
    globalRulestackName,
    priority,
    resource
  );
  console.log(result);
}
