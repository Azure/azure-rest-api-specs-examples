/** Samples for Appliances GetUpgradeGraph. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/UpgradeGraph.json
     */
    /**
     * Sample code: Get Appliance Upgrade Graph.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void getApplianceUpgradeGraph(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager
            .appliances()
            .getUpgradeGraphWithResponse(
                "testresourcegroup", "appliance01", "stable", com.azure.core.util.Context.NONE);
    }
}
