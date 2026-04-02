
/**
 * Samples for SharedPrivateLinkResources ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ListSharedPrivateLinkResourcesByService.json
     */
    /**
     * Sample code: ListSharedPrivateLinkResourcesByService.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        listSharedPrivateLinkResourcesByService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getSharedPrivateLinkResources().listByService("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
