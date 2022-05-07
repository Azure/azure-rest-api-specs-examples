Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.1.1/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the details of the protection intent up item. This is an asynchronous operation. To know the status of the operation,
call the GetItemOperationResult API.
 *
 * @summary Provides the details of the protection intent up item. This is an asynchronous operation. To know the status of the operation,
call the GetItemOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureWorkload/BackupProtectionIntent_Get.json
 */
async function getProtectionIntentForAnItem() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = "myRG";
  const fabricName = "Azure";
  const intentObjectName = "249D9B07-D2EF-4202-AA64-65F35418564E";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.get(
    vaultName,
    resourceGroupName,
    fabricName,
    intentObjectName
  );
  console.log(result);
}

getProtectionIntentForAnItem().catch(console.error);
```
