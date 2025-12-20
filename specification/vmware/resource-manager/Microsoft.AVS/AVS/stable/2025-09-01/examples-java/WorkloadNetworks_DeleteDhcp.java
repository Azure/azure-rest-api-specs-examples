
/**
 * Samples for WorkloadNetworks DeleteDhcp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeleteDhcp.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteDhcp.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteDhcp("group1", "cloud1", "dhcp1", com.azure.core.util.Context.NONE);
    }
}
