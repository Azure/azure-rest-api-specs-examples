const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all maintenance events for a flexible server.
 *
 * @summary lists all maintenance events for a flexible server.
 * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsListByServerWithFilter.json
 */
async function listMaintenanceEventsFilteredByStatusForAServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.maintenanceEvents.list("exampleresourcegroup", "exampleserver", {
    maintenanceStatus: "Upcoming",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
