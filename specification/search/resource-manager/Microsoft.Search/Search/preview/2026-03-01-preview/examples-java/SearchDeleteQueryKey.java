
/**
 * Samples for QueryKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchDeleteQueryKey.json
     */
    /**
     * Sample code: SearchDeleteQueryKey.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchDeleteQueryKey(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getQueryKeys().deleteWithResponse("rg1", "mysearchservice", "<a query API key>",
            com.azure.core.util.Context.NONE);
    }
}
