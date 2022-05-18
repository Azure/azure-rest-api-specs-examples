Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a pageable list of all intents that are present within a vault.
 *
 * @summary Provides a pageable list of all intents that are present within a vault.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureWorkload/BackupProtectionIntent_List.json
 */
async function listProtectionIntentWithBackupManagementTypeFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = "myRG";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectionIntent.list(vaultName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectionIntentWithBackupManagementTypeFilter().catch(console.error);
```
