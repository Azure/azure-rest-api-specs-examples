const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List PrefixListGlobalRulestackResource resources by Tenant
 *
 * @summary List PrefixListGlobalRulestackResource resources by Tenant
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PrefixListGlobalRulestack_List_MaximumSet_Gen.json
 */
async function prefixListGlobalRulestackListMaximumSetGen() {
  const globalRulestackName = "praval";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const resArray = new Array();
  for await (let item of client.prefixListGlobalRulestack.list(globalRulestackName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
