/** Samples for UpgradableVersions Details. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/UpgradableVersions_Details.json
     */
    /**
     * Sample code: UpgradableVersions_Details.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void upgradableVersionsDetails(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .upgradableVersions()
            .detailsWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
