
/**
 * Samples for Fleetspace Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceGet.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().getWithResponse("rg1", "fleet1", "fleetspace1",
            com.azure.core.util.Context.NONE);
    }
}
