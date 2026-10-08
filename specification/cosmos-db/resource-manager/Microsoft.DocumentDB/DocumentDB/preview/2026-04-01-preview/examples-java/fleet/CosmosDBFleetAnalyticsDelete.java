
/**
 * Samples for FleetAnalytics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/fleet/CosmosDBFleetAnalyticsDelete.json
     */
    /**
     * Sample code: CosmosDB FleetAnalytics Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetAnalyticsDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetAnalytics().delete("rg1", "fleet1", "storageAccount",
            com.azure.core.util.Context.NONE);
    }
}
