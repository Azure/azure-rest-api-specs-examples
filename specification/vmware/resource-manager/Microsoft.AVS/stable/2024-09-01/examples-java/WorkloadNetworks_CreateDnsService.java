
import com.azure.resourcemanager.avs.models.DnsServiceLogLevelEnum;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks CreateDnsService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/WorkloadNetworks_CreateDnsService.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDnsService.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().defineDnsService("dnsService1").withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("dnsService1").withDnsServiceIp("5.5.5.5").withDefaultDnsZone("defaultDnsZone1")
            .withFqdnZones(Arrays.asList("fqdnZone1")).withLogLevel(DnsServiceLogLevelEnum.INFO).withRevision(1L)
            .create();
    }
}
