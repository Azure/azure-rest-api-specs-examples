
/**
 * Samples for FleetAnalytics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/fleet/CosmosDBFleetAnalyticsGet.json
     */
    /**
     * Sample code: CosmosDB FleetAnalytics Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetAnalyticsGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetAnalytics().getWithResponse("rg1", "fleet1", "storageAccount",
            com.azure.core.util.Context.NONE);
    }
}
