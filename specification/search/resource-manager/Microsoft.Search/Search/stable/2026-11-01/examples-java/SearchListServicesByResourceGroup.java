
/**
 * Samples for Services ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/SearchListServicesByResourceGroup.json
     */
    /**
     * Sample code: SearchListServicesByResourceGroup.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        searchListServicesByResourceGroup(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
