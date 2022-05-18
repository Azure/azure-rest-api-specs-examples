Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAzureVirtualMachineBackupStatus() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const azureRegion = "southeastasia";
  const parameters = {
    resourceId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Compute/VirtualMachines/testVm",
    resourceType: "VM",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupStatus.get(azureRegion, parameters);
  console.log(result);
}

getAzureVirtualMachineBackupStatus().catch(console.error);
```
