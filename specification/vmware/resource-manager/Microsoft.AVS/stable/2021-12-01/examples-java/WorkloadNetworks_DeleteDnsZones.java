import com.azure.core.util.Context;

/** Samples for WorkloadNetworks DeleteDnsZone. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteDnsZone.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteDnsZone("group1", "dnsZone1", "cloud1", Context.NONE);
    }
}
