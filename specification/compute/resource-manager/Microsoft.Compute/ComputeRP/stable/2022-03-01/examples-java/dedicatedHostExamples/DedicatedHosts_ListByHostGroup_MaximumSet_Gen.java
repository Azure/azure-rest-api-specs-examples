import com.azure.core.util.Context;

/** Samples for DedicatedHosts ListByHostGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHosts_ListByHostGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHosts_ListByHostGroup_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostsListByHostGroupMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .listByHostGroup("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
