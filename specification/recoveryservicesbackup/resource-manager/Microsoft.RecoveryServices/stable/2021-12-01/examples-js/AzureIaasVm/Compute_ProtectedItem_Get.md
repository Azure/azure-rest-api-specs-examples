Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getProtectedVirtualMachineDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName = "PythonSDKBackupTestRg";
  const fabricName = "Azure";
  const containerName = "iaasvmcontainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1";
  const protectedItemName = "vm;iaasvmcontainerv2;iaasvm-rg;iaasvm-1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItems.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName
  );
  console.log(result);
}

getProtectedVirtualMachineDetails().catch(console.error);
```
