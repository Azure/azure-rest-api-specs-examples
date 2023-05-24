const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a shared scheduled action within the given scope.
 *
 * @summary Create or update a shared scheduled action within the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-createOrUpdate-shared.json
 */
async function createOrUpdateScheduledActionByScope() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const name = "monthlyCostByResource";
  const ifMatch = "";
  const scheduledAction = {
    displayName: "Monthly Cost By Resource",
    fileDestination: { fileFormats: ["Csv"] },
    kind: "Email",
    notification: {
      subject: "Cost by resource this month",
      to: ["user@gmail.com", "team@gmail.com"],
    },
    schedule: {
      daysOfWeek: ["Monday"],
      endDate: new Date("2021-06-19T22:21:51.1287144Z"),
      frequency: "Monthly",
      hourOfDay: 10,
      startDate: new Date("2020-06-19T22:21:51.1287144Z"),
      weeksOfMonth: ["First", "Third"],
    },
    status: "Enabled",
    viewId: "/providers/Microsoft.CostManagement/views/swaggerExample",
  };
  const options = {
    ifMatch,
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.scheduledActions.createOrUpdateByScope(
    scope,
    name,
    scheduledAction,
    options
  );
  console.log(result);
}
