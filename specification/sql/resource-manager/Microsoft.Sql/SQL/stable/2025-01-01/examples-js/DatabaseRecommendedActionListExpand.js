const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of database advisors.
 *
 * @summary gets a list of database advisors.
 * x-ms-original-file: 2025-01-01/DatabaseRecommendedActionListExpand.json
 */
async function listOfDatabaseRecommendedActionsForAllAdvisors() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseAdvisors.listByDatabase(
    "workloadinsight-demos",
    "misosisvr",
    "IndexAdvisor_test_3",
    { expand: "recommendedActions" },
  );
  console.log(result);
}
