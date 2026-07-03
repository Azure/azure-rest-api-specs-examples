
/**
 * Samples for ResourceNavigationLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGetResourceNavigationLinks.json
     */
    /**
     * Sample code: Get Resource Navigation Links.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getResourceNavigationLinks(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getResourceNavigationLinks().listWithResponse("rg1", "vnet", "subnet",
            com.azure.core.util.Context.NONE);
    }
}
