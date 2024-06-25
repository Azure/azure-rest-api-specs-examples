
/**
 * Samples for WorkloadNetworks GetDnsZone.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_GetDnsZone.json
     */
    /**
     * Sample code: WorkloadNetworks_GetDnsZone.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getDnsZoneWithResponse("group1", "cloud1", "dnsZone1",
            com.azure.core.util.Context.NONE);
    }
}
