
/**
 * Samples for Bookshelves List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Bookshelves_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Bookshelves_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        bookshelvesListBySubscriptionMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelves().list(com.azure.core.util.Context.NONE);
    }
}
