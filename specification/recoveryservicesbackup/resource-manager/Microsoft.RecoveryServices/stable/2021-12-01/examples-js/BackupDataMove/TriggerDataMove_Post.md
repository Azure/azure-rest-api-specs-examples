Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers Data Move Operation on target vault
 *
 * @summary Triggers Data Move Operation on target vault
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/BackupDataMove/TriggerDataMove_Post.json
 */
async function triggerDataMove() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "target-rsv";
  const resourceGroupName = "targetRG";
  const parameters = {
    correlationId: "MTg2OTcyMzM4NzYyMjc1NDY3Nzs1YmUzYmVmNi04YjJiLTRhOTItOTllYi01NTM0MDllYjk2NjE=",
    dataMoveLevel: "Vault",
    sourceRegion: "USGov Iowa",
    sourceResourceId:
      "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/sourceRG/providers/Microsoft.RecoveryServices/vaults/source-rsv",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.beginBMSTriggerDataMoveAndWait(
    vaultName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

triggerDataMove().catch(console.error);
```
