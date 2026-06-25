
/**
 * Samples for DedicatedHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHost_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void dedicatedHostDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().delete("rgcompute", "aaaaaa", "aaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
