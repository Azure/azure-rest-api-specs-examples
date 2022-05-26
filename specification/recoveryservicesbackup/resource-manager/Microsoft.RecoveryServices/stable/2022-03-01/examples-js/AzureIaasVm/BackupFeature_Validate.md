Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_9.0.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to It will validate if given feature with resource properties is supported in service
 *
 * @summary It will validate if given feature with resource properties is supported in service
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/BackupFeature_Validate.json
 */
async function checkAzureVMBackupFeatureSupport() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const azureRegion = "southeastasia";
  const parameters = {
    featureType: "AzureVMResourceBackup",
    vmSize: "Basic_A0",
    vmSku: "Premium",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.featureSupport.validate(azureRegion, parameters);
  console.log(result);
}

checkAzureVMBackupFeatureSupport().catch(console.error);
```
