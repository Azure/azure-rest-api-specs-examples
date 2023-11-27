const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List FqdnListLocalRulestackResource resources by LocalRulestacks
 *
 * @summary List FqdnListLocalRulestackResource resources by LocalRulestacks
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/FqdnListLocalRulestack_ListByLocalRulestacks_MaximumSet_Gen.json
 */
async function fqdnListLocalRulestackListByLocalRulestacksMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "rgopenapi";
  const localRulestackName = "lrs1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fqdnListLocalRulestack.listByLocalRulestacks(
    resourceGroupName,
    localRulestackName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
