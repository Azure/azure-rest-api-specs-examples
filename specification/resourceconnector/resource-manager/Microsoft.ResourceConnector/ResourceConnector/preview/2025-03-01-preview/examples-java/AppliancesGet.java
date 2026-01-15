
/**
 * Samples for Appliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesGet.json
     */
    /**
     * Sample code: Get Appliance.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void getAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().getByResourceGroupWithResponse("testresourcegroup", "appliance01",
            com.azure.core.util.Context.NONE);
    }
}
