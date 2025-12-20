
/**
 * Samples for WorkloadNetworks DeletePublicIp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeletePublicIP.json
     */
    /**
     * Sample code: WorkloadNetworks_DeletePublicIP.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeletePublicIP(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deletePublicIp("group1", "publicIP1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
