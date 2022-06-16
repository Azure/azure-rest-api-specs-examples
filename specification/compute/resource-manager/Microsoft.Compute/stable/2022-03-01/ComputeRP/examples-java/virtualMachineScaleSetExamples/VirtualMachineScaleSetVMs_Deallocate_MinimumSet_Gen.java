import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Deallocate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Deallocate_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsDeallocateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .deallocate("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaa", Context.NONE);
    }
}
