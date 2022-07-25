const { OpenEnergyPlatformManagementServiceAPIs } = require("@azure/arm-oep");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns oep resource for a given name.
 *
 * @summary Returns oep resource for a given name.
 * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Get.json
 */
async function oepResourceGet() {
  const subscriptionId = "0000000-0000-0000-0000-000000000001";
  const resourceGroupName = "DummyResourceGroupName";
  const resourceName = "DummyResourceName";
  const credential = new DefaultAzureCredential();
  const client = new OpenEnergyPlatformManagementServiceAPIs(credential, subscriptionId);
  const result = await client.energyServices.get(resourceGroupName, resourceName);
  console.log(result);
}

oepResourceGet().catch(console.error);
