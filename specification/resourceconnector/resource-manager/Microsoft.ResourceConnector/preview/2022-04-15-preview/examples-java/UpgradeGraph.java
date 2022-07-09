import com.azure.core.util.Context;

/** Samples for Appliances GetUpgradeGraph. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/UpgradeGraph.json
     */
    /**
     * Sample code: Get Appliance Upgrade Graph.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void getApplianceUpgradeGraph(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().getUpgradeGraphWithResponse("testresourcegroup", "appliance01", "stable", Context.NONE);
    }
}
