
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks CreateDnsZone.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_CreateDnsZone.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDnsZone.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().defineDnsZone("dnsZone1").withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("dnsZone1").withDomain(Arrays.asList()).withDnsServerIps(Arrays.asList("1.1.1.1"))
            .withSourceIp("8.8.8.8").withRevision(1L).create();
    }
}
