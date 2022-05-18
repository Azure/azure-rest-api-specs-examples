Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.2.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous
operation. Status of the operation can be fetched using GetPolicyOperationResult API.
 *
 * @summary Provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous
operation. Status of the operation can be fetched using GetPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectionPolicies_Get.json
 */
async function getAzureIaasVMProtectionPolicyDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.get(vaultName, resourceGroupName, policyName);
  console.log(result);
}

getAzureIaasVMProtectionPolicyDetails().catch(console.error);
```
