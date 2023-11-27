const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a PrefixListGlobalRulestackResource
 *
 * @summary Get a PrefixListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_Get_MinimumSet_Gen.json
 */
async function prefixListGlobalRulestackGetMinimumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.prefixListGlobalRulestack.get(globalRulestackName, name);
  console.log(result);
}
