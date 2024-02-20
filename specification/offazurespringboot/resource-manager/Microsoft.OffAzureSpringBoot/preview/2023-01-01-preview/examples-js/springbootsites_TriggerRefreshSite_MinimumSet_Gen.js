const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Trigger refresh springbootsites action
 *
 * @summary Trigger refresh springbootsites action
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_TriggerRefreshSite_MinimumSet_Gen.json
 */
async function springbootsitesTriggerRefreshSiteMinimumSetGen() {
  const subscriptionId = process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "z";
  const resourceGroupName = process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootsites";
  const springbootsitesName = "czarpuxwoafaqsuptutcwyu";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootsites.beginTriggerRefreshSiteAndWait(
    resourceGroupName,
    springbootsitesName,
  );
  console.log(result);
}
