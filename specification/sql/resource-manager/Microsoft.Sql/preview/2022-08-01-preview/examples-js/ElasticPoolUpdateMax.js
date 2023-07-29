const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an elastic pool.
 *
 * @summary Updates an elastic pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolUpdateMax.json
 */
async function updateAnElasticPoolWithAllParameter() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-2369";
  const serverName = "sqlcrudtest-8069";
  const elasticPoolName = "sqlcrudtest-8102";
  const parameters = {
    licenseType: "LicenseIncluded",
    perDatabaseSettings: { maxCapacity: 1, minCapacity: 0.25 },
    sku: { name: "BC_Gen4", capacity: 2, tier: "BusinessCritical" },
    zoneRedundant: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.elasticPools.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    elasticPoolName,
    parameters
  );
  console.log(result);
}
