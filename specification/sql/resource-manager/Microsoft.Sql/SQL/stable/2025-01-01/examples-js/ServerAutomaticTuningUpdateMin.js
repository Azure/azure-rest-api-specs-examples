const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update automatic tuning options on server.
 *
 * @summary update automatic tuning options on server.
 * x-ms-original-file: 2025-01-01/ServerAutomaticTuningUpdateMin.json
 */
async function updatesServerAutomaticTuningSettingsWithMinimalProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c3aa9078-0000-0000-0000-e36f151182d7";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAutomaticTuningOperations.update(
    "default-sql-onebox",
    "testsvr11",
    { desiredState: "Auto" },
  );
  console.log(result);
}
