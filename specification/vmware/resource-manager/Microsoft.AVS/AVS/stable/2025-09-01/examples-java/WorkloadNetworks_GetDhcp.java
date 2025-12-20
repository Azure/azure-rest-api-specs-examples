
/**
 * Samples for WorkloadNetworks GetDhcp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetDhcp.json
     */
    /**
     * Sample code: WorkloadNetworks_GetDhcp.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getDhcpWithResponse("group1", "dhcp1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
