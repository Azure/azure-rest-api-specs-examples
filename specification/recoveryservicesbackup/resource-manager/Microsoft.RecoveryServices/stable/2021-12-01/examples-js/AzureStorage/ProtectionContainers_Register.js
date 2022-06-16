const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function registerAzureStorageProtectionContainers() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "swaggertestvault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const containerName = "StorageContainer;Storage;SwaggerTestRg;swaggertestsa";
  const parameters = {
    properties: {
      acquireStorageAccountLock: "Acquire",
      backupManagementType: "AzureStorage",
      containerType: "StorageContainer",
      friendlyName: "swaggertestsa",
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.register(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    parameters
  );
  console.log(result);
}

registerAzureStorageProtectionContainers().catch(console.error);
