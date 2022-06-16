import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachineExtensions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensions_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionsDeleteMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .delete("rgcompute", "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
