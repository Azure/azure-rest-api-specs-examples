
/**
 * Samples for AppLinks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinks_Get.json
     */
    /**
     * Sample code: AppLinks_Get.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinksGet(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinks().getByResourceGroupWithResponse("test_rg", "applink-test-01",
            com.azure.core.util.Context.NONE);
    }
}
