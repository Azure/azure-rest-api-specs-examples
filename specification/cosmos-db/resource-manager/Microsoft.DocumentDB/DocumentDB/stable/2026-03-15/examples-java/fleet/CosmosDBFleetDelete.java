
/**
 * Samples for Fleet Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetDelete.json
     */
    /**
     * Sample code: CosmosDB Fleet Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleets().delete("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
