
import com.azure.resourcemanager.avs.models.WorkloadNetworkDnsZone;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks UpdateDnsZone.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/WorkloadNetworks_UpdateDnsZone.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateDnsZone.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkDnsZone resource = manager.workloadNetworks()
            .getDnsZoneWithResponse("group1", "cloud1", "dnsZone1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDisplayName("dnsZone1").withDomain(Arrays.asList())
            .withDnsServerIps(Arrays.asList("1.1.1.1")).withSourceIp("8.8.8.8").withRevision(1L).apply();
    }
}
