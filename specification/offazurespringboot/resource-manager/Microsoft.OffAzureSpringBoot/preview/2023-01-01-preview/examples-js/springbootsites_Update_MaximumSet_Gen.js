const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a springbootsites resource.
 *
 * @summary Update a springbootsites resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_Update_MaximumSet_Gen.json
 */
async function springbootsitesUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "chshxczdscjpcyvyethat";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootsites";
  const springbootsitesName = "xrmzlavpewxtfeitghdrj";
  const springbootsites = {
    location: "icnumzvzzeqhuxtcefuqdcro",
    tags: { key9581: "cgdqvbknjrwcwuesquddsxu" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootsites.beginUpdateAndWait(
    resourceGroupName,
    springbootsitesName,
    springbootsites,
  );
  console.log(result);
}
