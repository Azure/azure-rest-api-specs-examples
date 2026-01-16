
/**
 * Samples for Appliances ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesListKeys.json
     */
    /**
     * Sample code: ListKeys Appliance.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void listKeysAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().listKeysWithResponse("testresourcegroup", "appliance01", null,
            com.azure.core.util.Context.NONE);
    }
}
