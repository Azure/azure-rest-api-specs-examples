
/**
 * Samples for WorkloadNetworks GetPublicIp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetPublicIP.json
     */
    /**
     * Sample code: WorkloadNetworks_GetPublicIP.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetPublicIP(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getPublicIpWithResponse("group1", "cloud1", "publicIP1",
            com.azure.core.util.Context.NONE);
    }
}
