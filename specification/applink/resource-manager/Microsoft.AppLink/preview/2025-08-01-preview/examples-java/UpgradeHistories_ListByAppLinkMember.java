
/**
 * Samples for UpgradeHistories ListByAppLinkMember.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/UpgradeHistories_ListByAppLinkMember.json
     */
    /**
     * Sample code: UpgradeHistories_ListByAppLinkMember.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void
        upgradeHistoriesListByAppLinkMember(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.upgradeHistories().listByAppLinkMember("test_rg", "applink-test-01", "member-01",
            com.azure.core.util.Context.NONE);
    }
}
