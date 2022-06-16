const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the default Service Registry or update the existing Service Registry.
 *
 * @summary Create the default Service Registry or update the existing Service Registry.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ServiceRegistries_CreateOrUpdate.json
 */
async function serviceRegistriesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const serviceRegistryName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.serviceRegistries.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    serviceRegistryName
  );
  console.log(result);
}

serviceRegistriesCreateOrUpdate().catch(console.error);
