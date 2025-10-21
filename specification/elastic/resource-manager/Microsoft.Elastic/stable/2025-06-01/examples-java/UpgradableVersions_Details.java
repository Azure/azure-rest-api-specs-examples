
/**
 * Samples for UpgradableVersions Details.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/UpgradableVersions_Details.
     * json
     */
    /**
     * Sample code: UpgradableVersions_Details.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void upgradableVersionsDetails(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.upgradableVersions().detailsWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
