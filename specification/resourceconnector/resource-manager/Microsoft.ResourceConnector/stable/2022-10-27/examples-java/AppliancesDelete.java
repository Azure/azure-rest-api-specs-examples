/** Samples for Appliances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesDelete.json
     */
    /**
     * Sample code: Delete Appliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void deleteAppliance(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().delete("testresourcegroup", "appliance01", com.azure.core.util.Context.NONE);
    }
}
