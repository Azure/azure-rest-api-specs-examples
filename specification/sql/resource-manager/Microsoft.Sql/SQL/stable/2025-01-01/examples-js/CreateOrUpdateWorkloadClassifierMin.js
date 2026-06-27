const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a workload classifier.
 *
 * @summary creates or updates a workload classifier.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateWorkloadClassifierMin.json
 */
async function createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.workloadClassifiers.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    "wlm_workloadgroup",
    "wlm_workloadclassifier",
    { memberName: "dbo" },
  );
  console.log(result);
}
