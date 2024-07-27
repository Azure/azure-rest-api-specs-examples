const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation will stop backup for a backup instance and retains the backup data as per the policy (except latest Recovery point, which will be retained forever)
 *
 * @summary This operation will stop backup for a backup instance and retains the backup data as per the policy (except latest Recovery point, which will be retained forever)
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/BackupInstanceOperations/SuspendBackup_ResourceGuardEnabled.json
 */
async function suspendBackupsWithMua() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "testrg";
  const vaultName = "testvault";
  const backupInstanceName = "testbi";
  const parameters = {
    resourceGuardOperationRequests: [
      "/subscriptions/754ec39f-8d2a-44cf-bfbf-13107ac85c36/resourcegroups/mua-testing/providers/Microsoft.DataProtection/resourceGuards/gvjreddy-test-ecy-rg-reader/dppDisableSuspendBackupsRequests/default",
    ],
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.beginSuspendBackupsAndWait(
    resourceGroupName,
    vaultName,
    backupInstanceName,
    options,
  );
  console.log(result);
}
