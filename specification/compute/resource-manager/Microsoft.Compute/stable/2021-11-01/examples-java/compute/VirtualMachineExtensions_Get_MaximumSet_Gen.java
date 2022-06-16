import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensions_Get_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionsGetMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .getWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", "aaaaaa", Context.NONE);
    }
}
