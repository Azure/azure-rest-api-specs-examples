const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disable the default Application Configuration Service.
 *
 * @summary Disable the default Application Configuration Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ConfigurationServices_Delete.json
 */
async function configurationServicesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const configurationServiceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.configurationServices.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    configurationServiceName
  );
  console.log(result);
}

configurationServicesDelete().catch(console.error);
