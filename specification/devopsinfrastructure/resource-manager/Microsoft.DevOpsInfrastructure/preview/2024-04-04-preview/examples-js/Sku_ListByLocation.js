const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List ResourceSku resources by subscription ID
 *
 * @summary List ResourceSku resources by subscription ID
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/Sku_ListByLocation.json
 */
async function skuListByLocation() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const locationName = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sku.listByLocation(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
