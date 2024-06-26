const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Catalog resources by subscription ID
 *
 * @summary List Catalog resources by subscription ID
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalogsSub.json
 */
async function catalogsListBySubscription() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.catalogs.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
