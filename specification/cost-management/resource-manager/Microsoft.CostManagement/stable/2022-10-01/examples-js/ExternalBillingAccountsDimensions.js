const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the dimensions by the external cloud provider type.
 *
 * @summary Lists the dimensions by the external cloud provider type.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalBillingAccountsDimensions.json
 */
async function externalBillingAccountDimensionList() {
  const externalCloudProviderType = "externalBillingAccounts";
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
