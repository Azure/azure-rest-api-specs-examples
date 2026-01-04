
import com.azure.resourcemanager.cosmos.models.FleetResourceUpdate;

/**
 * Samples for Fleet Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetUpdate.json
     */
    /**
     * Sample code: CosmosDB Fleet Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleets().updateWithResponse("rg1", "fleet1",
            new FleetResourceUpdate(), com.azure.core.util.Context.NONE);
    }
}
