
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMReimageParameters;

/**
 * Samples for VirtualMachineScaleSetVMs Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Reimage_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Reimage_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMReimageMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().reimage("rgcompute",
            "aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            new VirtualMachineScaleSetVMReimageParameters().withTempDisk(true).withForceUpdateOSDiskForEphemeral(true),
            com.azure.core.util.Context.NONE);
    }
}
