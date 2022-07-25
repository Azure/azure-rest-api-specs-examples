const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a site update tags.
 *
 * @summary Updates a site update tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SiteUpdateTags.json
 */
async function updateMobileNetworkSiteTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const siteName = "testSite";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sites.updateTags(
    resourceGroupName,
    mobileNetworkName,
    siteName,
    parameters
  );
  console.log(result);
}

updateMobileNetworkSiteTags().catch(console.error);
