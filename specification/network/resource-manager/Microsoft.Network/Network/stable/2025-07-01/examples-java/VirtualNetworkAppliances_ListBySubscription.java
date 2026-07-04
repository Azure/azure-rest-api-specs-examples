
/**
 * Samples for VirtualNetworkAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkAppliances_ListBySubscription.json
     */
    /**
     * Sample code: List all virtual network appliances.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllVirtualNetworkAppliances(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkAppliances().list(com.azure.core.util.Context.NONE);
    }
}
