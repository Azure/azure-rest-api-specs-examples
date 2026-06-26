const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an elastic pool.
 *
 * @summary updates an elastic pool.
 * x-ms-original-file: 2025-01-01/ElasticPoolUpdateMin.json
 */
async function updateAnElasticPoolWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.elasticPools.update(
    "sqlcrudtest-2369",
    "sqlcrudtest-8069",
    "sqlcrudtest-8102",
    {},
  );
  console.log(result);
}
