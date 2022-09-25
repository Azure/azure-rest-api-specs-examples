import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachineExtensions_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensions_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionsDeleteMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaa", "aa", Context.NONE);
    }
}
