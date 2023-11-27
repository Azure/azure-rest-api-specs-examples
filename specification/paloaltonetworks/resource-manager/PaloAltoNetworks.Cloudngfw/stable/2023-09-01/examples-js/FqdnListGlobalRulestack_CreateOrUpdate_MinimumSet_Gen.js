const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FqdnListGlobalRulestackResource
 *
 * @summary Create a FqdnListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/FqdnListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function fqdnListGlobalRulestackCreateOrUpdateMinimumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const resource = {
    fqdnList: ["string1", "string2"],
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.fqdnListGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource
  );
  console.log(result);
}
