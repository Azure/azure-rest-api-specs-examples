
/**
 * Samples for WorkloadNetworks DeleteSegment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeleteSegment.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteSegment.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteSegment(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteSegment("group1", "cloud1", "segment1", com.azure.core.util.Context.NONE);
    }
}
