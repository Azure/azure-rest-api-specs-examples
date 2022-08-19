const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a schedule resource.
 *
 * @summary Gets a schedule resource.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_Get.json
 */
async function schedulesGetByPool() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "TestProject";
  const poolName = "DevPool";
  const scheduleName = "autoShutdown";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.schedules.get(resourceGroupName, projectName, poolName, scheduleName);
  console.log(result);
}

schedulesGetByPool().catch(console.error);
