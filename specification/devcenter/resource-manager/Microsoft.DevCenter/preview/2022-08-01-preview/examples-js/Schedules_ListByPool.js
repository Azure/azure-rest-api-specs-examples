const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists schedules for a pool
 *
 * @summary Lists schedules for a pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_ListByPool.json
 */
async function schedulesListByPool() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "TestProject";
  const poolName = "DevPool";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.schedules.listByPool(resourceGroupName, projectName, poolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

schedulesListByPool().catch(console.error);
