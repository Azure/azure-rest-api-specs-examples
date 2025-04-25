const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to stops the database instance of the SAP system.
 *
 * @summary stops the database instance of the SAP system.
 * x-ms-original-file: 2024-09-01/SapDatabaseInstances_StopInstanceSoft.json
 */
async function softStopTheDatabaseInstanceOfTheSAPSystem() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapDatabaseInstances.stop("test-rg", "X00", "db0", {
    body: { softStopTimeoutSeconds: 300 },
  });
  console.log(result);
}
