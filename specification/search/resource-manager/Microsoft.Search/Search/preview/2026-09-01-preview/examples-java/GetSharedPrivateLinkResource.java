
/**
 * Samples for SharedPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/GetSharedPrivateLinkResource.json
     */
    /**
     * Sample code: SharedPrivateLinkResourceGet.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void sharedPrivateLinkResourceGet(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getSharedPrivateLinkResources().getWithResponse("rg1", "mysearchservice",
            "testResource", com.azure.core.util.Context.NONE);
    }
}
