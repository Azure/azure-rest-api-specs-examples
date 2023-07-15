const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a FqdnListGlobalRulestackResource
 *
 * @summary Delete a FqdnListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/FqdnListGlobalRulestack_Delete_MinimumSet_Gen.json
 */
async function fqdnListGlobalRulestackDeleteMinimumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.fqdnListGlobalRulestack.beginDeleteAndWait(globalRulestackName, name);
  console.log(result);
}
