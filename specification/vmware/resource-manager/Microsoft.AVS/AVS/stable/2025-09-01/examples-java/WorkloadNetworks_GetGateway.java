
/**
 * Samples for WorkloadNetworks GetGateway.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetGateway.json
     */
    /**
     * Sample code: WorkloadNetworks_GetGateway.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetGateway(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getGatewayWithResponse("group1", "cloud1", "gateway1",
            com.azure.core.util.Context.NONE);
    }
}
