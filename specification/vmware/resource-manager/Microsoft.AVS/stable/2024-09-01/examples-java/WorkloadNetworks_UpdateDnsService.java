
import com.azure.resourcemanager.avs.models.DnsServiceLogLevelEnum;
import com.azure.resourcemanager.avs.models.WorkloadNetworkDnsService;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks UpdateDnsService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/WorkloadNetworks_UpdateDnsService.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateDnsService.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateDnsService(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkDnsService resource = manager.workloadNetworks()
            .getDnsServiceWithResponse("group1", "cloud1", "dnsService1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDisplayName("dnsService1").withDnsServiceIp("5.5.5.5")
            .withDefaultDnsZone("defaultDnsZone1").withFqdnZones(Arrays.asList("fqdnZone1"))
            .withLogLevel(DnsServiceLogLevelEnum.INFO).withRevision(1L).apply();
    }
}
