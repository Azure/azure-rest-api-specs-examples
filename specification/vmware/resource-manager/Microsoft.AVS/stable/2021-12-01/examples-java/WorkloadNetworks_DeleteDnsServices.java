import com.azure.core.util.Context;

/** Samples for WorkloadNetworks DeleteDnsService. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteDnsServices.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteDnsService.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteDnsService("group1", "dnsService1", "cloud1", Context.NONE);
    }
}
