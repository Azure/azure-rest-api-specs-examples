const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a GlobalRulestackResource
 *
 * @summary Update a GlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/GlobalRulestack_Update_MinimumSet_Gen.json
 */
async function globalRulestackUpdateMinimumSetGen() {
  const globalRulestackName = "praval";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.update(globalRulestackName, properties);
  console.log(result);
}
