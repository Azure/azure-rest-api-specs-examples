/** Samples for DedicatedHosts ListByHostGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/dedicatedHostExamples/DedicatedHosts_ListByHostGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHosts_ListByHostGroup_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostsListByHostGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .listByHostGroup("rgcompute", "aaaa", com.azure.core.util.Context.NONE);
    }
}
