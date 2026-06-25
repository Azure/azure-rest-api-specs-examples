
/**
 * Samples for DedicatedHosts ListByHostGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_ListByHostGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        dedicatedHostListByHostGroupMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().listByHostGroup("rgcompute", "aaaa",
            com.azure.core.util.Context.NONE);
    }
}
