
/**
 * Samples for AppLinkMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinkMembers_Delete.json
     */
    /**
     * Sample code: AppLinkMembers_Delete.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinkMembersDelete(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinkMembers().delete("test_rg", "applink-test-01", "member-01", com.azure.core.util.Context.NONE);
    }
}
