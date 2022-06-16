import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachineExtensions_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensions_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionsListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .listWithResponse("rgcompute", "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
