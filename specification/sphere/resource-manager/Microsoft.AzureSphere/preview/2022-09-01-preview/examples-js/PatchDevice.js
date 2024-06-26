const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Device. Use '.unassigned' or '.default' for the device group and product names to move a device to the catalog level.
 *
 * @summary Update a Device. Use '.unassigned' or '.default' for the device group and product names to move a device to the catalog level.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PatchDevice.json
 */
async function devicesUpdate() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const productName = "MyProduct1";
  const deviceGroupName = "MyDeviceGroup1";
  const deviceName =
    "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.devices.beginUpdateAndWait(
    resourceGroupName,
    catalogName,
    productName,
    deviceGroupName,
    deviceName,
    properties
  );
  console.log(result);
}
