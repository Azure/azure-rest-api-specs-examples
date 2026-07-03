
/**
 * Samples for PrivateLinkServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceGet.json
     */
    /**
     * Sample code: Get private link service.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPrivateLinkService(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().getByResourceGroupWithResponse("rg1", "testPls", null,
            com.azure.core.util.Context.NONE);
    }
}
