
/**
 * Samples for Services Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/
     * UpgradeSearchServiceToLatestVersion.json
     */
    /**
     * Sample code: UpgradeSearchServiceToLatestVersion.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void upgradeSearchServiceToLatestVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().upgrade("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
