/** Samples for Appliances ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListKeys.json
     */
    /**
     * Sample code: ListKeys Appliance.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void listKeysAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager
            .appliances()
            .listKeysWithResponse("testresourcegroup", "appliance01", null, com.azure.core.util.Context.NONE);
    }
}
