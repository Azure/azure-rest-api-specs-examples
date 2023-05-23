const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a shared scheduled action within the given scope.
 *
 * @summary Create or update a shared scheduled action within the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-insightAlert-createOrUpdate-shared.json
 */
async function createOrUpdateInsightAlertScheduledActionByScope() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const name = "dailyAnomalyByResource";
  const ifMatch = "";
  const scheduledAction = {
    displayName: "Daily anomaly by resource",
    kind: "InsightAlert",
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
