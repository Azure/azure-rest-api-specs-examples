const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates the capability image for the device. Use '.unassigned' or '.default' for the device group and product names to generate the image for a device that does not belong to a specific device group and product.
 *
 * @summary Generates the capability image for the device. Use '.unassigned' or '.default' for the device group and product names to generate the image for a device that does not belong to a specific device group and product.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostGenerateDeviceCapabilityImage.json
 */
async function devicesGenerateCapabilityImage() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const productName = "MyProduct1";
  const deviceGroupName = "myDeviceGroup1";
  const deviceName =
    "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
  const generateDeviceCapabilityRequest = {
    capabilities: ["ApplicationDevelopment"],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.devices.beginGenerateCapabilityImageAndWait(
    resourceGroupName,
    catalogName,
    productName,
    deviceGroupName,
    deviceName,
    generateDeviceCapabilityRequest,
  );
  console.log(result);
}
