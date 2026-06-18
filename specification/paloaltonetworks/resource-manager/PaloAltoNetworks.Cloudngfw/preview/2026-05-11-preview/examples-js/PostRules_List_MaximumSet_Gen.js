const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list PostRulesResource resources by Tenant
 *
 * @summary list PostRulesResource resources by Tenant
 * x-ms-original-file: 2026-05-11-preview/PostRules_List_MaximumSet_Gen.json
 */
async function postRulesListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const resArray = new Array();
  for await (const item of client.postRules.list("lrs1")) {
    resArray.push(item);
  }

  console.log(resArray);
}
