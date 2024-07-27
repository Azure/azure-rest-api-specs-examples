
/**
 * Samples for VirtualMachineScaleSetVMs Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Deallocate_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMDeallocateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().deallocate("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
