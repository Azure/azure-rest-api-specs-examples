Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_9.0.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Registers the container with Recovery Services vault.
This is an asynchronous operation. To track the operation status, use location header to call get latest status of
the operation.
 *
 * @summary Registers the container with Recovery Services vault.
This is an asynchronous operation. To track the operation status, use location header to call get latest status of
the operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureStorage/ProtectionContainers_Register.json
 */
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
```
