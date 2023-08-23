/** Samples for Appliances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListBySubscription.json
     */
    /**
     * Sample code: List Appliances by subscription.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void listAppliancesBySubscription(
        com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().list(com.azure.core.util.Context.NONE);
    }
}
