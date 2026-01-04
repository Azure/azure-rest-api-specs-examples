
/**
 * Samples for FleetspaceAccount Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceAccountGet.json
     */
    /**
     * Sample code: CosmosDB FleetspaceAccount Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceAccountGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaceAccounts().getWithResponse("rg1", "fleet1",
            "fleetspace1", "db1", com.azure.core.util.Context.NONE);
    }
}
