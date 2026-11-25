
/**
 * Samples for QueryKeys ListBySearchService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/SearchListQueryKeysBySearchService.json
     */
    /**
     * Sample code: SearchListQueryKeysBySearchService.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        searchListQueryKeysBySearchService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getQueryKeys().listBySearchService("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
