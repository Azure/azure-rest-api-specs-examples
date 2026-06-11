const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the alerts for scope defined.
 *
 * @summary lists the alerts for scope defined.
 * x-ms-original-file: 2025-03-01/ResourceGroupAlerts.json
 */
async function resourceGroupAlerts() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.list(
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer",
  );
  console.log(result);
}
