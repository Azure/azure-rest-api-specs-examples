
/**
 * Samples for Fleetspace List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceList.json
     */
    /**
     * Sample code: CosmosDB Fleetspace List.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().list("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
