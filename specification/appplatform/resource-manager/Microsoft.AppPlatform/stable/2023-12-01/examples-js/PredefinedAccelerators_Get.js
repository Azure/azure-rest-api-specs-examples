const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the predefined accelerator.
 *
 * @summary Get the predefined accelerator.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/PredefinedAccelerators_Get.json
 */
async function predefinedAcceleratorsGet() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const applicationAcceleratorName = "default";
  const predefinedAcceleratorName = "acc-name";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.predefinedAccelerators.get(
    resourceGroupName,
    serviceName,
    applicationAcceleratorName,
    predefinedAcceleratorName,
  );
  console.log(result);
}
