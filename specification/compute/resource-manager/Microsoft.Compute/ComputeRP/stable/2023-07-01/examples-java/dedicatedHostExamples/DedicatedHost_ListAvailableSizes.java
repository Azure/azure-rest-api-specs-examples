/** Samples for DedicatedHosts ListAvailableSizes. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/dedicatedHostExamples/DedicatedHost_ListAvailableSizes.json
     */
    /**
     * Sample code: Get Available Dedicated Host Sizes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableDedicatedHostSizes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .listAvailableSizes("myResourceGroup", "myDedicatedHostGroup", "myHost", com.azure.core.util.Context.NONE);
    }
}
