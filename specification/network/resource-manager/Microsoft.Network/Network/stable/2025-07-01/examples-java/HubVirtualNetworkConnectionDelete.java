
/**
 * Samples for HubVirtualNetworkConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/HubVirtualNetworkConnectionDelete.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void hubVirtualNetworkConnectionDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getHubVirtualNetworkConnections().delete("rg1", "virtualHub1", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
