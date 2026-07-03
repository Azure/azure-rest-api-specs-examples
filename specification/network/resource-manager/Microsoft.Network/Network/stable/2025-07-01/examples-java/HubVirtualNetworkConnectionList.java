
/**
 * Samples for HubVirtualNetworkConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/HubVirtualNetworkConnectionList.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void hubVirtualNetworkConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getHubVirtualNetworkConnections().list("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
