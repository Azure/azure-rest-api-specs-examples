
/**
 * Samples for Appliances ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesListOperations.json
     */
    /**
     * Sample code: List Appliances operations.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void
        listAppliancesOperations(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().listOperations(com.azure.core.util.Context.NONE);
    }
}
