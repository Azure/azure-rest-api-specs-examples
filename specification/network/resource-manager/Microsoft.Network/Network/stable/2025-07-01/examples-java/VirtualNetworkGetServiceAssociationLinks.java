
/**
 * Samples for ServiceAssociationLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGetServiceAssociationLinks.json
     */
    /**
     * Sample code: Get Service Association Links.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getServiceAssociationLinks(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceAssociationLinks().listWithResponse("rg1", "vnet", "subnet",
            com.azure.core.util.Context.NONE);
    }
}
