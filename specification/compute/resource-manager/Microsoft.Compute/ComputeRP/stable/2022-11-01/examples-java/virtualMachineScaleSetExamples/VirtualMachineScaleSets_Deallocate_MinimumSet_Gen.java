/** Samples for VirtualMachineScaleSets Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_Deallocate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Deallocate_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsDeallocateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .deallocate("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
