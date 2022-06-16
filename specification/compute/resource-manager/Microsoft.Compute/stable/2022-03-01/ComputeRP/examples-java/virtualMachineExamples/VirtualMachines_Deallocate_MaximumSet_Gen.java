import com.azure.core.util.Context;

/** Samples for VirtualMachines Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_Deallocate_MaximumSet_Gen.json
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
            .deallocate("rgcompute", "aaaaaaaaaa", true, Context.NONE);
    }
}
