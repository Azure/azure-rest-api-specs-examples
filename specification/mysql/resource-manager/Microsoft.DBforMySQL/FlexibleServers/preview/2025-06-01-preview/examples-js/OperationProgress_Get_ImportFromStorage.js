const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the operation result for a long running operation.
 *
 * @summary get the operation result for a long running operation.
 * x-ms-original-file: 2025-06-01-preview/OperationProgress_Get_ImportFromStorage.json
 */
async function operationProgressGetImportFromStorage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.operationProgress.get(
    "westus",
    "00000000-0000-0000-0000-000000000000",
  );
  console.log(result);
}
