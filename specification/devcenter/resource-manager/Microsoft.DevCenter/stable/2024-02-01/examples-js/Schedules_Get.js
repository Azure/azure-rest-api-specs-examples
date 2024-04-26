const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a schedule resource.
 *
 * @summary Gets a schedule resource.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Schedules_Get.json
 */
async function schedulesGetByPool() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "TestProject";
  const poolName = "DevPool";
  const scheduleName = "autoShutdown";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.schedules.get(resourceGroupName, projectName, poolName, scheduleName);
  console.log(result);
}
