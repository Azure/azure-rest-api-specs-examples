
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/SearchListOperations.json
     */
    /**
     * Sample code: SearchListOperations.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchListOperations(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
