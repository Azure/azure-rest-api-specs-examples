
/**
 * Samples for WorkloadNetworks ListDnsZones.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_ListDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_ListDnsZones.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListDnsZones(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listDnsZones("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
