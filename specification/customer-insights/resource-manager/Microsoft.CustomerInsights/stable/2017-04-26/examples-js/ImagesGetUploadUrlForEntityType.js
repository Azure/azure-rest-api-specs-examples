const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets entity type (profile or interaction) image upload URL.
 *
 * @summary Gets entity type (profile or interaction) image upload URL.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ImagesGetUploadUrlForEntityType.json
 */
async function imagesGetUploadUrlForEntityType() {
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
  const result = await client.images.getUploadUrlForEntityType(
    resourceGroupName,
    hubName,
    parameters
  );
  console.log(result);
}

imagesGetUploadUrlForEntityType().catch(console.error);
