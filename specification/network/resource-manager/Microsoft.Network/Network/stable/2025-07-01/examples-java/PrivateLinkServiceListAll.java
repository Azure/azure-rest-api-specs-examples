
/**
 * Samples for PrivateLinkServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceListAll.json
     */
    /**
     * Sample code: List all private list service.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllPrivateListService(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().list(com.azure.core.util.Context.NONE);
    }
}
