
import com.azure.resourcemanager.compute.models.AttachDetachDataDisksRequest;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DataDisksToAttach;
import com.azure.resourcemanager.compute.models.DataDisksToDetach;
import com.azure.resourcemanager.compute.models.DiskDeleteOptionTypes;
import com.azure.resourcemanager.compute.models.DiskDetachOptionTypes;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import java.util.Arrays;

/**
 * Samples for VirtualMachines AttachDetachDataDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_AttachDetachDataDisks_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_AttachDetachDataDisks_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineAttachDetachDataDisksMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().attachDetachDataDisks("rgcompute",
            "aaaaaaaaaaaaaaaaaaaa",
            new AttachDetachDataDisksRequest().withDataDisksToAttach(Arrays.asList(new DataDisksToAttach().withDiskId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d")
                .withLun(1).withCaching(CachingTypes.READ_ONLY).withDeleteOption(DiskDeleteOptionTypes.DELETE)
                .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))
                .withWriteAcceleratorEnabled(true),
                new DataDisksToAttach().withDiskId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_2_disk3_7d5e664bdafa49baa780eb2d128ff38e")
                    .withLun(2).withCaching(CachingTypes.READ_WRITE).withDeleteOption(DiskDeleteOptionTypes.DETACH)
                    .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))
                    .withWriteAcceleratorEnabled(false)))
                .withDataDisksToDetach(Arrays.asList(new DataDisksToDetach().withDiskId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x")
                    .withDetachOption(DiskDetachOptionTypes.FORCE_DETACH),
                    new DataDisksToDetach().withDiskId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_4_disk4_4d4e784bdafa49baa780eb2d256ff41z")
                        .withDetachOption(DiskDetachOptionTypes.FORCE_DETACH))),
            com.azure.core.util.Context.NONE);
    }
}
