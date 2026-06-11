const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a shared scheduled action within the given scope.
 *
 * @summary create or update a shared scheduled action within the given scope.
 * x-ms-original-file: 2025-03-01/scheduledActions/scheduledAction-insightAlert-createOrUpdate-shared.json
 */
async function createOrUpdateInsightAlertScheduledActionByScope() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.scheduledActions.createOrUpdateByScope(
    "subscriptions/00000000-0000-0000-0000-000000000000",
    "dailyAnomalyByResource",
    {
      kind: "InsightAlert",
      displayName: "Daily anomaly by resource",
      notification: {
        subject: "Cost anomaly detected in the resource",
        to: ["user@gmail.com", "team@gmail.com"],
      },
      schedule: {
        endDate: new Date("2021-06-19T22:21:51.1287144Z"),
        frequency: "Daily",
        startDate: new Date("2020-06-19T22:21:51.1287144Z"),
      },
      status: "Enabled",
      viewId: "/providers/Microsoft.CostManagement/views/swaggerExample",
    },
    { ifMatch: "" },
  );
  console.log(result);
}
