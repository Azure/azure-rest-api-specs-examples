
/**
 * Samples for WorkloadNetworks DeleteDnsService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeleteDnsService.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteDnsService.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteDnsService("group1", "dnsService1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
