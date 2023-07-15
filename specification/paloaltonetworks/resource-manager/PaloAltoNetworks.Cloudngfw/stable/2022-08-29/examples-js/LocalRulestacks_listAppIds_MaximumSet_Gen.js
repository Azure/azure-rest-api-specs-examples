const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of AppIds for LocalRulestack ApiVersion
 *
 * @summary List of AppIds for LocalRulestack ApiVersion
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/LocalRulestacks_listAppIds_MaximumSet_Gen.json
 */
async function localRulestacksListAppIdsMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "rgopenapi";
  const localRulestackName = "lrs1";
  const appIdVersion = "8543";
  const appPrefix = "pref";
  const skip = "a6a321";
  const top = 20;
  const options = {
    appIdVersion,
    appPrefix,
    skip,
    top,
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.localRulestacks.listAppIds(
    resourceGroupName,
    localRulestackName,
    options
  );
  console.log(result);
}
