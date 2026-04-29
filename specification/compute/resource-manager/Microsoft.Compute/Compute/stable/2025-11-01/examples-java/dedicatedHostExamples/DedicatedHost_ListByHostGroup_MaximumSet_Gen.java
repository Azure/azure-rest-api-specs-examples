
/**
 * Samples for DedicatedHosts ListByHostGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_ListByHostGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        dedicatedHostListByHostGroupMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().listByHostGroup("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
