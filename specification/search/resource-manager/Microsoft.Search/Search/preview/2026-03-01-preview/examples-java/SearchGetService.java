
/**
 * Samples for Services GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchGetService.json
     */
    /**
     * Sample code: SearchGetService.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchGetService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().getByResourceGroupWithResponse("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
