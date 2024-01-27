
import com.azure.resourcemanager.compute.models.AttachDetachDataDisksRequest;
import com.azure.resourcemanager.compute.models.DataDisksToAttach;
import com.azure.resourcemanager.compute.models.DataDisksToDetach;
import com.azure.resourcemanager.compute.models.DiskDetachOptionTypes;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSetVMs AttachDetachDataDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_AttachDetachDataDisks_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_AttachDetachDataDisks_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMAttachDetachDataDisksMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().attachDetachDataDisks(
            "rgcompute", "azure-vmscaleset", "0",
            new AttachDetachDataDisksRequest().withDataDisksToAttach(Arrays.asList(new DataDisksToAttach().withDiskId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d")
                .withLun(1),
                new DataDisksToAttach().withDiskId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_2_disk3_7d5e664bdafa49baa780eb2d128ff38e")
                    .withLun(2)))
                .withDataDisksToDetach(Arrays.asList(new DataDisksToDetach().withDiskId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x")
                    .withDetachOption(DiskDetachOptionTypes.FORCE_DETACH),
                    new DataDisksToDetach().withDiskId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_4_disk4_4d4e784bdafa49baa780eb2d256ff41z")
                        .withDetachOption(DiskDetachOptionTypes.FORCE_DETACH))),
            com.azure.core.util.Context.NONE);
    }
}
