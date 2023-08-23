import java.util.Arrays;

/** Samples for WorkloadNetworks CreateDnsZone. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_CreateDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDnsZone.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .defineDnsZone("dnsZone1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("dnsZone1")
            .withDomain(Arrays.asList())
            .withDnsServerIps(Arrays.asList("1.1.1.1"))
            .withSourceIp("8.8.8.8")
            .withRevision(1L)
            .create();
    }
}
