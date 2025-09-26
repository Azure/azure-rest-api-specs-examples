const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an elastic pool.
 *
 * @summary Creates or updates an elastic pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolCreateOrUpdateMax.json
 */
async function createOrUpdateElasticPoolWithAllParameter() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-2369";
  const serverName = "sqlcrudtest-8069";
  const elasticPoolName = "sqlcrudtest-8102";
  const parameters = {
    location: "Japan East",
    perDatabaseSettings: { maxCapacity: 2, minCapacity: 0.25 },
    sku: { name: "GP_Gen4_2", capacity: 2, tier: "GeneralPurpose" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.elasticPools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    elasticPoolName,
    parameters,
  );
  console.log(result);
}
