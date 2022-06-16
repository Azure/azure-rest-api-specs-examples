const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables backup of an item or to modifies the backup policy information of an already backed up item. This is an
asynchronous operation. To know the status of the operation, call the GetItemOperationResult API.
 *
 * @summary Enables backup of an item or to modifies the backup policy information of an already backed up item. This is an
asynchronous operation. To know the status of the operation, call the GetItemOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/StopProtection.json
 */
async function stopProtectionWithRetainDataOnAzureIaasVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const parameters = {
    properties: {
      protectedItemType: "Microsoft.Compute/virtualMachines",
      protectionState: "ProtectionStopped",
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItems.createOrUpdate(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    parameters
  );
  console.log(result);
}

stopProtectionWithRetainDataOnAzureIaasVM().catch(console.error);
