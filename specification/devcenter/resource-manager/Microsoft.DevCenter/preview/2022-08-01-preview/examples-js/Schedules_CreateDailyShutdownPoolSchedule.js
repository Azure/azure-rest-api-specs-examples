const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Schedule.
 *
 * @summary Creates or updates a Schedule.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_CreateDailyShutdownPoolSchedule.json
 */
async function schedulesCreateDailyShutdownPoolSchedule() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "DevProject";
  const poolName = "DevPool";
  const scheduleName = "autoShutdown";
  const body = {
    typePropertiesType: "StopDevBox",
    frequency: "Daily",
    state: "Enabled",
    time: "17:30",
    timeZone: "America/Los_Angeles",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.schedules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    projectName,
    poolName,
    scheduleName,
    body
  );
  console.log(result);
}

schedulesCreateDailyShutdownPoolSchedule().catch(console.error);
