
/**
 * Samples for WorkloadNetworks GetDnsService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetDnsService.json
     */
    /**
     * Sample code: WorkloadNetworks_GetDnsService.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getDnsServiceWithResponse("group1", "cloud1", "dnsService1",
            com.azure.core.util.Context.NONE);
    }
}
