const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a pageable list of workload item of a specific container according to the query filter and the pagination
parameters.
 *
 * @summary Provides a pageable list of workload item of a specific container according to the query filter and the pagination
parameters.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureWorkload/BackupWorkloadItems_List.json
 */
async function listWorkloadItemsInContainer() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "suchandr-seacan-rsv";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "testRg";
  const fabricName = "Azure";
  const containerName = "VMAppContainer;Compute;bvtdtestag;sqlserver-1";
  const filter = "backupManagementType eq 'AzureWorkload'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupWorkloadItems.list(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
