const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a PrefixListGlobalRulestackResource
 *
 * @summary Delete a PrefixListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PrefixListGlobalRulestack_Delete_MaximumSet_Gen.json
 */
async function prefixListGlobalRulestackDeleteMaximumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.prefixListGlobalRulestack.beginDeleteAndWait(
    globalRulestackName,
    name
  );
  console.log(result);
}
