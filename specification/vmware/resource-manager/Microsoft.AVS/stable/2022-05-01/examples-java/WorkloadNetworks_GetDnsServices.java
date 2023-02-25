/** Samples for WorkloadNetworks GetDnsService. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetDnsServices.json
     */
    /**
     * Sample code: WorkloadNetworks_GetDnsService.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .getDnsServiceWithResponse("group1", "cloud1", "dnsService1", com.azure.core.util.Context.NONE);
    }
}
