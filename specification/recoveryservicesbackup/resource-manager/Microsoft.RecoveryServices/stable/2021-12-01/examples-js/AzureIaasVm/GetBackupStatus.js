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
