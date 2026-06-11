const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the dimensions by the external cloud provider type.
 *
 * @summary lists the dimensions by the external cloud provider type.
 * x-ms-original-file: 2025-03-01/ExternalSubscriptionsDimensions.json
 */
async function externalSubscriptionDimensionList() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.dimensions.listByExternalCloudProviderType(
    "externalSubscriptions",
    "100",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
