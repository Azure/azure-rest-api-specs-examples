const { PostgresClient } = require("@azure/arm-neonpostgres");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a Endpoint
 *
 * @summary delete a Endpoint
 * x-ms-original-file: 2025-03-01/Endpoints_Delete_MaximumSet_Gen.json
 */
async function endpointsDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "9B8E3300-C5FA-442B-A259-3F6F614D5BD4";
  const client = new PostgresClient(credential, subscriptionId);
  await client.endpoints.delete("rgneon", "test-org", "entity-name", "entity-name", "entity-name");
}
