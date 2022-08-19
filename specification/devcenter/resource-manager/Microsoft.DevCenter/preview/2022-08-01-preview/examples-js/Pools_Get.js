const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a machine pool
 *
 * @summary Gets a machine pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Pools_Get.json
 */
async function poolsGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const poolName = "{poolName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.get(resourceGroupName, projectName, poolName);
  console.log(result);
}

poolsGet().catch(console.error);
