
/**
 * Samples for TrunkedNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2025-07-01-preview/examples/
     * TrunkedNetworks_Delete.json
     */
    /**
     * Sample code: Delete trunked network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteTrunkedNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().delete("resourceGroupName", "trunkedNetworkName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
