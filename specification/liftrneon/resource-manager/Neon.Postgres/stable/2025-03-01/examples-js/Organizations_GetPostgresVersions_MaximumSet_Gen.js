const { PostgresClient } = require("@azure/arm-neonpostgres");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to action to retrieve the PostgreSQL versions.
 *
 * @summary action to retrieve the PostgreSQL versions.
 * x-ms-original-file: 2025-03-01/Organizations_GetPostgresVersions_MaximumSet_Gen.json
 */
async function organizationsGetPostgresVersionsMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "9B8E3300-C5FA-442B-A259-3F6F614D5BD4";
  const client = new PostgresClient(credential, subscriptionId);
  const result = await client.organizations.getPostgresVersions("rgneon", {
    parameters: { version: 7 },
  });
  console.log(result);
}
