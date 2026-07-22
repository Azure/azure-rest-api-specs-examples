
/**
 * Samples for FleetspaceAccount Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceAccountGet.json
     */
    /**
     * Sample code: CosmosDB FleetspaceAccount Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceAccountGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaceAccounts().getWithResponse("rg1", "fleet1", "fleetspace1", "db1",
            com.azure.core.util.Context.NONE);
    }
}
