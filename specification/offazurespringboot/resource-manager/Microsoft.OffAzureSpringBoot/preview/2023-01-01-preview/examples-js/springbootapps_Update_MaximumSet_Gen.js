const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a springbootapps resource.
 *
 * @summary Update a springbootapps resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Update_MaximumSet_Gen.json
 */
async function springbootappsUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "jnetwlorzmxpxmcucorv";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootapps";
  const siteName = "pdfosfhtemfsaglvwjdyqlyeipucrd";
  const springbootappsName = "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf";
  const springbootapps = {};
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootapps.beginUpdateAndWait(
    resourceGroupName,
    siteName,
    springbootappsName,
    springbootapps,
  );
  console.log(result);
}
