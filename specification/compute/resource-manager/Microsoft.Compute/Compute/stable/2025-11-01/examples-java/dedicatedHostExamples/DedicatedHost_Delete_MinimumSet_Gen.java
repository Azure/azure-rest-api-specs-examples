
/**
 * Samples for DedicatedHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void dedicatedHostDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().delete("rgcompute", "aaaaaaaaaaaaaaa", "aaaaa",
            com.azure.core.util.Context.NONE);
    }
}
