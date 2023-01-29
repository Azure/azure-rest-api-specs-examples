/** Samples for VirtualMachines Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachines_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_Deallocate_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesDeallocateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .deallocate("rgcompute", "aaaaaaaaaa", true, com.azure.core.util.Context.NONE);
    }
}
