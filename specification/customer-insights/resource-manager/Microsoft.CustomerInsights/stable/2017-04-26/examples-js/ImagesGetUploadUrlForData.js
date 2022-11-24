const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets data image upload URL.
 *
 * @summary Gets data image upload URL.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ImagesGetUploadUrlForData.json
 */
async function imagesGetUploadUrlForData() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const parameters = {
    entityType: "Profile",
    entityTypeName: "Contact",
    relativePath: "images/profile1.png",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.images.getUploadUrlForData(resourceGroupName, hubName, parameters);
  console.log(result);
}

imagesGetUploadUrlForData().catch(console.error);
