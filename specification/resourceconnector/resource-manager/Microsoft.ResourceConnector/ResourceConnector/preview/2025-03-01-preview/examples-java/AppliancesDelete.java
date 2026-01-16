
/**
 * Samples for Appliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesDelete.json
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
