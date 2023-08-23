/** Samples for Appliances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesGet.json
     */
    /**
     * Sample code: Get Appliance.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void getAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager
            .appliances()
            .getByResourceGroupWithResponse("testresourcegroup", "appliance01", com.azure.core.util.Context.NONE);
    }
}
