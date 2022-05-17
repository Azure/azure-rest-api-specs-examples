Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listProtectableItemsWithBackupManagementTypeFilterAsAzureIaasVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const filter = "backupManagementType eq 'AzureIaasVM'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectableItems.list(
    vaultName,
    resourceGroupName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectableItemsWithBackupManagementTypeFilterAsAzureIaasVM().catch(console.error);
```
