
/**
 * Samples for VirtualNetworkAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkAppliances_ListBySubscription.json
     */
    /**
     * Sample code: List all virtual network appliances.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllVirtualNetworkAppliances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkAppliances().list(com.azure.core.util.Context.NONE);
    }
}
