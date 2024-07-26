
import com.azure.resourcemanager.compute.models.AttachDetachDataDisksRequest;
import com.azure.resourcemanager.compute.models.DataDisksToAttach;
import com.azure.resourcemanager.compute.models.DataDisksToDetach;
import java.util.Arrays;

/**
 * Samples for VirtualMachines AttachDetachDataDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_AttachDetachDataDisks_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_AttachDetachDataDisks_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineAttachDetachDataDisksMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().attachDetachDataDisks("rgcompute",
            "azure-vm",
            new AttachDetachDataDisksRequest().withDataDisksToAttach(Arrays.asList(new DataDisksToAttach().withDiskId(
                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d")))
                .withDataDisksToDetach(Arrays.asList(new DataDisksToDetach().withDiskId(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x"))),
            com.azure.core.util.Context.NONE);
    }
}
