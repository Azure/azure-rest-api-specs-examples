const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an resource upload URL for an App, which may be artifacts or source archive.
 *
 * @summary Get an resource upload URL for an App, which may be artifacts or source archive.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_GetResourceUploadUrl.json
 */
async function appsGetResourceUploadUrl() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apps.getResourceUploadUrl(resourceGroupName, serviceName, appName);
  console.log(result);
}

appsGetResourceUploadUrl().catch(console.error);
