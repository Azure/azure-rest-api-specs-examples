const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List predefined URL categories for rulestack
 *
 * @summary List predefined URL categories for rulestack
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/LocalRulestacks_listPredefinedUrlCategories_MaximumSet_Gen.json
 */
async function localRulestacksListPredefinedUrlCategoriesMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "rgopenapi";
  const localRulestackName = "lrs1";
  const skip = "a6a321";
  const top = 20;
  const options = {
    skip,
    top,
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.localRulestacks.listPredefinedUrlCategories(
    resourceGroupName,
    localRulestackName,
    options
  );
  console.log(result);
}
