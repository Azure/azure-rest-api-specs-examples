
/**
 * Samples for FleetAnalytics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/fleet/CosmosDBFleetAnalyticsList.json
     */
    /**
     * Sample code: CosmosDB FleetAnalytics List.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetAnalyticsList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetAnalytics().list("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
