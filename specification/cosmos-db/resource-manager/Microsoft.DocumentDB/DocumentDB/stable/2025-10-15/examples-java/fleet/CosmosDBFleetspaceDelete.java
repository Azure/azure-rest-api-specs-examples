
/**
 * Samples for Fleetspace Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceDelete.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaces().delete("rg1", "fleet1", "fleetspace1",
            com.azure.core.util.Context.NONE);
    }
}
