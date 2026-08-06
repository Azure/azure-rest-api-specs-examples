
/**
 * Samples for BookshelfPrivateLinkResources ListByBookshelf.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BookshelfPrivateLinkResources_ListByBookshelf_MaximumSet_Gen.json
     */
    /**
     * Sample code: BookshelfPrivateLinkResources_ListByBookshelf_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void bookshelfPrivateLinkResourcesListByBookshelfMaximumSet(
        com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelfPrivateLinkResources().listByBookshelf("rgdiscovery", "4ee70172cf125c4793",
            com.azure.core.util.Context.NONE);
    }
}
