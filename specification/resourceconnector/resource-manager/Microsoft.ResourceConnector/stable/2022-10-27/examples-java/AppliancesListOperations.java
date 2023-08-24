/** Samples for Appliances ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListOperations.json
     */
    /**
     * Sample code: List Appliances operations.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void listAppliancesOperations(
        com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().listOperations(com.azure.core.util.Context.NONE);
    }
}
