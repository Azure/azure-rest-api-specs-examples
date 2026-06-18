const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list PrefixListGlobalRulestackResource resources by Tenant
 *
 * @summary list PrefixListGlobalRulestackResource resources by Tenant
 * x-ms-original-file: 2026-05-11-preview/PrefixListGlobalRulestack_List_MaximumSet_Gen.json
 */
async function prefixListGlobalRulestackListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const resArray = new Array();
  for await (const item of client.prefixListGlobalRulestack.list("praval")) {
    resArray.push(item);
  }

  console.log(resArray);
}
