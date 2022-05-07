Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.1.1/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Discovers all the containers in the subscription that can be backed up to Recovery Services Vault. This is an
asynchronous operation. To know the status of the operation, call GetRefreshOperationResult API.
 *
 * @summary Discovers all the containers in the subscription that can be backed up to Recovery Services Vault. This is an
asynchronous operation. To know the status of the operation, call GetRefreshOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/RefreshContainers.json
 */
async function triggerAzureVMDiscovery() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.refresh(
    vaultName,
    resourceGroupName,
    fabricName
  );
  console.log(result);
}

triggerAzureVMDiscovery().catch(console.error);
```
