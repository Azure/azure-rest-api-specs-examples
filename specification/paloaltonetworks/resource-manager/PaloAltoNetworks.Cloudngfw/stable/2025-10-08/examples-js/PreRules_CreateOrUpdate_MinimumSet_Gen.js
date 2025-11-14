const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a PreRulesResource
 *
 * @summary Create a PreRulesResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2025-10-08/examples/PreRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function preRulesCreateOrUpdateMinimumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const resource = { ruleName: "preRule1" };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.beginCreateOrUpdateAndWait(
    globalRulestackName,
    priority,
    resource,
  );
  console.log(result);
}
