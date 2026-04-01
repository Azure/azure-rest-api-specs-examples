
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchDeleteService.json
     */
    /**
     * Sample code: SearchDeleteService.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchDeleteService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().deleteWithResponse("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
