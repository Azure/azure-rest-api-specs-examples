
/**
 * Samples for SharedPrivateLinkResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DeleteSharedPrivateLinkResource.json
     */
    /**
     * Sample code: SharedPrivateLinkResourceDelete.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void sharedPrivateLinkResourceDelete(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getSharedPrivateLinkResources().delete("rg1", "mysearchservice", "testResource",
            com.azure.core.util.Context.NONE);
    }
}
