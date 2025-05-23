const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Provides the status of the delete operations such as deleting backed up item. Once the operation has started, the
status code in the response would be Accepted. It will continue to be in this state till it reaches completion. On
successful completion, the status code will be OK. This method expects OperationID as an argument. OperationID is
part of the Location header of the operation response.
 *
 * @summary Provides the status of the delete operations such as deleting backed up item. Once the operation has started, the
status code in the response would be Accepted. It will continue to be in this state till it reaches completion. On
successful completion, the status code will be OK. This method expects OperationID as an argument. OperationID is
part of the Location header of the operation response.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/ProtectedItem_Delete_OperationResult.json
 */
async function getResultForProtectedItemDeleteOperation() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName =
    process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "PythonSDKBackupTestRg";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupOperationResults.get(vaultName, resourceGroupName, operationId);
  console.log(result);
}
