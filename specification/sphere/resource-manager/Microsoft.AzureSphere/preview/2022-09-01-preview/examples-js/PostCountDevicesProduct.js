const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Counts devices in product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
 *
 * @summary Counts devices in product. '.default' and '.unassigned' are system defined values and cannot be used for product name.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostCountDevicesProduct.json
 */
async function productsCountDevices() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const productName = "MyProduct1";
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.products.countDevices(resourceGroupName, catalogName, productName);
  console.log(result);
}
