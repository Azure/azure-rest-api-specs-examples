
/**
 * Samples for AppLinkMembers ListByAppLink.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinkMembers_ListByAppLink.json
     */
    /**
     * Sample code: AppLinkMembers_ListByAppLink.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinkMembersListByAppLink(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinkMembers().listByAppLink("test_rg", "applink-test-01", com.azure.core.util.Context.NONE);
    }
}
