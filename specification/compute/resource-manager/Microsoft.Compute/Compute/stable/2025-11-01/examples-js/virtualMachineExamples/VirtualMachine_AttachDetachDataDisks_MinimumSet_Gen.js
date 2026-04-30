const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to attach and detach data disks to/from the virtual machine.
 *
 * @summary attach and detach data disks to/from the virtual machine.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_AttachDetachDataDisks_MinimumSet_Gen.json
 */
async function virtualMachineAttachDetachDataDisksMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.attachDetachDataDisks("rgcompute", "azure-vm", {
    dataDisksToAttach: [
      {
        diskId:
          "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d",
      },
    ],
    dataDisksToDetach: [
      {
        diskId:
          "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x",
      },
    ],
  });
  console.log(result);
}
