const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Service Registry and its properties.
 *
 * @summary Get the Service Registry and its properties.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ServiceRegistries_Get.json
 */
async function serviceRegistriesGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const serviceRegistryName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.serviceRegistries.get(
    resourceGroupName,
    serviceName,
    serviceRegistryName
  );
  console.log(result);
}

serviceRegistriesGet().catch(console.error);
