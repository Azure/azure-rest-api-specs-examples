const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a devcenter.
 *
 * @summary Gets a devcenter.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/DevCenters_Get.json
 */
async function devCentersGet() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devCenters.get(resourceGroupName, devCenterName);
  console.log(result);
}

devCentersGet().catch(console.error);
