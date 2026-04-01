
/**
 * Samples for Services Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/UpgradeSearchServiceToLatestVersion.json
     */
    /**
     * Sample code: UpgradeSearchServiceToLatestVersion.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        upgradeSearchServiceToLatestVersion(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().upgrade("rg1", "mysearchservice", com.azure.core.util.Context.NONE);
    }
}
