Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.1.1/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listProtectableItemsWithBackupManagementTypeFilterAsAzureStorage() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = "testRg";
  const fabricName = "Azure";
  const filter = "backupManagementType eq 'AzureStorage' and workloadType eq 'AzureFileShare'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.protectableContainers.list(
    vaultName,
    resourceGroupName,
    fabricName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectableItemsWithBackupManagementTypeFilterAsAzureStorage().catch(console.error);
```
