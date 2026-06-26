const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a workload group.
 *
 * @summary creates or updates a workload group.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateWorkloadGroupMin.json
 */
async function createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.workloadGroups.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    "smallrc",
    { maxResourcePercent: 100, minResourcePercent: 0, minResourcePercentPerRequest: 3 },
  );
  console.log(result);
}
