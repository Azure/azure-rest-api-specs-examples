const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers restore for a BackupInstance
 *
 * @summary Triggers restore for a BackupInstance
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/BackupInstanceOperations/TriggerRestoreWithRehydration.json
 */
async function triggerRestoreWithRehydration() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "000pikumar";
  const vaultName = "PratikPrivatePreviewVault1";
  const backupInstanceName = "testInstance1";
  const parameters = {
    objectType: "AzureBackupRestoreWithRehydrationRequest",
    recoveryPointId: "hardcodedRP",
    rehydrationPriority: "High",
    rehydrationRetentionDuration: "7D",
    restoreTargetInfo: {
      datasourceInfo: {
        datasourceType: "OssDB",
        objectType: "Datasource",
        resourceID:
          "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb",
        resourceLocation: "",
        resourceName: "testdb",
        resourceType: "Microsoft.DBforPostgreSQL/servers/databases",
        resourceUri: "",
      },
      datasourceSetInfo: {
        datasourceType: "OssDB",
        objectType: "DatasourceSet",
        resourceID:
          "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest",
        resourceLocation: "",
        resourceName: "viveksipgtest",
        resourceType: "Microsoft.DBforPostgreSQL/servers",
        resourceUri: "",
      },
      objectType: "RestoreTargetInfo",
      recoveryOption: "FailIfExists",
      restoreLocation: "southeastasia",
    },
    sourceDataStoreType: "VaultStore",
    sourceResourceId:
      "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.beginTriggerRestoreAndWait(
    resourceGroupName,
    vaultName,
    backupInstanceName,
    parameters
  );
  console.log(result);
}
