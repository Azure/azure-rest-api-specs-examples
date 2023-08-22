/** Samples for WorkloadNetworks ListDhcp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_ListDhcpConfigurations.json
     */
    /**
     * Sample code: WorkloadNetworks_ListDhcp.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listDhcp("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
