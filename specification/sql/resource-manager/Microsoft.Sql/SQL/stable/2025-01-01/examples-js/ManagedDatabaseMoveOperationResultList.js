const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists managed database move operations.
 *
 * @summary lists managed database move operations.
 * x-ms-original-file: 2025-01-01/ManagedDatabaseMoveOperationResultList.json
 */
async function getsAllManagedDatabaseMoveOperationsForSpecifiedSubscriptionResourceGroupAndLocation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedDatabaseMoveOperations.listByLocation(
    "rg1",
    "westeurope",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
