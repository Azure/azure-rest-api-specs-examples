
/**
 * Samples for WorkloadNetworks DeleteDnsZone.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeleteDnsZone.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteDnsZone.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteDnsZone("group1", "dnsZone1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
