const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Monitoring Setting and its properties.
 *
 * @summary Get the Monitoring Setting and its properties.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/MonitoringSettings_Get.json
 */
async function monitoringSettingsGet() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.monitoringSettings.get(resourceGroupName, serviceName);
  console.log(result);
}
