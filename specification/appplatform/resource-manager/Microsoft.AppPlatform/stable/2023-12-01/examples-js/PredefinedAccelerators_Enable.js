const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable predefined accelerator.
 *
 * @summary Enable predefined accelerator.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/PredefinedAccelerators_Enable.json
 */
async function predefinedAcceleratorsEnable() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const applicationAcceleratorName = "default";
  const predefinedAcceleratorName = "acc-name";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.predefinedAccelerators.beginEnableAndWait(
    resourceGroupName,
    serviceName,
    applicationAcceleratorName,
    predefinedAcceleratorName,
  );
  console.log(result);
}
