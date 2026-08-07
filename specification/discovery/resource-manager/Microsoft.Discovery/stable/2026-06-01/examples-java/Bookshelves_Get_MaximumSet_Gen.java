
/**
 * Samples for Bookshelves GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Bookshelves_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Bookshelves_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void bookshelvesGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.bookshelves().getByResourceGroupWithResponse("rgdiscovery", "cfa586c95413ca2f8a",
            com.azure.core.util.Context.NONE);
    }
}
