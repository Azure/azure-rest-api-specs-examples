Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.1.1/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getOperationStatus() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "source-rsv";
  const resourceGroupName = "sourceRG";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.getOperationStatus(vaultName, resourceGroupName, operationId);
  console.log(result);
}

getOperationStatus().catch(console.error);
```
