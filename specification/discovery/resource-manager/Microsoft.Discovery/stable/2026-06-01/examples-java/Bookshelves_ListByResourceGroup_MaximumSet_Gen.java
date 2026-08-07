
/**
 * Samples for Bookshelves ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Bookshelves_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Bookshelves_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        bookshelvesListByResourceGroupMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelves().listByResourceGroup("rgdiscovery", com.azure.core.util.Context.NONE);
    }
}
