
/**
 * Samples for Fleetspace Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceDelete.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().delete("rg1", "fleet1", "fleetspace1",
            com.azure.core.util.Context.NONE);
    }
}
