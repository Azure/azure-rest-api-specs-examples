const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates a Scheduled.
 *
 * @summary Partially updates a Scheduled.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Schedules_Patch.json
 */
async function schedulesUpdate() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "TestProject";
  const poolName = "DevPool";
  const scheduleName = "autoShutdown";
  const body = { time: "18:00" };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.schedules.beginUpdateAndWait(
    resourceGroupName,
    projectName,
    poolName,
    scheduleName,
    body
  );
  console.log(result);
}
