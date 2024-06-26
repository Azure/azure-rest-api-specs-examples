
/**
 * Samples for WorkloadNetworks ListPublicIPs.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListPublicIPs.
     * json
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
