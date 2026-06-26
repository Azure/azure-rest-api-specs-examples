const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an elastic pool.
 *
 * @summary creates or updates an elastic pool.
 * x-ms-original-file: 2025-01-01/ElasticPoolCreateWithVBSPreferredEnclaveType.json
 */
async function createOrUpdateElasticPoolWithPreferredEnclaveTypeParameterAsVBS() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.elasticPools.createOrUpdate(
    "sqlcrudtest-2369",
    "sqlcrudtest-8069",
    "sqlcrudtest-8102",
    { location: "Japan East", preferredEnclaveType: "VBS", sku: { name: "GP_Gen5_4" } },
  );
  console.log(result);
}
