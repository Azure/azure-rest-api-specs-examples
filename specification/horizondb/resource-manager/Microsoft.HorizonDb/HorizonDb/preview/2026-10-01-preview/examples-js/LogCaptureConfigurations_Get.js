const { HorizonDbClient } = require("@azure/arm-horizondb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a log capture configuration for a HorizonDB cluster.
 *
 * @summary gets a log capture configuration for a HorizonDB cluster.
 * x-ms-original-file: 2026-10-01-preview/LogCaptureConfigurations_Get.json
 */
async function getThePostgreSQLServerLogCaptureConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new HorizonDbClient(credential, subscriptionId);
  const result = await client.logCaptureConfigurations.get(
    "example-resource-group",
    "example-cluster",
    "PostgreSqlServer",
  );
  console.log(result);
}
