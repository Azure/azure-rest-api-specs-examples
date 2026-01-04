
/**
 * Samples for HubVirtualNetworkConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * HubVirtualNetworkConnectionGet.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void hubVirtualNetworkConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getHubVirtualNetworkConnections().getWithResponse("rg1",
            "virtualHub1", "connection1", com.azure.core.util.Context.NONE);
    }
}
