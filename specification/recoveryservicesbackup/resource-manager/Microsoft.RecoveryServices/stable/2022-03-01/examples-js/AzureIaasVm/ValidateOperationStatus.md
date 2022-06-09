```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches the status of a triggered validate operation. The status can be in progress, completed
or failed. You can refer to the OperationStatus enum for all the possible states of the operation.
If operation has completed, this method returns the list of errors obtained while validating the operation.
 *
 * @summary Fetches the status of a triggered validate operation. The status can be in progress, completed
or failed. You can refer to the OperationStatus enum for all the possible states of the operation.
If operation has completed, this method returns the list of errors obtained while validating the operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/ValidateOperationStatus.json
 */
async function getOperationStatusOfValidateOperation() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.validateOperationStatuses.get(
    vaultName,
    resourceGroupName,
    operationId
  );
  console.log(result);
}

getOperationStatusOfValidateOperation().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_9.0.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
