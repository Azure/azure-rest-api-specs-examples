
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: Get PrivateLinkResource.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void getPrivateLinkResource(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateLinkResources().getWithResponse("res4303", "testfileshare01", "fileshare",
            com.azure.core.util.Context.NONE);
    }
}
