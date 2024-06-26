const { OpenEnergyPlatformManagementServiceAPIs } = require("@azure/arm-oep");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Update.json
 */
async function oepResourceUpdate() {
  const subscriptionId = "0000000-0000-0000-0000-000000000001";
  const resourceGroupName = "DummyResourceGroupName";
  const resourceName = "DummyResourceName";
  const credential = new DefaultAzureCredential();
  const client = new OpenEnergyPlatformManagementServiceAPIs(credential, subscriptionId);
  const result = await client.energyServices.update(resourceGroupName, resourceName);
  console.log(result);
}

oepResourceUpdate().catch(console.error);
