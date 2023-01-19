const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the default Application Live View or update the existing Application Live View.
 *
 * @summary Create the default Application Live View or update the existing Application Live View.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/ApplicationLiveViews_CreateOrUpdate.json
 */
async function applicationLiveViewsCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const applicationLiveViewName = "default";
  const applicationLiveViewResource = {
    properties: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.applicationLiveViews.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    applicationLiveViewName,
    applicationLiveViewResource
  );
  console.log(result);
}
