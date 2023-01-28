/** Samples for DedicatedHosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/dedicatedHostExamples/DedicatedHosts_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHosts_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostsDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .delete("rgcompute", "aaaaaaaaaaaaaaa", "aaaaa", com.azure.core.util.Context.NONE);
    }
}
