
/**
 * Samples for AppLinkMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinkMembers_Get.json
     */
    /**
     * Sample code: AppLinkMembers_Get.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinkMembersGet(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinkMembers().getWithResponse("test_rg", "applink-test-01", "member-01",
            com.azure.core.util.Context.NONE);
    }
}
