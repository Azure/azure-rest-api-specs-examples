const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Triggers restore for a BackupInstance
 *
 * @summary Triggers restore for a BackupInstance
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/BackupInstanceOperations/TriggerRestore.json
 */
async function triggerRestore() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "000pikumar";
  const vaultName = "PratikPrivatePreviewVault1";
  const backupInstanceName = "testInstance1";
  const parameters = {
    identityDetails: {
      useSystemAssignedIdentity: false,
      userAssignedIdentityArmUrl:
        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourcegroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUami",
    },
    objectType: "AzureBackupRecoveryPointBasedRestoreRequest",
    recoveryPointId: "hardcodedRP",
    restoreTargetInfo: {
      datasourceAuthCredentials: {
        objectType: "SecretStoreBasedAuthCredentials",
        secretStoreResource: {
          secretStoreType: "AzureKeyVault",
          uri: "https://samplevault.vault.azure.net/secrets/credentials",
        },
      },
      datasourceInfo: {
        datasourceType: "Microsoft.DBforPostgreSQL/servers/databases",
        objectType: "Datasource",
        resourceID:
          "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb",
        resourceLocation: "",
        resourceName: "targetdb",
        resourceType: "Microsoft.DBforPostgreSQL/servers/databases",
        resourceUri: "",
      },
      datasourceSetInfo: {
        datasourceType: "Microsoft.DBforPostgreSQL/servers/databases",
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
    parameters,
  );
  console.log(result);
}
