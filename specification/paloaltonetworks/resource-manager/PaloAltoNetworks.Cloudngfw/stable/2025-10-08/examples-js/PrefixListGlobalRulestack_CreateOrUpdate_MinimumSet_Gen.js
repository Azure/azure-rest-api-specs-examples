const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a PrefixListGlobalRulestackResource
 *
 * @summary Create a PrefixListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2025-10-08/examples/PrefixListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function prefixListGlobalRulestackCreateOrUpdateMinimumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const resource = {
    prefixList: ["1.0.0.0/24"],
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.prefixListGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource,
  );
  console.log(result);
}
