
/**
 * Samples for FleetspaceAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceAccountList.json
     */
    /**
     * Sample code: CosmosDB FleetspaceAccount List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceAccountList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaceAccounts().list("rg1", "fleet1", "fleetspace1",
            com.azure.core.util.Context.NONE);
    }
}
