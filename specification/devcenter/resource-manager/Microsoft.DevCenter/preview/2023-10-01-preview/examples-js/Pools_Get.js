const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a machine pool
 *
 * @summary Gets a machine pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Get.json
 */
async function poolsGet() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "DevProject";
  const poolName = "DevPool";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.get(resourceGroupName, projectName, poolName);
  console.log(result);
}
