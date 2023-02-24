/** Samples for WorkloadNetworks GetGateway. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetGateway.json
     */
    /**
     * Sample code: WorkloadNetworks_GetGateway.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetGateway(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .getGatewayWithResponse("group1", "cloud1", "gateway1", com.azure.core.util.Context.NONE);
    }
}
