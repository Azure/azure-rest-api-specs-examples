
/**
 * Samples for AdminKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchGetAdminKeys.json
     */
    /**
     * Sample code: SearchGetAdminKeys.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchGetAdminKeys(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getAdminKeys().getWithResponse("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
