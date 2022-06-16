import com.azure.core.util.Context;

/** Samples for WorkloadNetworks GetDhcp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetDhcpConfigurations.json
     */
    /**
     * Sample code: WorkloadNetworks_GetDhcp.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getDhcpWithResponse("group1", "dhcp1", "cloud1", Context.NONE);
    }
}
