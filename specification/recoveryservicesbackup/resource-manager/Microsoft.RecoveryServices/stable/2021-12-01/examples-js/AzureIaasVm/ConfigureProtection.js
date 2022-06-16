const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function enableProtectionOnAzureIaasVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const parameters = {
    properties: {
      policyId:
        "/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/DefaultPolicy",
      protectedItemType: "Microsoft.Compute/virtualMachines",
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

enableProtectionOnAzureIaasVM().catch(console.error);
