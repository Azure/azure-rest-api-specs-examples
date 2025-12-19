
/**
 * Samples for WorkloadNetworks GetSegment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetSegment.json
     */
    /**
     * Sample code: WorkloadNetworks_GetSegment.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetSegment(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getSegmentWithResponse("group1", "cloud1", "segment1",
            com.azure.core.util.Context.NONE);
    }
}
