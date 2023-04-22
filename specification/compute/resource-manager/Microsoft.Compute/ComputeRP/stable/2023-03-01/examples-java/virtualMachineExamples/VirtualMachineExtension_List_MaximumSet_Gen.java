/** Samples for VirtualMachineExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachineExtension_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .listWithResponse("rgcompute", "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
