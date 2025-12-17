
/**
 * Samples for TrunkedNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/TrunkedNetworks_Get
     * .json
     */
    /**
     * Sample code: Get Trunked network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getTrunkedNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().getByResourceGroupWithResponse("resourceGroupName", "trunkedNetworkName",
            com.azure.core.util.Context.NONE);
    }
}
