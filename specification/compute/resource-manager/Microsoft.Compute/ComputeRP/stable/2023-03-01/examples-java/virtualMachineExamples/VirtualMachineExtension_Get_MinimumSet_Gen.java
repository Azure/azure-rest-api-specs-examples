/** Samples for VirtualMachineExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachineExtension_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_Get_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionGetMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .getWithResponse("rgcompute", "myVM", "myVMExtension", null, com.azure.core.util.Context.NONE);
    }
}
