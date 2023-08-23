/** Samples for WorkloadNetworks GetSegment. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_GetSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_GetSegment.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetSegment(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .getSegmentWithResponse("group1", "cloud1", "segment1", com.azure.core.util.Context.NONE);
    }
}
