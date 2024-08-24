
/**
 * Samples for VirtualMachineExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExamples/VirtualMachineExtension_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_List_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensions().listWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
