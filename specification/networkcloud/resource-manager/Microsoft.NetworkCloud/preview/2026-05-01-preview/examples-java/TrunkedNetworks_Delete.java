
/**
 * Samples for TrunkedNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/TrunkedNetworks_Delete.json
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
