
/**
 * Samples for DedicatedHosts ListByHostGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * dedicatedHostExamples/DedicatedHost_ListByHostGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_ListByHostGroup_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostListByHostGroupMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().listByHostGroup("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
