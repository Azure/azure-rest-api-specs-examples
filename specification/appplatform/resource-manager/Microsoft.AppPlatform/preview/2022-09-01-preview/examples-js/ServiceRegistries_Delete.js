const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disable the default Service Registry.
 *
 * @summary Disable the default Service Registry.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ServiceRegistries_Delete.json
 */
async function serviceRegistriesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const serviceRegistryName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.serviceRegistries.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    serviceRegistryName
  );
  console.log(result);
}

serviceRegistriesDelete().catch(console.error);
