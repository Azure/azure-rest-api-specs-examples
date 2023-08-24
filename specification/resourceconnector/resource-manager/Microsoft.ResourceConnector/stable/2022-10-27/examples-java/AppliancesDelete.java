/** Samples for Appliances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesDelete.json
     */
    /**
     * Sample code: Delete Appliance.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void deleteAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().delete("testresourcegroup", "appliance01", com.azure.core.util.Context.NONE);
    }
}
