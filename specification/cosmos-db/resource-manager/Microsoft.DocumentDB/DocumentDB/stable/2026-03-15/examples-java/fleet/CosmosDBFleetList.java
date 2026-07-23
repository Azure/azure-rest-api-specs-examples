
/**
 * Samples for Fleet List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetList.json
     */
    /**
     * Sample code: CosmosDB Fleet List by subscription.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetListBySubscription(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleets().list(com.azure.core.util.Context.NONE);
    }
}
