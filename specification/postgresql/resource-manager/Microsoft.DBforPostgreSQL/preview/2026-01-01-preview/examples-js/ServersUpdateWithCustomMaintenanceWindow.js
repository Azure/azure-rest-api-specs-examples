const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 *
 * @summary updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 * x-ms-original-file: 2026-01-01-preview/ServersUpdateWithCustomMaintenanceWindow.json
 */
async function updateAnExistingServerWithCustomMaintenanceWindow() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.update("exampleresourcegroup", "exampleserver", {
    createMode: "Update",
    maintenanceWindow: { customWindow: "Enabled", dayOfWeek: 0, startHour: 8, startMinute: 0 },
  });
  console.log(result);
}
