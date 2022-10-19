import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.WorkloadNetworkDhcp;
import com.azure.resourcemanager.avs.models.WorkloadNetworkDhcpServer;

/** Samples for WorkloadNetworks UpdateDhcp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_UpdateDhcpConfigurations.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateDhcp.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkDhcp resource =
            manager.workloadNetworks().getDhcpWithResponse("group1", "dhcp1", "cloud1", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new WorkloadNetworkDhcpServer().withRevision(1L).withServerAddress("40.1.5.1/24").withLeaseTime(86400L))
            .apply();
    }
}
