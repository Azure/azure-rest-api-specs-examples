
/**
 * Samples for PrivateLinkServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceDelete.json
     */
    /**
     * Sample code: Delete private link service.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePrivateLinkService(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().delete("rg1", "testPls", com.azure.core.util.Context.NONE);
    }
}
