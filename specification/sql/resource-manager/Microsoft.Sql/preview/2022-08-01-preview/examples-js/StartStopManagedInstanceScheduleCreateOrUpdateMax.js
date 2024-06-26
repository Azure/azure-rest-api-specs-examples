const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the managed instance's Start/Stop schedule.
 *
 * @summary Creates or updates the managed instance's Start/Stop schedule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleCreateOrUpdateMax.json
 */
async function createsOrUpdatesTheManagedInstanceStartOrStopScheduleWithAllOptionalParametersSpecified() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "schedulerg";
  const managedInstanceName = "schedulemi";
  const startStopScheduleName = "default";
  const parameters = {
    description: "This is a schedule for our Dev/Test environment.",
    scheduleList: [
      {
        startDay: "Thursday",
        startTime: "18:00",
        stopDay: "Thursday",
        stopTime: "17:00",
      },
      {
        startDay: "Thursday",
        startTime: "15:00",
        stopDay: "Thursday",
        stopTime: "14:00",
      },
    ],
    timeZoneId: "Central European Standard Time",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.startStopManagedInstanceSchedules.createOrUpdate(
    resourceGroupName,
    managedInstanceName,
    startStopScheduleName,
    parameters,
  );
  console.log(result);
}
