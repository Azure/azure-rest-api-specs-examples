/** Samples for VirtualMachineExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachineExtension_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionDeleteMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaa", "aa", com.azure.core.util.Context.NONE);
    }
}
