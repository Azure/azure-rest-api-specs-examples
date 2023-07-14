const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List PreRulesResource resources by Tenant
 *
 * @summary List PreRulesResource resources by Tenant
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PreRules_List_MinimumSet_Gen.json
 */
async function preRulesListMinimumSetGen() {
  const globalRulestackName = "lrs1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const resArray = new Array();
  for await (let item of client.preRules.list(globalRulestackName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
