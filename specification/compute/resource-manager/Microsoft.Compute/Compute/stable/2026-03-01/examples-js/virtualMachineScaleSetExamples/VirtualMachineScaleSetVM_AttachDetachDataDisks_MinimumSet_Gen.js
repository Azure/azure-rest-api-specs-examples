const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to attach and detach data disks to/from a virtual machine in a VM scale set.
 *
 * @summary attach and detach data disks to/from a virtual machine in a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_AttachDetachDataDisks_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMAttachDetachDataDisksMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.attachDetachDataDisks(
    "rgcompute",
    "azure-vmscaleset",
    "0",
    {
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
    },
  );
  console.log(result);
}
