const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a springbootsites resource.
 *
 * @summary Delete a springbootsites resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_Delete_MaximumSet_Gen.json
 */
async function springbootsitesDeleteMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "chshxczdscjpcyvyethat";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootsites";
  const springbootsitesName = "xrmzlavpewxtfeitghdrj";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootsites.beginDeleteAndWait(
    resourceGroupName,
    springbootsitesName,
  );
  console.log(result);
}
