
/**
 * Samples for WorkloadNetworks ListSegments.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_ListSegments.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListSegments(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listSegments("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
