const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Counts devices in device group. '.default' and '.unassigned' are system defined values and cannot be used for product or device group name.
 *
 * @summary Counts devices in device group. '.default' and '.unassigned' are system defined values and cannot be used for product or device group name.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostCountDevicesDeviceGroup.json
 */
async function deviceGroupsCountDevices() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const productName = "MyProduct1";
  const deviceGroupName = "MyDeviceGroup1";
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.deviceGroups.countDevices(
    resourceGroupName,
    catalogName,
    productName,
    deviceGroupName,
  );
  console.log(result);
}
