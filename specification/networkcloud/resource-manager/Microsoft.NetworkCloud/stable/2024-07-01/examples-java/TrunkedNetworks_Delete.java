
/**
 * Samples for TrunkedNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/
     * TrunkedNetworks_Delete.json
     */
    /**
     * Sample code: Delete trunked network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteTrunkedNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().delete("resourceGroupName", "trunkedNetworkName", com.azure.core.util.Context.NONE);
    }
}
