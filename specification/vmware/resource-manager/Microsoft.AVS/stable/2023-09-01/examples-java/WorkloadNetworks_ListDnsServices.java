
/**
 * Samples for WorkloadNetworks ListDnsServices.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListDnsServices.
     * json
     */
    /**
     * Sample code: WorkloadNetworks_ListDnsServices.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListDnsServices(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listDnsServices("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
