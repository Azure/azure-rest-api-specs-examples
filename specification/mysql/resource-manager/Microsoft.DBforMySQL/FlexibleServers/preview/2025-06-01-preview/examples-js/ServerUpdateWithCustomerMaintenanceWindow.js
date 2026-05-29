const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: 2025-06-01-preview/ServerUpdateWithCustomerMaintenanceWindow.json
 */
async function updateServerCustomerMaintenanceWindow() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.update("testrg", "mysqltestserver", {
    maintenanceWindow: {
      batchOfMaintenance: "Batch1",
      customWindow: "Enabled",
      dayOfWeek: 1,
      startHour: 8,
      startMinute: 0,
    },
  });
  console.log(result);
}
