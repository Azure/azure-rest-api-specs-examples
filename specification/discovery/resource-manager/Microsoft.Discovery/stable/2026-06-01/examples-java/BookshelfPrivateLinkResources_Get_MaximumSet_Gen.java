
/**
 * Samples for BookshelfPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BookshelfPrivateLinkResources_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: BookshelfPrivateLinkResources_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        bookshelfPrivateLinkResourcesGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelfPrivateLinkResources().getWithResponse("rgdiscovery", "28b448d6fa86171ee3", "connection",
            com.azure.core.util.Context.NONE);
    }
}
