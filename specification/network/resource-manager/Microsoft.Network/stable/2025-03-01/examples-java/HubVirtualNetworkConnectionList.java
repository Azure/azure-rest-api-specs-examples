
/**
 * Samples for HubVirtualNetworkConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * HubVirtualNetworkConnectionList.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void hubVirtualNetworkConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getHubVirtualNetworkConnections().list("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
