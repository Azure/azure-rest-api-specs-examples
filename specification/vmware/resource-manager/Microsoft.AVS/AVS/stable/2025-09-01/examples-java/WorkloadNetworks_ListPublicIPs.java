
/**
 * Samples for WorkloadNetworks ListPublicIPs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_ListPublicIPs.json
     */
    /**
     * Sample code: WorkloadNetworks_ListPublicIPs.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListPublicIPs(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listPublicIPs("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
