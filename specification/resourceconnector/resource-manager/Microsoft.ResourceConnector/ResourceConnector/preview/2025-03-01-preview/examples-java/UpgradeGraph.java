
/**
 * Samples for Appliances GetUpgradeGraph.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/UpgradeGraph.json
     */
    /**
     * Sample code: Get Appliance Upgrade Graph.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void
        getApplianceUpgradeGraph(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().getUpgradeGraphWithResponse("testresourcegroup", "appliance01", "stable",
            com.azure.core.util.Context.NONE);
    }
}
