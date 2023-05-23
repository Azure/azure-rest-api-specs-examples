const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the dimensions by the external cloud provider type.
 *
 * @summary Lists the dimensions by the external cloud provider type.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionsDimensions.json
 */
async function externalSubscriptionDimensionList() {
  const externalCloudProviderType = "externalSubscriptions";
  const externalCloudProviderId = "100";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.dimensions.listByExternalCloudProviderType(
    externalCloudProviderType,
    externalCloudProviderId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
