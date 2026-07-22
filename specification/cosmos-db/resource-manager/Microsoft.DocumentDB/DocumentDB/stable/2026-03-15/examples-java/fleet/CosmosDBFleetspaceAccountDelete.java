
/**
 * Samples for FleetspaceAccount Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceAccountDelete.json
     */
    /**
     * Sample code: CosmosDB FleetspaceAccount Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceAccountDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaceAccounts().delete("rg1", "fleet1", "fleetspace1", "db1",
            com.azure.core.util.Context.NONE);
    }
}
