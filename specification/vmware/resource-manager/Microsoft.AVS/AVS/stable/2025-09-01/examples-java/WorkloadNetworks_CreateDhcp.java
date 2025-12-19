
import com.azure.resourcemanager.avs.models.WorkloadNetworkDhcpServer;

/**
 * Samples for WorkloadNetworks CreateDhcp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_CreateDhcp.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDhcp.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().defineDhcp("dhcp1").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new WorkloadNetworkDhcpServer().withDisplayName("dhcpConfigurations1").withRevision(1L)
                .withServerAddress("40.1.5.1/24").withLeaseTime(86400L))
            .create();
    }
}
