Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_8.1.1/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the status of the asynchronous operations like backup, restore. The status can be in progress, completed
or failed. You can refer to the Operation Status enum for all the possible states of an operation. Some operations
create jobs. This method returns the list of jobs associated with operation.
 *
 * @summary Provides the status of the asynchronous operations like backup, restore. The status can be in progress, completed
or failed. You can refer to the Operation Status enum for all the possible states of an operation. Some operations
create jobs. This method returns the list of jobs associated with operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectionPolicyOperationStatuses_Get.json
 */
async function getProtectionPolicyOperationStatus() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicyOperationStatuses.get(
    vaultName,
    resourceGroupName,
    policyName,
    operationId
  );
  console.log(result);
}

getProtectionPolicyOperationStatus().catch(console.error);
```
