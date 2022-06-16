import com.azure.core.util.Context;

/** Samples for WorkloadNetworks DeleteSegment. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteSegment.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteSegment(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteSegment("group1", "cloud1", "segment1", Context.NONE);
    }
}
