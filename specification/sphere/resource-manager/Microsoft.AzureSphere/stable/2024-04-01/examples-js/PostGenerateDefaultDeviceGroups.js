const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates default device groups for the product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
 *
 * @summary Generates default device groups for the product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostGenerateDefaultDeviceGroups.json
 */
async function productsGenerateDefaultDeviceGroups() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const productName = "MyProduct1";
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.products.listGenerateDefaultDeviceGroups(
    resourceGroupName,
    catalogName,
    productName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
