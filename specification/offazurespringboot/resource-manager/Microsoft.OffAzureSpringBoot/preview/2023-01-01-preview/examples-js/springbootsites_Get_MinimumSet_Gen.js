const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a springbootsites resource.
 *
 * @summary Get a springbootsites resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_Get_MinimumSet_Gen.json
 */
async function springbootsitesGetMinimumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "chshxczdscjpcyvyethat";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootsites";
  const springbootsitesName = "xrmzlavpewxtfeitghdrj";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootsites.get(resourceGroupName, springbootsitesName);
  console.log(result);
}
