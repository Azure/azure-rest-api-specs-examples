
/**
 * Samples for VirtualMachineScaleSetVMs Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Reimage_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Reimage_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMReimageMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().reimage("rgcompute", "aaaaaaa",
            "aaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
