import com.azure.core.util.Context;

/** Samples for WorkloadNetworks ListDnsZones. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_ListDnsZones.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListDnsZones(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listDnsZones("group1", "cloud1", Context.NONE);
    }
}
