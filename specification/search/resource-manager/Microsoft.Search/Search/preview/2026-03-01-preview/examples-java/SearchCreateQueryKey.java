
/**
 * Samples for QueryKeys Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchCreateQueryKey.json
     */
    /**
     * Sample code: SearchCreateQueryKey.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchCreateQueryKey(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getQueryKeys().createWithResponse("rg1", "mysearchservice",
            "An API key granting read-only access to the documents collection of an index.",
            com.azure.core.util.Context.NONE);
    }
}
