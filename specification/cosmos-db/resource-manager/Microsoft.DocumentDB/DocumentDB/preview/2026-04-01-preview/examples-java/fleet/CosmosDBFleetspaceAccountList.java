
/**
 * Samples for FleetspaceAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/fleet/CosmosDBFleetspaceAccountList.json
     */
    /**
     * Sample code: CosmosDB FleetspaceAccount List.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceAccountList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaceAccounts().list("rg1", "fleet1", "fleetspace1",
            com.azure.core.util.Context.NONE);
    }
}
