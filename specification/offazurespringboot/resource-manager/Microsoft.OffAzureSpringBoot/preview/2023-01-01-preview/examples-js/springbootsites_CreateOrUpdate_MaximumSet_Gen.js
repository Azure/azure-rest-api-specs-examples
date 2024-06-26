const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a springbootsites resource.
 *
 * @summary Create a springbootsites resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_CreateOrUpdate_MaximumSet_Gen.json
 */
async function springbootsitesCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "chshxczdscjpcyvyethat";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootsites";
  const springbootsitesName = "xrmzlavpewxtfeitghdrj";
  const springbootsites = {
    extendedLocation: { name: "rywvpbfsqovhlfirtwisugsdsfsgf", type: "lvsb" },
    location: "tgobtvxktootwhhvjtsmpddvlqlrq",
    properties: {
      masterSiteId: "xsoimrgshsactearljwuljmi",
      migrateProjectId: "wwuattybgco",
    },
    tags: { key3558: "xeuhtglamqzj" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootsites.createOrUpdate(
    resourceGroupName,
    springbootsitesName,
    springbootsites,
  );
  console.log(result);
}
