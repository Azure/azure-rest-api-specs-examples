
/**
 * Samples for Fleetspace List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceList.json
     */
    /**
     * Sample code: CosmosDB Fleetspace List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaces().list("rg1", "fleet1",
            com.azure.core.util.Context.NONE);
    }
}
