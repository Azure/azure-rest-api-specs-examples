
/**
 * Samples for HubVirtualNetworkConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/HubVirtualNetworkConnectionGet.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void hubVirtualNetworkConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getHubVirtualNetworkConnections().getWithResponse("rg1", "virtualHub1", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
