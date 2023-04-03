const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an resource upload URL for an App, which may be artifacts or source archive.
 *
 * @summary Get an resource upload URL for an App, which may be artifacts or source archive.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Apps_GetResourceUploadUrl.json
 */
async function appsGetResourceUploadUrl() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apps.getResourceUploadUrl(resourceGroupName, serviceName, appName);
  console.log(result);
}
