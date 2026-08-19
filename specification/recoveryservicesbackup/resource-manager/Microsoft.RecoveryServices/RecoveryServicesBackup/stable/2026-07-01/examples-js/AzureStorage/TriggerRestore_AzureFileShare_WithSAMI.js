const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 *
 * @summary restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 * x-ms-original-file: 2026-07-01/AzureStorage/TriggerRestore_AzureFileShare_WithSAMI.json
 */
async function restoreAzureFileShareToOriginalLocationWithManagedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  await client.restores.trigger(
    "swaggertestvault",
    "SwaggerTestRg",
    "Azure",
    "StorageContainer;Storage;SwaggerTestRg;swaggertestsa",
    "AzureFileShare;testshare",
    "932886657837421071",
    {
      properties: {
        objectType: "AzureFileShareRestoreRequest",
        recoveryType: "OriginalLocation",
        sourceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa",
        copyOptions: "Overwrite",
        restoreRequestType: "FullShareRestore",
        identityInfo: { isSystemAssignedIdentity: true },
      },
    },
  );
}
