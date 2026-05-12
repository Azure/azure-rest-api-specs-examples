const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to enables backup of an item or to modifies the backup policy information of an already backed up item. This is an
 * asynchronous operation. To know the status of the operation, call the GetItemOperationResult API.
 *
 * @summary enables backup of an item or to modifies the backup policy information of an already backed up item. This is an
 * asynchronous operation. To know the status of the operation, call the GetItemOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/StopProtection.json
 */
async function stopProtectionWithRetainDataOnAzureIaasVm() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItems.createOrUpdate(
    "NetSDKTestRsVault",
    "SwaggerTestRg",
    "Azure",
    "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    {
      properties: {
        protectedItemType: "Microsoft.Compute/virtualMachines",
        protectionState: "ProtectionStopped",
        sourceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
      },
    },
  );
  console.log(result);
}
