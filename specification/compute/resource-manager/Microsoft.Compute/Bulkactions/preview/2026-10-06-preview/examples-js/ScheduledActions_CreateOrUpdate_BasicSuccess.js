const { ComputeClient } = require("@azure/arm-compute-bulkactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a scheduled action.
 *
 * @summary creates or updates a scheduled action.
 * x-ms-original-file: 2026-10-06-preview/ScheduledActions_CreateOrUpdate_BasicSuccess.json
 */
async function _01CreateANewRecurringScheduledAction() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ComputeClient(credential, subscriptionId);
  const result = await client.scheduledActions.createOrUpdate("example-rg", "weekday-start", {
    properties: {
      resourceType: "VirtualMachine",
      actionType: "Start",
      startTime: "2026-09-15T07:00:00-07:00",
      schedule: {
        scheduledTime: "07:00:00",
        timeZone: "America/Los_Angeles",
        requestedWeekDays: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
      },
      notificationSettings: [],
    },
    location: "eastus",
  });
  console.log(result);
}
